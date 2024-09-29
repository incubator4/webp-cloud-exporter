package main

import (
	"github.com/incubator4/webp-cloud-exporter/pkg/server"
	"github.com/incubator4/webp-cloud-exporter/pkg/webpse"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var apiKey string

var command = &cobra.Command{
	Use: "exporter",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := webpse.New(apiKey)
		s := server.New(c)

		return s.Start(8080)
	},
}

func init() {
	command.PersistentFlags().StringVar(&apiKey, "api-key", os.Getenv("WEBP_API_KEY"), "API token for webp-cloud")
}

func main() {
	if err := command.Execute(); err != nil {
		log.Fatal("Failed to execute command", err)
	}
}
