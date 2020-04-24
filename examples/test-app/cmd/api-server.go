package cmd

import (
	"fmt"
	"os"
	"test-app/http"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(apiServerCmd)
}

var apiServerCmd = &cobra.Command{
	Use: "api-server",
	Run: func(cmd *cobra.Command, args []string) {
		startApiServer()
	},
}

func startApiServer() {
	logger, err := configureLogger(AppVersion)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer logger.Sync()
	pe, err := configurePromethuesExporter()
	if err != nil {
		logger.Fatal(err.Error())
	}

	s, err := http.New(logger, pe)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	s.Run(port, AppVersion)
}
