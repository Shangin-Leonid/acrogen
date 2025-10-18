package algo /* Algorithms */

func CalcFactorial(n uint) uint {
	var fact uint = 1
	for n > 1 {
		fact *= n
		n--
	}
	return fact
}
