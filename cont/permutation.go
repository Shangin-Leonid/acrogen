package cont /* Containers */

import (
	"errors"

	"github.com/Shangin-Leonid/acrogen/algo"
	"github.com/Shangin-Leonid/acrogen/utils"
)

// # Permutation is a bijection [0, 1, ... N] <--> [0, 1, ... N] with lexicographical order.
type Permutation struct {
	elems []int
}

// # NewIdPermutation creates new identity permutation (0, 1, ... length-1).
func NewIdPermutation(length int) *Permutation {
	idPerm := Permutation{make([]int, length)}
	for i := range idPerm.elems {
		idPerm.elems[i] = i
	}

	return &idPerm
}

// # Len returns a length (size) of the permutation.
func (p *Permutation) Len() int {
	return len(p.elems)
}

// # AsSlice returns the copy of underlying int slice.
func (p *Permutation) AsSlice() []int {
	return algo.GetCopy(p.elems)
}

// # Clone returns the copy of permutation.
func (p *Permutation) Clone() *Permutation {
	return &Permutation{elems: algo.GetCopy(p.elems)}
}

// # get returns the i-th element.
func (p *Permutation) Get(i int) int {
	return p.elems[i]
}

// # Resize resizes the permutation and set it to Id.
func (p *Permutation) Resize(newLength int) {
	*p = *NewIdPermutation(newLength)
}

// # Next returns the next (lexicographically) permutation.
func (p *Permutation) Next() *Permutation {
	nextP := p.Clone()
	nextP.shiftToNext()
	return nextP
}

// # Prev returns the previous (lexicographically) permutation.
func (p *Permutation) Prev() *Permutation {
	prevP := p.Clone()
	prevP.shiftToPrev()
	return prevP
}

// # Shift shifts in place with 's' steps toward or backward (lexicographically).
func (p *Permutation) Shift(s int) {

	shiftFunc := utils.TerOp(s < 0, (*Permutation).shiftToPrev, (*Permutation).shiftToNext)
	s = utils.AbsInt(s)

	for range s {
		shiftFunc(p)
	}
}

// # shiftToNext shifts the permutation to next (lexicographically).
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

// # shiftToPrev shifts the permutation to previous (lexicographically).
//
// TODOs:
//
//   - Optimize: reverse shifting instead.
func (p *Permutation) shiftToPrev() {
	p.Shift(PermutationsGroupOrder(p.Len()) - 1)
}

// # IsPermutation returns if slice can represent a valid permutation.
func IsPermutation(slice []int) bool {

	if len(slice) == 0 {
		return false
	}

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

// # GetPermutatedSlice returns copy of 'slice' permutated by 'perm' and error.
func GetPermutatedSlice[T any](slice []T, perm *Permutation) ([]T, error) {
	if len(slice) < perm.Len() {
		return nil, errors.New("incorrect slice and permutation sizes in 'GetPermutatedSlice()'")
	}

	permutated := algo.GetCopy(slice)

	for ind := 0; ind < perm.Len(); ind++ {
		permutated[ind] = slice[perm.Get(ind)]
	}

	return permutated, nil
}

// # PermutationsGroupOrder returns the order of S_n group that contains permutation of 'permLength'.
func PermutationsGroupOrder(permLength int) int {
	return int(algo.CalcFactorial(uint(permLength)))
}
