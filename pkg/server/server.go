package server

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"golang.fiber.template/domain/model"
)

func Start(a *fiber.App, cf *model.Config) {
	StartServerWithGracefulShutdown(a, cf)
}

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App, cf *model.Config) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	if err := a.Listen(":" + cf.APP.PORT); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

func StartServer(a *fiber.App, cf *model.Config) {
	if err := a.Listen(":" + cf.APP.PORT); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
