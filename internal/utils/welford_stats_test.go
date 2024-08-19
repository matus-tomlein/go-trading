package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariance(t *testing.T) {
	suite := assert.New(t)

	stats := NewWelfordStats()
	stats.Add(1)
	stats.Add(1)
	stats.Add(1)
	stats.Add(0)
	stats.Add(0)
	stats.Add(0)

	suite.EqualValues(0.3, stats.Variance())
}

func TestMean(t *testing.T) {
	suite := assert.New(t)

	stats := NewWelfordStats()
	stats.Add(1)
	stats.Add(1)
	stats.Add(1)
	stats.Add(0)
	stats.Add(0)
	stats.Add(0)

	suite.EqualValues(0.5, stats.Mean())
}

func TestRemove(t *testing.T) {
	suite := assert.New(t)

	stats := NewWelfordStats()
	stats.Add(3491)
	stats.Add(10921)
	stats.Add(12)
	stats.Add(4210)
	stats.Add(351)
	stats.Add(15291)

	variance := stats.Variance()
	mean := stats.Mean()

	stats.Add(239)
	stats.Remove(239)

	suite.EqualValues(mean, stats.Mean())
	suite.EqualValues(variance, stats.Variance())

}
