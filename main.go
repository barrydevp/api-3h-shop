package main

import (
	"github.com/barrydev/api-3h-shop/src"
	"github.com/barrydev/api-3h-shop/src/connections"
	"log"
	"net/http"
	"os"
)

func main() {
	const DefaultPort string = "4000"
	port := os.Getenv("PORT")

	if port == "" {
		log.Print("$PORT must be set => Using default port: ", DefaultPort)
		port = DefaultPort
	}

	app := src.App{}
	ginEngine := app.NewGinEngine()

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        ginEngine,
	}

	defer func() {
		log.Println("App shutting down.")
		connections.Mysql.Close()
		s.Close()
	}()

	s.ListenAndServe()

	log.Println("Listening on port ", port, "...")
}
