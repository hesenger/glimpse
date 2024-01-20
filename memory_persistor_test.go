package event_pulse

import "fmt"

type MemoryPersistor struct {
	events map[string][]PersistedEvent
}

func NewMemoryPersistor() *MemoryPersistor {
	return &MemoryPersistor{
		events: make(map[string][]PersistedEvent),
	}
}

func (p *MemoryPersistor) Persist(streamName string, streamId any, revision int, eventType string, eventData string) {
	key := fmt.Sprintf("%s:%s", streamName, streamId)
	events := p.events[key]
	if events == nil {
		events = make([]PersistedEvent, 0)
	}

	events = append(events, PersistedEvent{
		EventType: eventType,
		EventData: eventData,
	})
}

func (p *MemoryPersistor) GetEvents(streamName string, streamId any) []PersistedEvent {
	key := fmt.Sprintf("%s:%s", streamName, streamId)
	events := p.events[key]
	return events
}
