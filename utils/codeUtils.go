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
