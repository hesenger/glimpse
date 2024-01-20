package glimpse

type EventPersistor interface {
	Persist(streamName string, streamId any, revision int, eventType string, eventData string)
	GetEvents(streamName string, streamId any) []PersistedEvent
}
