package metrics

import (
	"contrib.go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/trace"
)

type TracingService struct {
	Exporter *jaeger.Exporter
}

func NewJaedgerExporter(agentEndpointURI, collectorEndpointURI, serviceName string) (*TracingService, error) {
	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint:     agentEndpointURI,
		CollectorEndpoint: collectorEndpointURI,
		ServiceName:       serviceName,
	})
	if err != nil {
		return nil, err
	}
	trace.RegisterExporter(je)
	return &TracingService{Exporter: je}, nil
}
