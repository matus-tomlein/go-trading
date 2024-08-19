package utils

import (
	"math"

	"github.com/matus-tomlein/go-trading/internal/models"
	"github.com/smjure/movingminmax"
)

// SymbolStatsComputer computes statistics for a symbol and a given window size.
// The stats are computed using online algorithms.
// Makes use of the Welford algorithm for computing the mean and variance.
// The min and max are computed using the movingminmax package which is using the algorithm in https://arxiv.org/abs/cs/0610046
type SymbolStatsComputer struct {
	Symbol       int
	K            int
	History      []float64
	Welford      *WelfordStats
	MovingMinMax *movingminmax.MovingMinMax
}

// Creates a new SymbolStatsComputer.
func NewSymbolStatsComputer(symbol int, k int) *SymbolStatsComputer {
	stats := new(SymbolStatsComputer)
	stats.Symbol = symbol
	stats.K = k
	stats.History = make([]float64, 0)
	stats.Welford = NewWelfordStats()
	stats.MovingMinMax = movingminmax.NewMovingMinMax(
		uint(math.Pow10(k)),
	)
	return stats
}

// Returns the size of the window that the stats are computed from.
func (s *SymbolStatsComputer) WindowSize() int {
	return int(math.Pow10(s.K))
}

// Updates the stats with a new value.
func (s *SymbolStatsComputer) AddValue(value float64) {
	s.History = append(s.History, value)
	s.Welford.Add(value)
	s.MovingMinMax.Update(value)
	if len(s.History) > s.WindowSize() {
		s.Welford.Remove(s.History[0])
		s.History = s.History[1:]
	}
}

// Returns the computed statistics.
func (s *SymbolStatsComputer) GetStats() models.Stats {
	if len(s.History) == 0 {
		return models.Stats{}
	}
	return models.Stats{
		Min:  s.MovingMinMax.Min(),
		Max:  s.MovingMinMax.Max(),
		Avg:  s.Welford.Mean(),
		Last: s.History[len(s.History)-1],
		Var:  s.Welford.Variance(),
	}
}

func (s *SymbolStatsComputer) GetSymbolStats() models.SymbolStats {
	return models.SymbolStats{
		Symbol: s.Symbol,
		K:      s.K,
		Stats:  s.GetStats(),
	}
}
