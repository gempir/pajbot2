package helper

import "math"

/*
Sum returns the sum of the given slice of ints
*/
func Sum(s []int) int {
	var x int
	for i := range s {
		x += s[i]
	}
	return x
}

// Round xD
func Round(val float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
