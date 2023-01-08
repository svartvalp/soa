package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	buckets = []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}

	ResponseTimeHistogram = &prometheus.HistogramVec{}
)

func init() {

	ResponseTimeHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "namespace",
		Name:      "http_server_request_duration_seconds",
		Help:      "Histogram of response time for handler in seconds",
		Buckets:   buckets,
	}, []string{"route", "method", "status_code"})
	prometheus.MustRegister(ResponseTimeHistogram)
}
