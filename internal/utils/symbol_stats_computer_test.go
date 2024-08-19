package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleValue(t *testing.T) {
	suite := assert.New(t)

	stats := NewSymbolStatsComputer(1, 1)
	stats.AddValue(1)

	result := stats.GetStats()

	suite.EqualValues(1, result.Min)
	suite.EqualValues(1, result.Max)
	suite.EqualValues(1, result.Avg)
	suite.EqualValues(1, result.Last)
	suite.EqualValues(0, result.Var)
}

func TestTwoValues(t *testing.T) {
	suite := assert.New(t)

	stats := NewSymbolStatsComputer(1, 2)
	stats.AddValue(1)
	stats.AddValue(2)

	result := stats.GetStats()

	suite.EqualValues(1, result.Min)
	suite.EqualValues(2, result.Max)
	suite.EqualValues(1.5, result.Avg)
	suite.EqualValues(2, result.Last)
	suite.EqualValues(0.5, result.Var)
}

func TestOverWindowSize(t *testing.T) {
	suite := assert.New(t)

	stats := NewSymbolStatsComputer(1, 0)
	stats.AddValue(1)
	stats.AddValue(2)

	result := stats.GetStats()

	suite.EqualValues(2, result.Min)
	suite.EqualValues(2, result.Max)
	suite.EqualValues(2, result.Avg)
	suite.EqualValues(2, result.Last)
	suite.EqualValues(0, result.Var)
}
