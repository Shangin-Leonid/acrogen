package utils /* Utils */

// #
// Ternary operator. The same as "(map[bool]T {true: vTrue, false: vFalse})[cond]".
// #
func TerOp[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	} else {
		return vFalse
	}
}
