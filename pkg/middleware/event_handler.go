package middleware

import (
	"log"
	"time"
)

func (m *Middleware) eventHandler() {
	for {
		select {
		case event := <-m.EventInput:
			// Duration limiting logic
			if m.LastRequestTime.Sub(m.FirstRequestTime) >= time.Duration(5)*time.Minute {
				log.Fatal("Total Request: ", m.RequestCount)
			}

			// System metrics (should be collected from monitoring tools)
			systemLoad := getSystemLoad()
			traffic := getTrafficPatterns()

			// Classify event
			m.classifyEvent(&event, systemLoad, traffic)

			// Add event to appropriate queue
			if event.Priority >= m.Config.Thresholds.Priority {
				event.IsSysteHighPriority = true
				m.HighPriorityQueue <- event
			} else {
				m.LowPriorityQueue <- event
			}
		case <-m.StopChan:
			return
		}
	}
}
