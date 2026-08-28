package rest

import (
	"context"
	"net/http"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/app"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/app/config"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// Server represents an HTTP server with an address
type Server struct {
	logger     ports.Logger
	httpServer *http.Server
}

// NewServer creates and returns a new Server instance, wiring the router to
// the given services so its handlers can execute use cases.
func NewServer(cfg config.AppConfig, logger ports.Logger, services *app.Services) *Server {
	// Create the router and register all routes
	mux := NewRouter(logger, services)

	httpServer := &http.Server{
		Addr:    cfg.App.RestAddr,
		Handler: mux,
	}

	return &Server{
		logger:     logger,
		httpServer: httpServer,
	}
}

// Start begins listening for HTTP requests on the configured address.
func (s *Server) Start() error {
	s.logger.Info("rest server listening", "addr", s.httpServer.Addr)

	err := s.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

// Shutdown stops the server gracefully.
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("shutting down rest server", "addr", s.httpServer.Addr)

	return s.httpServer.Shutdown(ctx)
}
