package event_pulse

type EventStream struct {
	session    Session
	streamName string
	streamId   any
	events     []any
}

func NewEventStream(session Session, streamName string, streamId any) *EventStream {
	return &EventStream{
		session:    session,
		streamName: streamName,
		streamId:   streamId,
		events:     make([]any, 0),
	}
}

func (s *EventStream) Append(event any) {
	s.events = append(s.events, event)
	s.session.Track(&EventEntry{
		StreamName:  s.streamName,
		StreamId:    s.streamId,
		Revision:    len(s.events),
		EventObject: event,
	})
}
