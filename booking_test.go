package glimpse

import (
	"encoding/json"
	"errors"
	"time"
)

type Booking struct {
	events *EventStream

	bookingId  string
	checkIn    time.Time
	checkOut   time.Time
	totalPrice float32
	amountPaid float32
}

func NewBooking(session Session, event *BookingCreated) *Booking {
	res := Booking{
		events:     NewEventStream(session, "booking", event.BookingId),
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

type PaymentRegistered struct {
	Date   time.Time
	Amount float32
}

func (s *Booking) RegisterPayment(event *PaymentRegistered) {
	s.amountPaid += event.Amount
	s.events.Append(event)
}

func (s *Booking) PendingAmount() float32 {
	return s.totalPrice - s.amountPaid
}

type BookingSerializer struct{}

func (s *BookingSerializer) Serialize(event any) (string, string, error) {
	switch event.(type) {
	case *BookingCreated:
		json, err := json.Marshal(event)
		if err != nil {
			return "", "", err
		}

		return "BookingCreated", string(json), nil
	}

	return "", "", errors.New("Unknown event type")
}

func (s *BookingSerializer) Deserialize(eventType string, eventData string) (any, error) {
	switch eventType {
	case "BookingCreated":
		var event BookingCreated
		err := json.Unmarshal([]byte(eventData), &event)
		if err != nil {
			return nil, err
		}

		return &event, nil
	}

	return nil, errors.New("Unknown event type")
}

func (s *BookingSerializer) Aggregate(session Session, stream any, untyped any) any {
	switch event := untyped.(type) {
	case *BookingCreated:
		return NewBooking(session, event)
	}

	return stream
}
