package meth

import "math"

func Max(ns ...int) int {
	if len(ns) == 1 {
		return ns[0]
	}

	max := math.MinInt
	for _, n := range ns {
		if n > max {
			max = n
		}
	}
	return max
}

func Min(ns ...int) int {
	if len(ns) == 1 {
		return ns[0]
	}

	min := math.MaxInt
	for _, n := range ns {
		if n < min {
			min = n
		}
	}
	return min
}
