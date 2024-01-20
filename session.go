package event_pulse

type Session interface {
	Add(stream any)
	Track(entry *EventEntry)
	Find(streamName string, streamId any) (any, error)
	Complete()
	Close() error
}

type SessionImpl struct {
	provider      *StreamSerializerProvider
	persistor     EventPersistor
	pendingEvents []*EventEntry
	hydrating     bool
	completed     bool
}

func NewSession(provider *StreamSerializerProvider, persistor EventPersistor) Session {
	return &SessionImpl{
		provider:      provider,
		persistor:     persistor,
		pendingEvents: make([]*EventEntry, 0),
		hydrating:     false,
		completed:     false,
	}
}

func (s *SessionImpl) Add(stream any) {
}

func (s *SessionImpl) Track(entry *EventEntry) {
	if s.hydrating {
		return
	}

	s.pendingEvents = append(s.pendingEvents, entry)
}

func (s *SessionImpl) Find(streamName string, streamId any) (any, error) {
	s.hydrating = true
	defer func() { s.hydrating = false }()

	serializer := s.provider.Get(streamName)
	events := s.persistor.GetEvents(streamName, streamId)
	var stream any
	for _, event := range events {
		obj, err := serializer.Deserialize(event.EventType, event.EventData)
		if err != nil {
			return nil, err
		}

		stream = serializer.Aggregate(s, stream, obj)
	}

	return stream, nil
}

func (s *SessionImpl) Complete() {
	s.completed = true
}

func (s *SessionImpl) Close() error {
	if !s.completed {
		return nil
	}

	for _, entry := range s.pendingEvents {
		serializer := s.provider.Get(entry.StreamName)
		eventType, eventData, err := serializer.Serialize(entry.EventObject)
		if err != nil {
			return err
		}

		s.persistor.Persist(entry.StreamName, entry.StreamId, entry.Revision, eventType, eventData)
	}

	return nil
}
