package models

// Computed statistics for a symbol and window size
type Stats struct {
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
	Avg  float64 `json:"avg"`
	Last float64 `json:"last"`
	Var  float64 `json:"var"`
}
