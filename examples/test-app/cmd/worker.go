package cmd

import (
	"context"
	"fmt"
	"os"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(workerrCmd)
}

type Receiver struct {
	client cloudevents.Client
	Target string
}

var workerrCmd = &cobra.Command{
	Use: "worker",
	Run: func(cmd *cobra.Command, args []string) {
		startWorker()
	},
}

func startWorker() {
	logger, err := configureLogger(AppVersion)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer logger.Sync()
	logger.Info("start")
	client, err := cloudevents.NewDefaultClient()
	if err != nil {
		logger.Fatal(err.Error())
	}

	if err := client.StartReceiver(context.Background(), display); err != nil {
		logger.Fatal(err.Error())
	}

}

func display(event cloudevents.Event) {
	fmt.Printf("☁️  cloudevents.Event\n%s", event.String())
}
