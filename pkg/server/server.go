package server

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"golang.fiber.template/pkg/config"
)

func Start(a *fiber.App) {
	StartServerWithGracefulShutdown(a)
}

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	_config := config.Server()
	if err := a.Listen(":" + _config.Port); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	// Run server.
	_config := config.Server()
	if err := a.Listen(":" + _config.Port); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
