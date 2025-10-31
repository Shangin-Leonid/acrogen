package ag /* Acronyms Generation */

import (
	"slices"
	"sort"
)

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
