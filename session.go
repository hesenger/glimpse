package event_pulse

type Session interface {
	Add(stream any)
	Track(entry *EventEntry)
	Find(streamName string, streamId any) any
	Complete()
	Close()
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

func (s *SessionImpl) Find(streamName string, streamId any) any {
	serializer := s.provider.Get(streamName)
	events := s.persistor.GetEvents(streamName, streamId)
	var stream any
	for _, event := range events {
		obj := serializer.Deserialize(event.EventType, event.EventData)
		stream = serializer.Aggregate(stream, obj)
	}

	return stream
}

func (s *SessionImpl) Complete() {
	s.completed = true
}

func (s *SessionImpl) Close() {
	if !s.completed {
		return
	}

	for _, entry := range s.pendingEvents {
		serializer := s.provider.Get(entry.StreamName)
		eventType, eventData := serializer.Serialize(entry.EventObject)
		s.persistor.Persist(entry.StreamName, entry.StreamId, entry.Revision, eventType, eventData)
	}
}
