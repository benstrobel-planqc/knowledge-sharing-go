package mycoollib

// Functions with a lowercase letter are not exported
func divideWithRemainder(dividend int, divisor int) (int, int) {
	return dividend / divisor, dividend % divisor
}

// Functions with a cased letter are exported
func Divide(dividend float64, divisor float64) float64 {
	return dividend / divisor
}
