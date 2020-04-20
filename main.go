package main

import (
	"github.com/barrydev/api-3h-shop/src"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/constants"
	"log"
	"net/http"
)

func main() {
	app := src.App{}
	ginEngine := app.NewGinEngine()

	s := &http.Server{
		Addr:           ":" + constants.PORT,
		Handler:        ginEngine,
	}

	defer func() {
		log.Println("App shutting down.")
		connections.Mysql.Close()
		s.Close()
	}()

	s.ListenAndServe()

	log.Println("Listening on port ", constants.PORT, "...")
}
