package ag /* Acronyms Generation */

import (
	"slices"

	"acrogen/algo"
)

// #
// Describes a parameter for acronym generating function.
// #
type AcrGeneratorMode int

const (
	Ordered    AcrGeneratorMode = 1
	NonOrdered AcrGeneratorMode = 2
)

// #
// Generates acronyms from 'src': check all possible (non-)ordered (depends on 'agm' param) letter combinations and take all that are in dictionary.
// Returns Acronyms collection.
// #
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

// #
// Generates acronyms from 'src': check all possible ordered letter combinations and take all that are in dictionary.
// Returns Acronyms collection.
// #
func generateAcronymsWithOrder(src Src, dict Dict) Acronyms {
	letterCombs, _ := algo.CalcOrderedCartesianProduct(src)

	convertToAcronym := func(lo LetterOpts) Acronym {
		word := make([]rune, 0, len(lo))
		sumEstimation := 0
		letterDecodings := []string{}

		for i := range lo {
			word = append(word, lo[i].Letter)
			sumEstimation += lo[i].Estimation
			letterDecodings = append(letterDecodings, lo[i].Decoding)
		}

		return Acronym{string(word), sumEstimation, letterDecodings}
	}

	isRealWord := func(s string) bool {
		if _, exist := dict[s]; exist {
			return true
		} else {
			return false
		}
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

// #
// Generates acronyms from 'src': check all possible non-ordered letter combinations and take all that are in dictionary.
// Returns Acronyms collection.
// #
func generateAcronymsWithoutOrder(src Src, dict Dict) Acronyms {
	var acrs Acronyms

	perm := algo.GetIdPermutation(len(src))
	nPermutations := int(algo.CalcFactorial(uint(len(src))))
	for range nPermutations {
		permSrc, _ := algo.GetPermutatedSlice(src, perm)
		newAcrs := generateAcronymsWithOrder(permSrc, dict)
		acrs = slices.Concat(acrs, newAcrs)
		algo.ChangeToNextPermutation(perm)
	}

	return acrs
}
