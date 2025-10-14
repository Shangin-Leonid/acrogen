package main

import (
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
	word            string
	sumEstimation   int
	letterDecodings []string
}
type Acronyms = []Acronym

// #
// Describes a set of existing and valid words, candidates for acronyms.
// #
type Dict map[string]struct{}

// #
// Generates acronyms from 'src': check all possible letter sequences and take some that are in dictionary.
// Returns Acronyms collection and error.
// #
func generateAcronyms(src Src, dict Dict) (acrs Acronyms, err error) {
	letterCombs, err := algo.CalcOrderedCartesianProduct(src)
	if err != nil {
		return nil, err
	}

	convertToAcronym := func(lo LetterOpts) Acronym {
		word := []rune{}
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

	for i := range letterCombs {
		acrCandidate := convertToAcronym(letterCombs[i])
		if isRealWord(acrCandidate.word) {
			acrs = append(acrs, acrCandidate)
		}
	}

	return acrs, nil
}

// #
// A wrapper for sorting Acronyms collection by summary estimations of its elements.
// Returns nothing, just sorts in place.
// #
func SortAcronymsBySumEstimation(acrs Acronyms) {
	isMoreSumEstimationFunc := func(i, j int) bool {
		return acrs[i].sumEstimation > acrs[j].sumEstimation
	}
	sort.Slice(acrs, isMoreSumEstimationFunc)
}
