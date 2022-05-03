package metric

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	Srv     *http.ServeMux
	Address string
}

// NewServer creates a new monitoring server.
func NewServer(cfg Config) Server {
	var srv *http.ServeMux

	if cfg.Enable {
		srv = http.NewServeMux()
		srv.Handle("/metrics", promhttp.Handler())
	}

	return Server{
		Address: cfg.Host,
		Srv:     srv,
	}
}

// Start creates and run a metric server for prometheus in new go routine.
func (s Server) Start() {
	if s.Srv == nil {
		return
	}

	go func() {
		if err := http.ListenAndServe(s.Address, s.Srv); err != nil {
			panic(err)
		}
	}()
}
