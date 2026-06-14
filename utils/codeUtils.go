package utils /* Utils */

// # TerOp implements ternary operator.
//
// # Params:
//
//   - boolean flag
//   - value if true
//   - value if false
//
// # Returns:
//
//   - one value
//
// # Description:
//
// The same as 'map[[bool]]T {true: vTrue, false: vFalse}[cond]'
func TerOp[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	} else {
		return vFalse
	}
}

// # AbsInt implements standart 'abs' function for int argument.
func AbsInt(num int) int {
	return TerOp(num < 0, -num, num)
}
