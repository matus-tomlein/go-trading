package models

// Batch represents a batch of values for a symbol
// It is used in the /add_batch endpoint request body
type Batch struct {
	// The financial instrument's identifier
	Symbol int `json:"symbol"`
	// Array of floating-point numbers representing sequential trading prices
	Values []float64 `json:"values"`
}
