package models

// Wrapper around Stats to include symbol and window size
type SymbolStats struct {
	Symbol int   `json:"symbol"`
	K      int   `json:"k"`
	Stats  Stats `json:"stats"`
}
