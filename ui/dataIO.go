package ui /* User Interface */

import (
	"fmt"

	"github.com/Shangin-Leonid/acrogen/ag"
	"github.com/Shangin-Leonid/acrogen/utils"
)

// # printAcronymInDetail prints an acronym in console in detailed format (decodes each letter).
//
// # Params:
//
//   - acronym to print
//
// # Returns: -
func printAcronymInDetail(acr ag.Acronym) {
	fmt.Printf("%s%s%d\n", acr.Word, TokenSeparator, acr.SumEstimation)
	for i, letter := range []rune(acr.Word) {
		fmt.Printf("%c -- %s\n", letter, acr.LetterDecodings[i])
	}
}

// # printAcronyms prints first acronyms from collection in console.
//
// # Params:
//
//   - collection of acronyms to print
//   - amount of acronyms to print
//
// # Returns:
//
//   - error
//
// # Description:
//
// The format is poor (acronym only, without any decoding info).
//
// 'amount' == 0 means printing all acronyms.
//
// # TODOs:
//
//   - Use 'all' instead of '0' command to print all acronyms.
func printAcronyms(acrs ag.Acronyms, amount int) error {

	switch {
	case amount < 0:
		return utils.NewSTError("incorrect (negative) amount of acronyms")
	case amount == 0:
		amount = len(acrs)
	case amount > len(acrs):
		return utils.NewSTError("too many acronyms are requested to print")
	}

	SuccessColor.Printf("\nList of acronyms:\n")
	for i := 0; i < amount; i++ {
		fmt.Printf("%s%s%d\n", acrs[i].Word, TokenSeparator, acrs[i].SumEstimation)
	}

	return nil
}

// # printMostSuitableAcronyms prints most suitable (by SumEstimation) acronyms in console.
//
// # Params:
//
//   - collection of acronyms to select from
//   - amount of acronyms to print
//
// # Returns:
//
//   - error
//
// # Description:
//
// The format is poor (acronym only, without any decoding info).
//
// # TODOs:
//
//   - Use another data structure (skip list maybe) to avoid sorting and getting suitable acronyms faster.
func printMostSuitableAcronyms(acrs ag.Acronyms, amount int) error {
	sortedAcrs := make(ag.Acronyms, len(acrs))
	copy(sortedAcrs, acrs)
	ag.SortAcronymsBySumEstimation(sortedAcrs)

	err := printAcronyms(sortedAcrs, amount)

	return err
}
