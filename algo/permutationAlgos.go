package algo /* Algorithms */

import "errors" // TODO remake with OOP

// WARNING! All permutations are zero based (strats from zero).

// #
// TODO Docs
// #
func GetPermutatedSlice[T any](slice []T, perm []int) ([]T, error) {
	if len(slice) != len(perm) {
		return nil, errors.New("incorrect slice and permutation sizes in 'GetPermutatedSlice()'")
	}

	permutated := make([]T, len(slice))
	_ = copy(permutated, slice)

	for ind, indImage := range perm {
		permutated[ind] = slice[indImage]
	}

	return permutated, nil
}

// #
// TODO Docs
// #
func ChangeToNextPermutation(perm []int) {
	if len(perm) <= 1 {
		return
	}

	i := len(perm) - 2
	for i >= 0 && perm[i] >= perm[i+1] {
		i--
	}

	if i >= 0 {
		j := len(perm) - 1
		for perm[j] <= perm[i] {
			j--
		}
		perm[i], perm[j] = perm[j], perm[i]
	}

	reverseSlice(perm[i+1:])
}

// #
// TODO Docs
// #
func GetIdPermutation(length int) []int {
	idPerm := make([]int, length)
	for i := range idPerm {
		idPerm[i] = i
	}
	return idPerm
}

func reverseSlice[T any](slice []T) {
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
}
