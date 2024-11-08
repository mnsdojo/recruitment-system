package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mnsdojo/recruitment-system/config"
	"github.com/mnsdojo/recruitment-system/db"
	"github.com/mnsdojo/recruitment-system/routes"
)

func main() {
	cfg := config.NewConfig()
	database, err := db.InitializeDb(cfg)
	if err != nil {
		log.Fatalf("Could not set up the database: %v", err)
	}

	router := mux.NewRouter()
	routes.SetupRoutes(router, database)
	port := cfg.Port
	log.Printf("Server is starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
