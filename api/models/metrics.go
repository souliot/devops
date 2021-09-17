package models

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"public/libs_go/gateway/metrics/service"
	"public/libs_go/gateway/metrics/system"
)

var (
	DefaultMetrics = new(Metrics)
	Handler        http.Handler
)

type Metrics struct{}

func (m *Metrics) Init(name string) {
	r := prometheus.NewRegistry()
	system.RegisterSystemCollector(r)
	service.RegisterServiceCollector(r, &service.RegisterOptions{name})

	Handler = promhttp.HandlerFor(
		prometheus.Gatherers{r},
		promhttp.HandlerOpts{
			ErrorHandling: promhttp.ContinueOnError,
		},
	)
	return
}
