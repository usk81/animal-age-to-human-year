package dog

import (
	"math"
)

// Calc dog age to human year
//   formula: 16ln(dog age) + 31 = human year
func Calc(age float64) int {
	if age <= 0 {
		return 0
	}
	return int(16*math.Log(age) + 31)
}

// WrongCalc calcates dog age to human year.
// This formula is popular, but it is wrong
//   formula: 7 * dog age = human year
func WrongCalc(age float64) int {
	if age <= 0 {
		return 0
	}
	return int(7 * age)
}
