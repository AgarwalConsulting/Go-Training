package service

import (
	"net/http"
)

// HealthService represents a Service which can show if a service is up or not
type HealthService interface {
	HealthInfo(http.ResponseWriter, *http.Request)
}
