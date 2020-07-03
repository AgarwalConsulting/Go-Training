package service

import (
	"fmt"
	"net/http"
)

type basicHealthImpl struct{}

// HealthInfo implements the first end-point of health service
func (h *basicHealthImpl) HealthInfo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "success")
}

// New returns an instance of type which implements HealthService
func New() HealthService {
	return &basicHealthImpl{}
}
