package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Content struct {
	Key string
}

// Server wraps http.Server, handles multiple paths, and handles graceful
// shutdown.
type Server struct {
	shutdownTimeout time.Duration
	shutdown        bool
	server          *http.Server
	port            int
}

// NewServer wraps a new http.Server and registers all handlers.
func NewServer(host string, port int, shutdownTimeout time.Duration) *Server {
	mux := http.NewServeMux()

	s := &Server{
		shutdownTimeout: shutdownTimeout,
		server: &http.Server{
			Handler:      mux,
			Addr:         fmt.Sprintf("%s:%d", host, port),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
	}

	mux.HandleFunc("/", s.rootHandler)
	mux.HandleFunc("/healthz", s.healthzHandler)
	mux.HandleFunc("/content.json", s.jsonHandler)

	return s
}

func (s *Server) Handler() http.Handler {
	return s.server.Handler
}

// ListenAndServe wraps http.Server.ListenAndServe
// No error if server closed gracefully.
func (s *Server) ListenAndServe() error {
	err := s.server.ListenAndServe()
	// http.Server.ListenAndServe always returns err != nil
	if err != http.ErrServerClosed {
		return fmt.Errorf("ListenAndServe Errored: %v\n", err)
	}
	return nil
}

// Shutdown triggers healthz to return unhealthy and shuts down the httpServer
// with the configured timeout.
// In shutdown mode, the server will still respond to requests, but the
// unhealthy /healthz will eventually cause the load balancer to remove this app
// from the backend pool.
func (s *Server) ShutdownWithTimeout() error {
	s.shutdown = true

	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	err := s.server.Shutdown(ctx) // blocks until shut down is complete
	if err == context.Canceled {
		return fmt.Errorf("Graceful Shutdown Timed Out: %v", err)
	} else if err != nil {
		return fmt.Errorf("Graceful Shutdown Errored: %v\n", err)
	}
	return nil
}

// healthzHandler responds as healthy when the app instance is ready to recieve
// traffic.
// healthzHandler implements http.Handler.
func (s *Server) healthzHandler(w http.ResponseWriter, r *http.Request) {
	if s.shutdown {
		w.WriteHeader(http.StatusServiceUnavailable) // 503
		fmt.Fprintf(w, "shutdown mode")
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "healthy")
}

// rootHandler serves an HTTP page.
// rootHandler implements http.HandleFunc.
func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	w.Header().Set("Expires", "0")                                         // Proxi
	w.Header().Set("Kubernetes-Pod-Name", os.Getenv("KUBERNETES_POD_NAME"))
	fmt.Fprintf(w, "Hello World\n")
}

// jsonHandler serves a JSON file.
// jsonHandler implements http.HandleFunc.
func (s *Server) jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	w.Header().Set("Expires", "0")                                         // Proxi
	w.Header().Set("Kubernetes-Pod-Name", os.Getenv("KUBERNETES_POD_NAME"))
	w.WriteHeader(http.StatusOK)

	body := &Content{
		Key: "value",
	}
	json.NewEncoder(w).Encode(body)
}
