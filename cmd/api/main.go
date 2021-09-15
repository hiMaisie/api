package api

import (
	"log"
	"os"

	"github.com/himaisie/api/pkg/server"
	"github.com/himaisie/api/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg = &config.Config{}
	rootCmd = &cobra.Command{
		Short: "api",
		Run: execute,
	}
)

func init() {
}

func execute(cmd *cobra.Command, args []string) {
	logger := log.New(os.Stdout, "", 0)
	srv := server.New(logger, "0.0.0.0:8080")

	if err := srv.Start(); err != nil {
		logger.Println("Error starting server: ", err)
	}
}
