package algo /* Algorithms */

// # CalcFactorial evaluates factorial of number.
func CalcFactorial(n uint) uint {
	var fact uint = 1
	for n > 1 {
		fact *= n
		n--
	}
	return fact
}
