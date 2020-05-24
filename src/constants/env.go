package constants

import (
	"log"
	"os"
)

var (
	PRIMARY_HOST         string
	WEB_HOST             string
	ADMIN_HOST           string
	PORT                 string
	APP_ENV              string
	CLEARDB_DATABASE_URL string
	SECRET_KEY           string
)

func init() {
	PRIMARY_HOST = os.Getenv("PRIMARY_HOST")

	if PRIMARY_HOST == "" {
		PRIMARY_HOST = "localhost"
		log.Print("$PRIMARY_HOST must be set => Using default PRIMARY_HOST: ", "localhost")
	}

	WEB_HOST = os.Getenv("WEB_HOST")

	if WEB_HOST == "" {
		WEB_HOST = "http://localhost:3000"
		log.Print("$WEB_HOST must be set => Using default WEB_HOST: ", "http://localhost:3000")
	}

	ADMIN_HOST = os.Getenv("ADMIN_HOST")

	if ADMIN_HOST == "" {
		ADMIN_HOST = "http://localhost:3000"
		log.Print("$ADMIN_HOST must be set => Using default ADMIN_HOST: ", "http://localhost:3000")
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
		// CLEARDB_DATABASE_URL = "mysql://root:barry123456@localhost/3hshop"
		log.Print("CLEARDB_DATABASE_URL: ", CLEARDB_DATABASE_URL)
	}

	SECRET_KEY = os.Getenv("SECRET_KEY")

	if SECRET_KEY == "" {
		SECRET_KEY = "api3hshop"
		log.Print("SECRET_KEY: ", SECRET_KEY)
	}
}
