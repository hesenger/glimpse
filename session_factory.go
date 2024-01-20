package glimpse

type SessionFactory struct {
	provider  *StreamSerializerProvider
	persistor EventPersistor
}

func NewSessionFactory(provider *StreamSerializerProvider, persistor EventPersistor) *SessionFactory {
	return &SessionFactory{
		provider:  provider,
		persistor: persistor,
	}
}

func (sf *SessionFactory) OpenSession() Session {
	return NewSession(sf.provider, sf.persistor)
}
