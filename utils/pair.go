package utils /* Utils */

// # Pair implements common data structure called 'Pair'.
// Can storage 2 values of any different or equal types.
type Pair[T1, T2 any] struct {
	First  T1
	Second T2
}
