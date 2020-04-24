package metrics

import (
	"net/http"

	"contrib.go.opencensus.io/exporter/prometheus"
	prom "github.com/prometheus/client_golang/prometheus"
	"go.opencensus.io/stats/view"
)

type MetricsService interface {
	GetExporter() view.Exporter
	GetHandler() func(w http.ResponseWriter, r *http.Request)
}
type prometheusExporter struct {
	exporter *prometheus.Exporter
}

func NewPrometheusExporter() (*prometheusExporter, error) {
	pe, err := prometheus.NewExporter(prometheus.Options{
		Registry: prom.DefaultGatherer.(*prom.Registry),
	})
	if err != nil {
		return nil, err

	}

	view.RegisterExporter(pe)
	return &prometheusExporter{exporter: pe}, nil
}
func (s *prometheusExporter) GetExporter() view.Exporter {
	return s.exporter
}
func (s *prometheusExporter) GetHandler() func(w http.ResponseWriter, r *http.Request) {
	return s.exporter.ServeHTTP
}
