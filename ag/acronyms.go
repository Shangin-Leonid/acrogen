package ag /* Acronyms Generation */

import (
	"slices"
	"sort"
)

// # Acronym describes one acronym: the word, summary estimation (sum of letter estimations) and decodings of each letter.
//
// # TODOs:
//
//   - Maybe change the order of fields for better memory allocating.
//   - Maybe use []rune instead of string
type Acronym struct {
	Word            string
	SumEstimation   int
	LetterDecodings []string
}
type Acronyms = []Acronym

// # ContainsAcronym searches for acronym 'word' in Acronyms collection.
//
// # Params:
//
//   - acronym to find
//   - acronyms collection to find in
//
// # Returns:
//
//   - index (0 if not found)
//   - flag if found
func ContainsAcronym(word string, acrs Acronyms) (int, bool) {
	ind := slices.IndexFunc(acrs, func(acr Acronym) bool {
		return word == acr.Word
	})

	return ind, (0 <= ind) && (ind < len(acrs))
}

// # ContainsAcronymBS searches for acronym 'word' in Acronyms collection.
//
// # Params:
//
//   - acronym to find
//   - acronyms collection to find in
//
// # Returns:
//
//   - index (0 if not found)
//   - flag if found
//
// # Description:
//
// Uses binary search. Needs alphabet ordered acronyms collection.
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

// # TakeAcronym searches for acronym 'word' in Acronyms collection.
//
// # Params:
//
//   - acronym to find
//   - acronyms collection to find in
//
// # Returns:
//
//   - found acronym
//   - flag if found
func TakeAcronym(word string, acrs Acronyms) (Acronym, bool) {
	ind, ok := ContainsAcronym(word, acrs)

	if ok {
		return acrs[ind], true
	} else {
		return Acronym{}, false
	}
}

// # ContainsAcronymBS searches for acronym 'word' in Acronyms collection.
//
// # Params:
//
//   - acronym to find
//   - acronyms collection to find in
//
// # Returns:
//
//   - found acronym
//   - flag if found
//
// # Description:
//
// Uses binary search. Needs alphabet ordered acronyms collection.
func TakeAcronymBS(word string, acrs Acronyms) (Acronym, bool) {
	ind, ok := ContainsAcronymBS(word, acrs)

	if ok {
		return acrs[ind], true
	} else {
		return Acronym{}, false
	}
}

// # SortAcronymsBySumEstimation is a wrapper for sorting (in place) Acronyms collection by summary estimations of its elements.
func SortAcronymsBySumEstimation(acrs Acronyms) {
	decreasingSumEstimationComparator := func(i, j int) bool {
		return acrs[i].SumEstimation > acrs[j].SumEstimation
	}
	sort.Slice(acrs, decreasingSumEstimationComparator)
}

// # SortAcronymsByAlphabet is a wrapper for alphabetically sorting (in place) of Acronyms collection.
func SortAcronymsByAlphabet(acrs Acronyms) {
	increasingAlphabetComparator := func(i, j int) bool {
		return acrs[i].Word < acrs[j].Word
	}
	sort.Slice(acrs, increasingAlphabetComparator)
}
