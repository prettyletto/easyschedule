package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	database "github.com/prettyletto/easyschedule/db"
	"github.com/prettyletto/easyschedule/internal/server"
)

func waitForSignalToStop(s *server.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting Down the application...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Stop(ctx); err != nil {
		log.Fatalf("Server Shutdown failed: %v", err)
	}

}

func main() {

	db := database.Init()
	s := server.New(":8080", db)

	go func() {
		if err := s.Start(); err != nil && err.Error() != "http: Server closed" {
			log.Println("Server could not start properly: ", err)
		}
	}()

	waitForSignalToStop(s)
}
