package utils

// Adapted from https://github.com/axiomhq/variance
// Calculates variance using Welford's online method
// Is able to forget values using the Remove method (by setting a negative weight)
type WelfordStats struct {
	mean float64
	sum  float64
	s    float64
}

func NewWelfordStats() *WelfordStats {
	return new(WelfordStats)
}

func (sts *WelfordStats) Add(x float64) {
	sts.AddWeighted(x, 1)
}

func (sts *WelfordStats) Remove(x float64) {
	sts.AddWeighted(x, -1)
}

func (sts *WelfordStats) AddWeighted(val, weight float64) {
	sts.sum += weight
	meanOld := sts.mean
	sts.mean = meanOld + (weight/sts.sum)*(val-meanOld)
	sts.s = sts.s + weight*(val-meanOld)*(val-sts.mean)
}

func (sts *WelfordStats) Mean() float64 {
	return sts.mean
}

func (sts *WelfordStats) Variance() float64 {
	if sts.sum <= 1 {
		return 0
	}
	return sts.s / (sts.sum - 1)
}
