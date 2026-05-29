package algo /* Algorithms */

// # CalcOrderedCartesianProduct calculates ordered Cartesian product of sets.
//
// # Params:
//
//   - slice that represents a set
//
// # Returns:
//
//   - slice that represents a cartesian product of input set:
//     (nil if input is empty; empty if input contains at least one empty set)
//
// # Description:
//
// Example:
// {1, 2} x {3, 4} x {5, 6} -> {{135}, {136}, {145}, {146}, {235}, {236}, {245}, {246}} (result without commas)
//
// # TODOs:
//
//   - Optimize by using remainder arithmetic.
func CalcOrderedCartesianProduct[T any](inp [][]T) [][]T {
	if len(inp) == 0 {
		return nil
	}

	amountOfOutputSlices := 1
	for i := range inp {
		amountOfOutputSlices *= len(inp[i])
	}

	if amountOfOutputSlices == 0 { // means one of sets is empty, so the product is empty too
		return [][]T{}
	}

	// Prealloc enough memory
	outp := make([][]T, amountOfOutputSlices)
	for i := range outp {
		outp[i] = make([]T, len(inp))
	}

	// Calculate the product

	indicesToTake := make([]int, len(inp))
	updateIndicesToTake := func() {
		i := len(indicesToTake) - 1
		for i > 0 && indicesToTake[i] == len(inp[i])-1 {
			indicesToTake[i] = 0
			i--
		}
		indicesToTake[i]++
	}

	for i := range outp {
		for j, t := range indicesToTake {
			outp[i][j] = inp[j][t]
		}
		updateIndicesToTake()
	}

	return outp
}
