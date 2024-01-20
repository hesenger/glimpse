package event_pulse

type SessionFactory struct {
	provider *StreamSerializerProvider
}

func NewSessionFactory(provider *StreamSerializerProvider) *SessionFactory {
	return &SessionFactory{
		provider: provider,
	}
}

func (sf *SessionFactory) OpenSession() *Session {
	return NewSession(sf.provider)
}
