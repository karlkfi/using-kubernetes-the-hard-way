package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	serveAddr       = flag.String("addr", "", "Hostname to serve on.")
	servePort       = flag.Int("port", 8081, "Port to serve on.")
	shutdownTimeout = flag.Duration("shutdown_timeout", 1*time.Minute, "Time to wait on SIGTERM.")
)

func main() {
	flag.Parse()

	s := NewServer(*serveAddr, *servePort, *shutdownTimeout)
	stopped := make(chan struct{})

	// Trigger graceful shutdown on SIGINT (Ctrl+C) or SIGTERM (Ctrl+/)
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
		<-sigCh

		if err := s.ShutdownWithTimeout(); err != nil {
			log.Printf("%v", err)
		}
		close(stopped)
	}()

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("%v", err)
	}

	<-stopped
	log.Println("Server stopped gracefully")
}
