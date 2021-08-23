package models

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/souliot/gateway/metrics/system"
)

var (
	DefaultMetrics = new(Metrics)
	Handler        http.Handler
)

type Metrics struct{}

func (m *Metrics) Init() {
	r := prometheus.NewRegistry()
	system.RegisterSystemCollector(r)

	Handler = promhttp.HandlerFor(
		prometheus.Gatherers{r},
		promhttp.HandlerOpts{
			ErrorHandling: promhttp.ContinueOnError,
		},
	)
	return
}
