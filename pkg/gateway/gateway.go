package gateway

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/mmghobadi/traefik_batchwise/pkg/models"
)

type Gateway struct {
	Target       *url.URL
	Proxy        *httputil.ReverseProxy
	Logger       *log.Logger
	EventChannel *models.EventChannels
}

func NewGateway(targetURL string, eventChannels *models.EventChannels) (*Gateway, error) {

	target, err := url.Parse(targetURL)
	if err != nil {
		return nil, err
	}

	return &Gateway{
		Target:       target,
		Proxy:        httputil.NewSingleHostReverseProxy(target),
		Logger:       log.New(os.Stdout, "API Gateway: ", log.LstdFlags),
		EventChannel: eventChannels,
	}, nil
}

// Custom function to handle proxying
func (g *Gateway) HandleProxyRequest(w http.ResponseWriter, r *http.Request) {
	// Add custom headers
	r.Header.Add("X-Forwarded-Host", r.Host)
	r.Header.Add("X-Origin-Host", g.Target.Host)

	// Proxy the request
	g.Proxy.ServeHTTP(w, r)
}

func (g *Gateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	// Parse event urgency from request
	eventUrgency, _ := strconv.ParseFloat(r.Header.Get("X-Event-Urgency"), 64)

	// Parse event from request (simplified)
	event := models.Event{
		ID:           r.Header.Get("X-Event-ID"),
		Type:         r.Header.Get("X-Event-Type"),
		Urgency:      eventUrgency,
		ReceivedTime: time.Now(),
		// Payload: r.Body,
		Request: r,
	}

	// Add event to the input channel
	g.EventChannel.EventInput <- event

	// Call the custom proxy function
	// g.HandleProxyRequest(w, r)

	// Log the request
	g.Logger.Printf(
		"Proxied request: %s %s -> %s [urgent: %v] [total time: %v]",
		r.Method,
		r.URL.Path,
		g.Target.String(),
		r.Header.Get("X-Urgent") != "",
		time.Since(startTime),
	)
}
