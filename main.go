package main

import (
	"log"
	"net/http"
	"time"

	"countryinfo/confirm"
	"countryinfo/handles"
	"countryinfo/services"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func main() {

	//loads confirm
	port := confirm.Reconfirm().Port

	// starts up services
	countryService := services.NewCountryService()
	hand := handles.NewHandler(countryService)

	// Define routes usinf serveMux
	mux := http.NewServeMux()
	mux.HandleFunc("/countryinfo/v1/info/", hand.HandleCountryInfo)
	mux.HandleFunc("/countryinfo/v1/population/", hand.HandlePopulation)
	mux.HandleFunc("/countryinfo/v1/status/", hand.HandleStatus)

	// Starts server
	addr := ":" + port
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

