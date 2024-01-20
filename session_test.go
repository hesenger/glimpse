package event_pulse

import (
	"testing"
	"time"
)

var sf *SessionFactory

func init() {
	provider := NewStreamSerializerProvider()
	serializer := StreamSerializer(&BookingSerializer{})
	provider.Register("booking", &serializer)
	sf = NewSessionFactory(provider)
}

func TestEventsArePersistedAcrossSessions(t *testing.T) {
	session := sf.OpenSession()
	new := NewBooking(session, BookingCreated{
		BookingId:  "123",
		CheckIn:    time.Now(),
		CheckOut:   time.Now().AddDate(0, 0, 3),
		TotalPrice: 300,
		AmountPaid: 0,
	})

	session.Track(new)
	session.Complete()
	session.Close() // persist events

	session = sf.OpenSession()
	booking, ok := session.Find("booking", "123").(*Booking)
	if ok {
		t.Error("booking not found")
	}

	if booking.PendingAmount() != 300 {
		t.Error("pending amount should be 300")
	}

	session.Complete()
	session.Close()
}
