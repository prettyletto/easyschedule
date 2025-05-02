package server

import (
	"context"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Server struct {
	httpServer *http.Server
	db         *sqlx.DB
}

func New(addr string, db *sqlx.DB) *Server {

	mux := http.NewServeMux()
	RegisterRoutes(mux, db)

	return &Server{httpServer: &http.Server{
		Addr:    addr,
		Handler: mux,
	},
		db: db,
	}
}

func (s *Server) Start() error {
	log.Printf("Starting server at %s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	log.Println("Stopping server...")
	return s.httpServer.Shutdown(ctx)
}
