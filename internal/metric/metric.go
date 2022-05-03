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

func newCounter(counterOpts prometheus.CounterOpts) prometheus.Counter {
	ev := prometheus.NewCounter(counterOpts)

	if err := prometheus.Register(ev); err != nil {
		panic(err)
	}

	return ev
}

func NewMetrics() Metrics {
	return Metrics{
		TotalRequests: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "total_requests",
			Help:        "total number of requests to proxy server",
			ConstLabels: nil,
		}),
		SuccessfulRequests: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "successful_requests",
			Help:        "total number of successful requests",
			ConstLabels: nil,
		}),
		FailedRequests: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "failed_requests",
			Help:        "total number of failed requests",
			ConstLabels: nil,
		}),
	}
}
