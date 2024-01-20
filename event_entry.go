package event_pulse

type EventEntry struct {
	StreamName  string
	StreamId    any
	Revision    int
	EventObject any
}
