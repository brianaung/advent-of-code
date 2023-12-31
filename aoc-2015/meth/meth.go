package meth

import "math"

func MinInt(nums ...int) int {
	min := math.MaxInt
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

func MaxInt(nums ...int) int {
	max := math.MinInt
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}
