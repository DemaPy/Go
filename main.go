package main

import (
	"log"
	"net/http"

	"go_dev/campaign"
)

func main() {
	http.HandleFunc("/campaign", campaign.GetCampaign)

	// Start the server on port 8080
	log.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
