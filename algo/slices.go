package algo /* Algorithms */

// # GetCopy returns a copy of slice.
func GetCopy[T any](origin []T) []T {
	copied := make([]T, len(origin))
	copy(copied, origin)
	return copied
}

// # ReverseSlice reverses a slice in place.
func ReverseSlice[T any](slice []T) {
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
}

// # IndRange
//
// # TODO:
//   - docs
type IndRange struct {
	Beg int
	End int
}

// # SplitSlice
//
// # TODO:
//   - docs
func SplitSlice[T any](slice []T, nParts int) []IndRange {

	if nParts <= 0 {
		return []IndRange{}
	}

	if nParts > len(slice) {
		nParts = len(slice)
	}

	minLen := len(slice) / nParts
	maxLen := minLen + 1
	nWithMaxLen := len(slice) % nParts
	nWithMinLen := nParts - nWithMaxLen

	ir := make([]IndRange, 0, nParts)
	beg := 0
	for range nWithMaxLen {
		ir = append(ir, IndRange{beg, beg + maxLen})
		beg += maxLen
	}
	for range nWithMinLen {
		ir = append(ir, IndRange{beg, beg + minLen})
		beg += minLen
	}

	return ir
}
