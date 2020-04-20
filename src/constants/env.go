package constants

import (
	"log"
	"os"
)

var (
	PRIMARY_HOST string
	PORT string
	APP_ENV string
	CLEARDB_DATABASE_URL string
)

func init() {
	PRIMARY_HOST = os.Getenv("PRIMARY_HOST")

	if PRIMARY_HOST == "" {
		PRIMARY_HOST = "localhost"
		log.Print("$PRIMARY_HOST must be set => Using default PRIMARY_HOST: ", "localhost")
	}

	PORT = os.Getenv("PORT")

	if PORT == "" {
		PORT = "4000"
		log.Print("$PORT must be set => Using default PORT: ", PORT)
	}

	APP_ENV = os.Getenv("APP_ENV")

	if APP_ENV == "" {
		APP_ENV = "development"
		log.Print("APP_ENV: ", APP_ENV)
	}

	CLEARDB_DATABASE_URL = os.Getenv("CLEARDB_DATABASE_URL")

	if CLEARDB_DATABASE_URL == "" {
		CLEARDB_DATABASE_URL = "mysql://b08738ff9fff5e:e79a1d81@us-cdbr-iron-east-01.cleardb.net/heroku_e16926abf051efd?reconnect=true"
		log.Print("CLEARDB_DATABASE_URL: ", CLEARDB_DATABASE_URL)
	}
}