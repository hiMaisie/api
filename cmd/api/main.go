package api

import (
	"github.com/himaisie/api/pkg/server"
	"log"
	"os"
)

func Execute() {
	logger := log.New(os.Stdout, "", 0)
	srv := server.New(logger, "0.0.0.0:8080")

	if err := srv.Start(); err != nil {
		logger.Println("Error starting server: ", err)
	}
}
