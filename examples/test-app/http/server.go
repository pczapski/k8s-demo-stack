package http

import (
	"fmt"
	"net/http"
	"test-app/metrics"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.uber.org/zap"
)

type server struct {
	e      *gin.Engine
	logger *zap.Logger
	pe     metrics.MetricsService
}

func New(logger *zap.Logger, pe metrics.MetricsService) (*server, error) {
	r := gin.New()
	view.RegisterExporter(pe.GetExporter())

	// Register stat views
	err := view.Register(
		// Gin (HTTP) stats
		ochttp.ServerRequestCountView,
		ochttp.ServerRequestBytesView,
		ochttp.ServerResponseBytesView,
		ochttp.ServerLatencyView,
		ochttp.ServerRequestCountByMethod,
		ochttp.ServerResponseCountByStatusCode,
	)
	if err != nil {
		return nil, err

	}
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	return &server{e: r, pe: pe}, nil
}

func (s *server) Run(port, version string) error {
	s.setupRouter(version)
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", port),
		&ochttp.Handler{
			Handler: s.e,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *server) Close() {
}

func (s *server) setupRouter(version string) {
	s.e.GET("/", homeHandler(version))
	s.e.GET("/metrics", metricsHandler(s.pe))
}
