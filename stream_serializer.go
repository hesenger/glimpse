package glimpse

type StreamSerializer interface {
	// Serialize an event of that stream.
	// Returns the eventType and the eventData.
	Serialize(event any) (string, string, error)

	// Deserialize an event of that stream.
	// Returns the deserialized event.
	Deserialize(eventType string, eventData string) (any, error)

	// Aggregate the event into the stream.
	// Returns the aggregated stream.
	Aggregate(session Session, stream any, event any) any
}
