package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
}

const DefaultPort = "2000"

func Server() error {
	// getting PORT env variable
	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}
	// app for setting up http.Handler
	app := Config{}
	// http server configuration
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.Routes(),
	}
	log.Printf("Server up...")
	log.Printf("Listing on %s", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
