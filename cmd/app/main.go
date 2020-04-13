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

	s := NewServer()

	// Trigger graceful shutdown on SIGINT (Ctrl+C) or SIGTERM (Ctrl+/)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
		<-c
		s.Shutdown()
		os.Exit(0)
	}()

	log.Fatal(s.ListenAndServe())
}
