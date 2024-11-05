package main

import (
	"log"
	"net/http"
	"time"

	"github.com/mmghobadi/traefik_batchwise/pkg/config"
	"github.com/mmghobadi/traefik_batchwise/pkg/gateway"
	"github.com/mmghobadi/traefik_batchwise/pkg/middleware"
	"github.com/mmghobadi/traefik_batchwise/pkg/models"
)

func main() {

	// Initialize shared channels
	eventChannels := models.NewEventChannels()

	// Initialize middleware
	cfg := config.LoadConfig()
	mw := middleware.NewMiddleware(cfg, eventChannels)

	gateway, err := gateway.NewGateway("http://127.0.0.1:5011/event", eventChannels)
	if err != nil {
		log.Fatal(err)
	}

	mw.Proxy = gateway.Proxy
	go mw.Start()

	server := &http.Server{
		Addr:         ":8050",
		Handler:      gateway,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Printf("Starting API Gateway on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
