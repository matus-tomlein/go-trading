package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"

	"github.com/matus-tomlein/go-trading/internal/models"
)

// Provides rapid statistical analyses of recent trading data for specified symbols
//
// StatsHandler handles requests for symbol stats
// It returns the stats for the given symbol and k
// It uses a sync.Map to store the stats
// It listens on the processedChannel for new stats
func StatsHandler(processedChannel chan models.SymbolStats) http.HandlerFunc {
	var statsBySymbol sync.Map

	// Listen on the processedChannel for new stats
	go func() {
		for {
			stats := <-processedChannel
			statsBySymbol.Store((models.MapKey{Symbol: stats.Symbol, K: stats.K}), stats)
		}
	}()

	// Return the handler function
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		// Parse the symbol from the URL
		symbol, err := strconv.Atoi(vars["symbol"])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Parse the k parameter from the URL GET params
		kStr := r.URL.Query().Get("k")
		if kStr == "" {
			kStr = "8"
		}
		k, err := strconv.Atoi(kStr)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Check if the symbol and the k is in the map
		stats, ok := statsBySymbol.Load(models.MapKey{Symbol: symbol, K: k})
		if !ok {
			http.Error(w, "Symbol not found", http.StatusNotFound)
			return
		}

		// Respond with the stats
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stats.(models.SymbolStats).Stats)
	}
}
