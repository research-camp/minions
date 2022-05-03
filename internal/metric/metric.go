package metric

import "github.com/prometheus/client_golang/prometheus"

const (
	Namespace = "xerox"
	Subsystem = "proxy"
)

// Metrics has all the client metrics.
type Metrics struct {
	TotalRequests      prometheus.Counter
	SuccessfulRequests prometheus.Counter
	FailedRequests     prometheus.Counter
}
