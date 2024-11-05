package traefik_batchwise

import (
	"context"
	"net/http"
	"time"

	"github.com/mmghobadi/traefik_batchwise/pkg/config"
	"github.com/mmghobadi/traefik_batchwise/pkg/middleware"
	"github.com/mmghobadi/traefik_batchwise/pkg/models"
)

// New creates a new Middleware instance.
func New(ctx context.Context, next http.Handler) (http.Handler, error) {
	// Load configuration
	cfg := config.LoadConfig("configs/config.yaml")
	m := &middleware.Middleware{
		NextHandler:       next,
		Config:            cfg,
		EventInput:        make(chan models.Event, 1000),
		HighPriorityQueue: make(chan models.Event, 1000),
		LowPriorityQueue:  make(chan models.Event, 10000),
		BatchQueue:        make(chan models.Batch, 100),
		StopChan:          make(chan bool),
		LastBatchTime:     time.Now(),
	}
	m.Start()
	return m, nil
}
