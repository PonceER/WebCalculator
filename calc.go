package calc

// Add function to get sum of 2 integers
func Add(addend1, addend2 int) int {
	return addend1 + addend2
}

// Subtract function to get difference of 2 integers
func Subtract(minuend, subtrahend int) int {
	return minuend - subtrahend
}

// Multiply function to get product of 2 integers
func Multiply(multiplicant, multiplier int) int {
	return multiplicant * multiplier
} 

// Divide function to get quotient of 2 integers
func Divide(dividend, divisor int) float64 {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	return float64(dividend) / float64(divisor)
}
