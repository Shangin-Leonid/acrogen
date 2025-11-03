package algo /* Algorithms */

// #
// Returns a copy of slice.
// #
func GetCopy[T any](origin []T) []T {
	copied := make([]T, len(origin))
	copy(copied, origin)
	return copied
}

// #
// Reverses a slice in place.
// #
func ReverseSlice[T any](slice []T) {
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
}
