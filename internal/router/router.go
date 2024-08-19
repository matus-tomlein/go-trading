package router

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/matus-tomlein/go-trading/internal/handlers"
	"github.com/matus-tomlein/go-trading/internal/models"
	"github.com/matus-tomlein/go-trading/internal/utils"
)

// inBatchChannel is a channel for incoming batches
// that need to be processed
// It has a buffer of 1000 to prevent blocking
var inBatchChannel = make(chan models.Batch, 1000)

// outSymbolStatsChannel is a channel for processed symbol stats
// It is consumed by the StatsHandler
var outSymbolStatsChannel = make(chan models.SymbolStats)

// ProcessChannelMessages processes messages from the inBatchChannel
// and sends the results to the outSymbolStatsChannel
func ProcessChannelMessages() {
	symbolStatsComputers := make(map[models.MapKey]*utils.SymbolStatsComputer)

	for {
		batch := <-inBatchChannel
		if len(batch.Values) == 0 {
			continue
		}

		fmt.Println("Processing batch", batch)

		for k := 1; k <= 8; k++ {
			mapKey := models.MapKey{Symbol: batch.Symbol, K: k}

			computer, exists := symbolStatsComputers[mapKey]
			if !exists {
				computer = utils.NewSymbolStatsComputer(batch.Symbol, k)
				symbolStatsComputers[mapKey] = computer
			}

			for _, value := range batch.Values {
				computer.AddValue(value)
			}

			outSymbolStatsChannel <- computer.GetSymbolStats()
		}

		fmt.Println("Batch processed")
	}
}

// Sets up the API routes, starts the message processing goroutine
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Add the routes
	r.HandleFunc("/add_batch", handlers.AddBatchHandler(inBatchChannel)).Methods("POST")
	r.HandleFunc("/stats/{symbol}", handlers.StatsHandler(outSymbolStatsChannel)).Methods("GET")

	// Start the message processing async goroutine
	go ProcessChannelMessages()

	return r
}
