package main

type PaymentEvent struct {
	EventType string `json:"event_type"`
	PaymentId string `json:"payment_id"`
	Status    string `json:"status"`
}
