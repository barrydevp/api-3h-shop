package main

import (
	"log"
	"net/http"

	"github.com/barrydev/api-3h-shop/src"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/constants"
)

func main() {
	app := src.App{}

	ginEngine := app.NewGinEngine()

	s := &http.Server{
		Addr:    ":" + constants.PORT,
		Handler: ginEngine,
	}

	defer func() {
		log.Println("App shutting down.")
		connections.Mysql.Close()
		s.Close()
	}()

	log.Println("Listening on port ", constants.PORT, "...")

	s.ListenAndServe()

}
