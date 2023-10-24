package test

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"testing"
)

func TestPrometheus(t *testing.T) {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":2112", nil); err != nil {
		return
	}
}
