package event_pulse

type Session struct {
	provider      *StreamSerializerProvider
	pendingEvents []any
	hydrating     bool
	completed     bool
}

func NewSession(provider *StreamSerializerProvider) *Session {
	return &Session{
		provider:      provider,
		pendingEvents: make([]any, 0),
		hydrating:     false,
		completed:     false,
	}
}

func (s *Session) Track(obj any) {
	if s.hydrating {
		return
	}

	s.pendingEvents = append(s.pendingEvents, obj)
}

func (s *Session) Find(streanName string, streamId any) any {
	return nil
}

func (s *Session) Complete() {
	s.completed = true
}

func (s *Session) Close() {
	if !s.completed {
		return
	}
}
