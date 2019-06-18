package monitoring

import (
	"net/http"

	"github.com/Shopify/airbrake-go"
)

// AirbrakeConfig contains airbrake configurations
type AirbrakeConfig struct {
	APIKey      string
	Endpoint    string
	Environment string
}

// PanicNotifier contains airbrake notifier methods
type PanicNotifier interface {
	Capture(r *http.Request)
}

// Capture captures panic
func (ac *AirbrakeConfig) Capture(r *http.Request) {
	airbrake.CapturePanic(r)
}
