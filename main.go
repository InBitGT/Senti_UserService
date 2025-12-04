package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"UserService/db"
	"UserService/internal/config"
	"UserService/internal/server"

	"github.com/joho/godotenv"
)

func main() {

	config.Init()
	database := db.Database()

	// migration.Migration()

	handlers := server.InitializeHandlers(database)

	srv := server.NewServer(database, handlers)

	_ = godotenv.Load()

	httpServer := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      srv.Router,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
	}

	log.Fatal(httpServer.ListenAndServe())
}
