package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/matus-tomlein/go-trading/internal/models"
)

// Allows the bulk addition of consecutive trading data points for a specific symbol
//
// AddBatchHandler creates a handler function for a POST request
// It decodes the JSON request body into a Batch struct
// It sends the Batch struct to the incomingChannel
func AddBatchHandler(incomingChannel chan models.Batch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the JSON request body into the Batch struct
		var batch models.Batch
		err := json.NewDecoder(r.Body).Decode(&batch)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Send the Batch struct to the incomingChannel
		incomingChannel <- batch

		// Respond with a success message
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"status": "success",
		})
	}
}
