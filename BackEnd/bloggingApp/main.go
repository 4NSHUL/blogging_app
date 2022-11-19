package main

import (
	"blog/routing"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Load Configurations from config.json using Viper
	LoadAppConfig()
	// Initialize Database
	Connect(AppConfig.ConnectionString)

	// Initialize the router
	router := routing.GetRouter()
	// Register Routes
	RegisterRoutes(router)
	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))

}
