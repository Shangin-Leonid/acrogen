package ag /* Acronyms Generation */

import (
	"slices"
	"sort"

	"acrogen/algo"
)

// #
// Describes one source file entry (line), that represents a variant of acronym letter, its estimation and decoding (description).
// #
type LetterOpt struct {
	Letter     rune
	Estimation int
	Decoding   string
}
type LetterOpts = []LetterOpt
type Src = []LetterOpts

// #
// Describes one acronym: the word, summary estimation (sum of letter estimations) and decodings of each letter.
// #
// TODO maybe change the order of fields for better memory allocating
type Acronym struct {
	Word            string // TODO maybe use []rune instead of string
	SumEstimation   int
	LetterDecodings []string
}
type Acronyms = []Acronym

// #
// Searches for acronym 'word' in Acronyms collection.
// Returns index and true if have found, some int and false else.
// #
func ContainsAcronym(word string, acrs Acronyms) (int, bool) {
	ind := slices.IndexFunc(acrs, func(acr Acronym) bool {
		return word == acr.Word
	})

	return ind, (0 <= ind) && (ind < len(acrs))
}

// #
// Searches for acronym 'word' in Acronyms collection by binary search (collection must be in alphabet order).
// Returns index and true if have found, index of place for inserting and false else.
// #
func ContainsAcronymBS(word string, acrs Acronyms) (int, bool) {
	return slices.BinarySearchFunc(acrs, word, func(acr Acronym, word string) int {
		switch {
		case acr.Word < word:
			return -1
		case acr.Word > word:
			return 1
		default:
			return 0
		}
	})
}

// #
// Searches for acronym 'word' in Acronyms collection.
// Returns (acronym, true) if have found, (empty acronym, false) else.
// #
func TakeAcronym(word string, acrs Acronyms) (Acronym, bool) {
	ind, ok := ContainsAcronym(word, acrs)

	if ok {
		return acrs[ind], true
	} else {
		return Acronym{}, false
	}
}

// #
// Searches for acronym 'word' in Acronyms collection by binary search (collection must be in alphabet order).
// Returns (acronym, true) if have found, (empty acronym, false) else.
// #
func TakeAcronymBS(word string, acrs Acronyms) (Acronym, bool) {
	ind, ok := ContainsAcronymBS(word, acrs)

	if ok {
		return acrs[ind], true
	} else {
		return Acronym{}, false
	}
}

// #
// Describes a set of existing and valid words, candidates for acronyms.
// #
type Dict map[string]struct{}

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

// #
// A wrapper for sorting Acronyms collection by summary estimations of its elements.
// Returns nothing, just sorts in place.
// #
func SortAcronymsBySumEstimation(acrs Acronyms) {
	decreasingSumEstimationComparator := func(i, j int) bool {
		return acrs[i].SumEstimation > acrs[j].SumEstimation
	}
	sort.Slice(acrs, decreasingSumEstimationComparator)
}

// #
// A wrapper for alphabetically sorting of Acronyms collection.
// Returns nothing, just sorts in place.
// #
func SortAcronymsByAlphabet(acrs Acronyms) {
	increasingAlphabetComparator := func(i, j int) bool {
		return acrs[i].Word < acrs[j].Word
	}
	sort.Slice(acrs, increasingAlphabetComparator)
}
