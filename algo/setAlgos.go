package algo

import (
	"errors"
)

// #
// Calculates and returns ordered Cartesion product of sets.
// Input and output sets are actually represented by slices (as types).
// Here is an example:
// {1, 2} x {3, 4} x {5, 6} -> {{135}, {136}, {145}, {146}, {235}, {236}, {245}, {246}} (result without commas)
// #

func CalcOrderedCartesianProduct[T any](inp [][]T) ([][]T, error) {
	if len(inp) == 0 {
		return nil, errors.New("no sets (slices) were passed")
	}

	// Prealloc enough memory
	amountOfOutputSlices := 1
	for i := range inp {
		amountOfOutputSlices *= len(inp[i])
	}

	if amountOfOutputSlices == 0 { // means one of sets is empty, so the product is empty too
		return [][]T{}, nil // return empty set
	}

	outp := make([][]T, amountOfOutputSlices)
	for i := range outp {
		outp[i] = make([]T, len(inp))
	}

	// Calculate the product

	indicesToTake := make([]int, len(inp))
	// TODO optimize by using remainder arithmetic
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

	return outp, nil
}
