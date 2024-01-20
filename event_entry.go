package glimpse

type EventEntry struct {
	StreamName  string
	StreamId    any
	Revision    int
	EventObject any
}
