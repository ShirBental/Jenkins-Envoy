package main

import (
	"log"
	nethttp "net/http"
	"time"

	"github.com/shirbental/jenkins-envoy/internal/config"
	httpserver "github.com/shirbental/jenkins-envoy/internal/http"
	"github.com/shirbental/jenkins-envoy/internal/jenkins"
)

func main() {
	cfg := config.Load()

	client := jenkins.NewMockClient() // switch to real later
	server := httpserver.NewServer(client)

	s := &nethttp.Server{
		Addr:         ":" + cfg.Port,
		Handler:      server.Routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Jenkins Envoy listening on port %s", cfg.Port)
	log.Fatal(s.ListenAndServe())
}
