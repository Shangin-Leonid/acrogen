package ag /* Acronyms Generation */

import (
	"slices"

	"github.com/Shangin-Leonid/acrogen/algo"
	"github.com/Shangin-Leonid/acrogen/cont"
)

// # AcrGeneratorMode describes a parameter for acronym generating function.
type AcrGeneratorMode int

// Modes of acronym generation.
const (
	Ordered    AcrGeneratorMode = 1
	NonOrdered AcrGeneratorMode = 2
)

// GenerateAcronyms generates acronyms in passed mode.
//
// # Params:
//
//   - source data
//   - dictionary of existing words
//   - mode of generation (see 'AcrGeneratorMode')
//
// # Returns:
//
//   - generated acronyms collection
//
// # Description:
//
// Uses source and dictionary for checking all possible (non-)ordered (depends on 'agm' param) letter combinations and take all that are in the dictionary.
func GenerateAcronyms(src Src, dict Dict, agm AcrGeneratorMode) Acronyms {
	if len(src) == 0 || len(src) == 1 {
		return Acronyms{}
	}

	switch agm {
	case Ordered:
		return generateAcronymsWithOrder(src, dict)
	case NonOrdered:
		return generateAcronymsWithoutOrder(src, dict)
	}

	return Acronyms{}
}

// generateAcronymsWithOrder generates acronyms with strong order of letters.
//
// # Params:
//
//   - source data
//   - dictionary of existing words
//   - mode of generation (see 'AcrGeneratorMode')
//
// # Returns:
//
//   - generated acronyms collection
func generateAcronymsWithOrder(src Src, dict Dict) Acronyms {

	letterCombs := algo.CalcOrderedCartesianProduct(src)

	isRealWord := func(s string) bool {
		_, exist := dict[s]
		return exist
	}

	var acrs Acronyms
	for i := range letterCombs {
		acrCandidate := convertToAcronym(letterCombs[i])
		if isRealWord(acrCandidate.Word) {
			acrs = append(acrs, acrCandidate)
		}
	}

	return acrs
}

// generateAcronymsWithOrder generates acronyms without any order of letters.
//
// # Params:
//
//   - source data
//   - dictionary of existing words
//   - mode of generation (see 'AcrGeneratorMode')
//
// # Returns:
//
//   - generated acronyms collection
func generateAcronymsWithoutOrder(src Src, dict Dict) Acronyms {

	var acrs Acronyms

	perm := cont.NewIdPermutation(len(src))
	for range cont.PermutationsGroupOrder(len(src)) {
		permSrc, _ := cont.GetPermutatedSlice(src, perm)
		newAcrs := generateAcronymsWithOrder(permSrc, dict)
		acrs = slices.Concat(acrs, newAcrs)
		perm.Shift(1)
	}

	return acrs
}
