package models

import (
	"net/http"
	"time"
)

type Event struct {
	ID      string
	Type    string
	Urgency float64
	// Payload  interface{}
	Priority            float64
	ReceivedTime        time.Time
	HoldingTime         time.Duration
	CompletedTime       time.Time
	IsSysteHighPriority bool
	IsUserHighPriority  bool
	Request             *http.Request
	Writer              http.ResponseWriter
}

type Batch struct {
	Events []Event
}

type EventChannels struct {
	EventInput chan Event
}

func NewEventChannels() *EventChannels {
	return &EventChannels{
		EventInput: make(chan Event, 1000),
	}
}
