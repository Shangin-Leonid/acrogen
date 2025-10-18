package main

import (
	"slices"
	"sort"

	"acrgen/algo"
)

// #
// Describes one source file entry (line), that represents a variant of acronym letter, its estimation and decoding (description).
// #
type LetterOpt struct {
	letter     rune
	estimation int
	decoding   string
}
type LetterOpts = []LetterOpt
type Src = []LetterOpts

type Acronym struct {
	word            string // TODO maybe use []rune instead of string
	sumEstimation   int
	letterDecodings []string
}
type Acronyms = []Acronym

// #
// Searches for acronym 'word' in Acronyms collection.
// Returns true if have found, false else.
// #
func containsAcronym(word string, acrs Acronyms) bool {
	ind := slices.IndexFunc(acrs, func(acr Acronym) bool {
		return word == acr.word
	})

	return (0 <= ind) && (ind < len(acrs))
}

// #
// Searches for acronym 'word' in Acronyms collection.
// Returns (acronym, true) if have found, (nil, false) else.
// #
func takeAcronym(word string, acrs Acronyms) (Acronym, bool) {
	ind := slices.IndexFunc(acrs, func(acr Acronym) bool {
		return word == acr.word
	})

	if 0 <= ind && ind < len(acrs) {
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
func generateAcronyms(src Src, dict Dict, agm AcrGeneratorMode) Acronyms {
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
			word = append(word, lo[i].letter)
			sumEstimation += lo[i].estimation
			letterDecodings = append(letterDecodings, lo[i].decoding)
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
		if isRealWord(acrCandidate.word) {
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
	for _ = range nPermutations {
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
	isMoreSumEstimationComparator := func(i, j int) bool {
		return acrs[i].sumEstimation > acrs[j].sumEstimation
	}
	sort.Slice(acrs, isMoreSumEstimationComparator)
}
