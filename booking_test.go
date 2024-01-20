package event_pulse

import "time"

type Booking struct {
	events EventStream

	bookingId  string
	checkIn    time.Time
	checkOut   time.Time
	totalPrice float32
	amountPaid float32
}

func NewBooking(session *Session, event BookingCreated) *Booking {
	res := Booking{
		events:     EventStream{session: session},
		bookingId:  event.BookingId,
		checkIn:    event.CheckIn,
		checkOut:   event.CheckOut,
		totalPrice: event.TotalPrice,
		amountPaid: event.AmountPaid,
	}
	res.events.Append(event)

	return &res
}

type BookingCreated struct {
	BookingId  string
	CheckIn    time.Time
	CheckOut   time.Time
	TotalPrice float32
	AmountPaid float32
}

func (s *Booking) PendingAmount() float32 {
	return s.totalPrice - s.amountPaid
}

type BookingSerializer struct{}

func (s *BookingSerializer) Serialize(stream *any) (string, string) {
	return "", ""
}

func (s *BookingSerializer) Deserialize(eventType string, eventData string) any {
	return Booking{}
}

func (s *BookingSerializer) Aggregate(stream *any, event *any) *any {
	if stream == nil {
		res := any(Booking{})
		return &res
	}
	return stream
}
