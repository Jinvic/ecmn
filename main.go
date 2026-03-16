package main

import (
	"fmt"
	"os"

	"ecmn/config"
	"ecmn/handlers"
	"ecmn/logger"
	"ecmn/router"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config: %v\n", err)
		os.Exit(1)
	}

	logger.Init(&cfg.Logging)
	defer logger.Sync()

	logger.Info("Starting application", logger.String("port", fmt.Sprintf("%d", cfg.Server.Port)))

	whHandler := handlers.NewWebhookHandler()
	r := router.Setup(whHandler)

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	if err := r.Run(addr); err != nil {
		logger.Fatal("Failed to start server", logger.Err(err))
	}
}
