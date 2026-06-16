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

// # IndRange storages a range of indices.
type IndRange struct {
	Beg int
	End int
}

// # SplitSlice splits (decomposes) a slice into as equal as possible chunks.
//
// # Params:
//
//   - slice
//   - expected amount of chunks
//
// # Returns:
//
//   - slice of IndRange representing the chunks
//
// # Description:
//
// It is guaranteed that len of every chunk is no less than 1.
// It is also guaranteed that diff between max and min chunk is no more than 1 (maybe 0).
// If 'len(slice)' == 0 or 'nParts' <= 0 then empty slice is returned.
// If 'nParts' > 'len(slice)' then 'len(slice)' is used instead of 'nParts'.
// If there is no opportunity to get equal chunks, then max len chunks are stored before min.
//
// Example:
// slice of len 5 + 'nParts' == 3 -> []IndRange{{0, 2} {2, 4} {4, 5}}
func SplitSlice[T any](slice []T, nParts int) []IndRange {

	if nParts <= 0 || len(slice) == 0 {
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
