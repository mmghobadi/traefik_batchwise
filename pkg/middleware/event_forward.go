package middleware

import (
	"fmt"
	"time"

	"github.com/mmghobadi/traefik_batchwise/pkg/models"
)

// Handler function to route requests
// func (m *Middleware) ForwardHandler(event models.Event) {

// 	metrics.LogEvent(event)

// }

// Handler function to route requests
func (m *Middleware) ForwardEvent(event models.Event) {

	// Set the request headers
	event.Request.Header.Set("X-Event-ID", event.ID)
	event.Request.Header.Set("X-Event-Type", event.Type)
	event.Request.Header.Set("X-API-Received-Time", event.ReceivedTime.Format(time.RFC3339))
	event.Request.Header.Set("X-API-Forwarded-Time", time.Now().Format(time.RFC3339))
	event.Request.Header.Set("X-Event-Urgency", fmt.Sprintf("%f", event.Urgency))
	event.Request.Header.Set("X-Event-System-Priority", fmt.Sprintf("%f", event.Priority))

	m.NextHandler.ServeHTTP(event.Writer, event.Request)

	// // Create a new request
	// client := &http.Client{}
	// req, err := http.NewRequest("GET", "http://127.0.0.1:5011/event", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // Set the request headers
	// req.Header.Set("X-Event-ID", event.ID)
	// req.Header.Set("X-Event-Type", event.Type)
	// req.Header.Set("X-API-Received-Time", event.ReceivedTime.Format(time.RFC3339))
	// req.Header.Set("X-API-Forwarded-Time", time.Now().Format(time.RFC3339))
	// req.Header.Set("X-Event-Urgency", fmt.Sprintf("%f", event.Urgency))
	// req.Header.Set("X-Event-System-Priority", fmt.Sprintf("%f", event.Priority))

	// client.Do(req)

}
