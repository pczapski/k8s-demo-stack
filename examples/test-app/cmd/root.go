package cmd

import (
	"fmt"
	"os"
	"test-app/logger"
	"test-app/metrics"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var AppVersion string
var GitCommit string

var rootCmd = &cobra.Command{
	Use: "test-app",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

type Services struct {
	metrics metrics.MetricsService
	logger  *zap.Logger
	tracing *metrics.TracingService
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func configureLogger(appVersion string) (*zap.Logger, error) {
	return logger.NewZap(appVersion)
}

func configurePromethuesExporter() (metrics.MetricsService, error) {
	return metrics.NewPrometheusExporter()
}
