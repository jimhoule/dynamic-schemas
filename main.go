package main

import (
	"fmt"
	"log"
	"main/collections"
	"main/database"
	"main/database/arango"
	"main/documents"
	"main/router"
	"main/schemas"
	"main/swagger"
	"main/tenants"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	clientDbHandler := database.NewDbHandler(
		arango.NewArangoDb(),
		os.Getenv("CLIENT_DB_NAME"),
		os.Getenv("CLIENT_DB_ADDRESS"),
		os.Getenv("CLIENT_DB_USERNAME"),
		os.Getenv("CLIENT_DB_PASSWORD"),
	)

	appDbHandler := database.NewDbHandler(
		arango.NewArangoDb(),
		"",
		os.Getenv("APP_DB_ADDRESS"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
	)

	mainRouter := router.Get()

	tenants.Init(mainRouter, clientDbHandler)
	schemas.Init(mainRouter, appDbHandler)
	collections.Init(mainRouter, appDbHandler)
	documents.Init(mainRouter, appDbHandler)
	swagger.Init(mainRouter)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", os.Getenv("HTTP_URL"), os.Getenv("HTTP_PORT")),
		Handler: mainRouter,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
