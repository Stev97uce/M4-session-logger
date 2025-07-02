package main

import (
	"log"
	"net/http"
	"os"

	"session-logger/app/db"
	"session-logger/app/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectMongo()
	defer db.CloseMongo()

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	apiPort := os.Getenv("API_PORT")
	log.Println("Session Logger running on port", apiPort)

	log.Fatal(http.ListenAndServe(":"+apiPort, r))
}
