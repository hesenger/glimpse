package event_pulse

import (
	"testing"
	"time"
)

var sf *SessionFactory

func init() {
	provider := NewStreamSerializerProvider()
	persistor := NewMemoryPersistor()
	serializer := &BookingSerializer{}
	provider.Register("booking", serializer)
	sf = NewSessionFactory(provider, persistor)
}

func TestEventsArePersistedAcrossSessions(t *testing.T) {
	session := sf.OpenSession()
	new := NewBooking(session, &BookingCreated{
		BookingId:  "123",
		CheckIn:    time.Now(),
		CheckOut:   time.Now().AddDate(0, 0, 3),
		TotalPrice: 300,
		AmountPaid: 0,
	})

	session.Add(new)
	session.Complete()
	session.Close() // persist events

	session = sf.OpenSession()
	res, err := session.Find("booking", "123")
	if err != nil {
		t.Error(err)
	}

	booking, ok := res.(*Booking)
	if !ok {
		t.Error("booking should be of type Booking")
	}

	if booking.PendingAmount() != 300 {
		t.Error("pending amount should be 300")
	}

	session.Complete()
	session.Close()
}
