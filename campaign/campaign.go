package campaign

import (
	"log"
	"net/http"

	"go_dev/campaign/model"
	"go_dev/utils"
)

func GetCampaign(w http.ResponseWriter, r *http.Request) {
	campaign := model.Get("My campaign")

	jsonResponse, err := utils.ResponseDataToJson(campaign)
	if err != nil {
		log.Printf("Error marshalling response: %v", err)

		// Prepare an error response in JSON format
		jsonError, err := utils.ToJson(utils.ResponseDataError(err))
		if err != nil {
			// If there's an error while preparing the error response, handle it
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set the Content-Type and status code for the error response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500 status code for errors
		w.Write(jsonError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
