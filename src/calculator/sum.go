package calculator

// Private
var lastResult int

// Public
var Version = "2.0"

// Somme de deux entiers
func Sum(number1 int, number2 int) int {
	lastResult = number1 + number2
	return lastResult
}
