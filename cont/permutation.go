package cont /* Containers */

import (
	"errors"

	"acrogen/algo"
	"acrogen/utils"
)

// TODO refactor with OOP
// TODO operator < > ==

// Permutation is a bijection [0, 1, ... N] <--> [0, 1, ... N] with lexicographical order.
type Permutation struct {
	elems []int
}

// #
// Creates new identity permutation (0, 1, ... length-1).
// #
func NewIdPermutation(length int) Permutation {
	idPerm := Permutation{make([]int, length)}
	for i := range idPerm.elems {
		idPerm.elems[i] = i
	}
	return idPerm
}

// #
// Returns a length (size) of the permutation.
// #
func (p Permutation) Len() int {
	return len(p.elems)
}

// #
// Returns the copy of underlying slice.
// #
func (p Permutation) AsSlice() []int {
	return algo.GetCopy(p.elems)
}

// #
// Returns the i-th element.
// #
func (p Permutation) get(i int) int {
	return p.elems[i]
}

// #
// Resizes the permutation and set it to Id.
// #
func (p *Permutation) Resize(newLength int) {
	*p = NewIdPermutation(newLength)
}

// #
// Returns the next (lexicographically) permutation.
// #
func (p Permutation) Next() Permutation {
	nextP := GetCopy(p)
	nextP.shiftToNext()
	return nextP
}

// #
// Returns the previous (lexicographically) permutation.
// #
func (p Permutation) Prev() Permutation {
	prevP := GetCopy(p)
	prevP.shiftToPrev()
	return prevP
}

// #
// Shifts in place with 's' steps toward or backward (lexicographically).
// #
func (p *Permutation) Shift(s int) {
	var shiftFunc func(*Permutation)
	if s < 0 {
		shiftFunc = (*Permutation).shiftToPrev
		s = -s
	} else {
		shiftFunc = (*Permutation).shiftToNext
	}

	for range s {
		shiftFunc(p)
	}
}

// #
// Shifts the permutation to next (lexicographically).
// #
func (p *Permutation) shiftToNext() {
	elems := p.elems

	if len(elems) <= 1 {
		return
	}

	i := len(elems) - 2
	for i >= 0 && elems[i] >= elems[i+1] {
		i--
	}

	if i >= 0 {
		j := len(elems) - 1
		for elems[j] <= elems[i] {
			j--
		}
		elems[i], elems[j] = elems[j], elems[i]
	}

	algo.ReverseSlice(elems[i+1:])
}

// #
// Shifts the permutation to previous (lexicographically).
// #
func (p *Permutation) shiftToPrev() {
	// TODO reverse shifting instead
	p.Shift(PermutationsGroupOrder(p.Len()) - 1)
}

// #
// Returns if slice can represent a valid permutation.
// #
func IsPermutation(slice []int) bool {
	type void = utils.Void
	elems := make(map[int]void, len(slice))
	for i := range slice {
		// Check validity of elements range
		if (slice[i] < 0) || (slice[i] > len(slice)-1) {
			return false
		}
		elems[slice[i]] = void{}
	}

	// Check elements uniqueness
	return len(elems) == len(slice)
}

// #
// Returns a copy of permutation.
// #
func GetCopy(p Permutation) Permutation {
	return Permutation{algo.GetCopy(p.elems)}
}

// #
// Returns copy of 'slice' permutated by 'perm'.
// #
func GetPermutatedSlice[T any](slice []T, perm Permutation) ([]T, error) {
	if len(slice) < perm.Len() {
		return nil, errors.New("incorrect slice and permutation sizes in 'GetPermutatedSlice()'")
	}

	permutated := algo.GetCopy(slice)

	for ind := 0; ind < perm.Len(); ind++ {
		permutated[ind] = slice[perm.get(ind)]
	}

	return permutated, nil
}

// #
// Returns the order of S_n group that contains permutation of 'permLength'.
// #
func PermutationsGroupOrder(permLength int) int {
	return int(algo.CalcFactorial(uint(permLength)))
}
