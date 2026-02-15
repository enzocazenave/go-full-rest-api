package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/enzocazenave/golang-full-rest-api/dbconfig"
	"github.com/enzocazenave/golang-full-rest-api/internal/handlers"
	"github.com/enzocazenave/golang-full-rest-api/internal/routes"
	"github.com/enzocazenave/golang-full-rest-api/serverconfig"
)

func main() {
	config, err := serverconfig.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db := dbconfig.ConnectDB(config.DatabaseURL)
	defer db.Close()

	handler := handlers.NewHandlers()

	mux := http.NewServeMux()

	routes.SetupRoutes(mux, handler)

	serverAddr := fmt.Sprintf(":%s", config.ServerPort)
	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}

	fmt.Printf(
		"Server is running on %s in %s environment with log level %s\n",
		serverAddr,
		config.Environment,
		config.LogLevel,
	)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
