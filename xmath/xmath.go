package xmath

// MinInt returns the smaller of x or y.
func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// MaxInt returns the gratest of x or y.
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
