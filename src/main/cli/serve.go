package cli

import (
	"farm/src/main/env"
	"github.com/spf13/cobra"
	"log"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serving farm service.",
	Long:  `Serving farm service.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() {
	go env.RunInsecureGRPCServer()
	err := env.RunInsecureHTTPServer()
	if err != nil {
		log.Fatal(err)
	}
}
