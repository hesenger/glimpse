package event_pulse

type StreamSerializer interface {
	// Serialize an event of that stream.
	// Returns the eventType and the eventData.
	Serialize(stream any) (string, string)

	// Deserialize an event of that stream.
	// Returns the deserialized event.
	Deserialize(eventType string, eventData string) any

	// Aggregate the event into the stream.
	// Returns the aggregated stream.
	Aggregate(stream any, event any) any
}
