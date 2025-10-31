package ui /* User Interface */

import (
	"errors"
	"fmt"

	"acrogen/ag"
)

// #
// Prints acronym in console in detailed format (decodes each letter).
// #
func printAcronymInDetail(acr ag.Acronym) {
	fmt.Printf("%s%s%d\n", acr.Word, TokenSeparator, acr.SumEstimation)
	for i, letter := range []rune(acr.Word) {
		fmt.Printf("%c -- %s\n", letter, acr.LetterDecodings[i])
	}
}

// #
// Prints acronyms in console in poor format (acronym only, without any decoding info).
// 'amount' == 0 means printing all acronyms.
// #
func printAcronyms(acrs ag.Acronyms, amount int) error {

	switch {
	case amount < 0:
		return errors.New("incorrect (negative) amount of acronyms")
	case amount == 0:
		amount = len(acrs)
	case amount > len(acrs):
		return errors.New("too many acronyms are requested to print")
	}

	SuccessColor.Printf("\nList of acronyms:\n")
	for i := 0; i < amount; i++ {
		fmt.Printf("%s%s%d\n", acrs[i].Word, TokenSeparator, acrs[i].SumEstimation)
	}

	return nil
}

// #
// Prints most suitable (by SumEstimation) acronyms in console in poor format (acronym only, without any decoding info).
// #
func printMostSuitableAcronyms(acrs ag.Acronyms, amount int) error {

	// TODO optimize by not sorting all elements, but by taking the amount of best ones
	sortedAcrs := make(ag.Acronyms, len(acrs))
	copy(sortedAcrs, acrs)
	ag.SortAcronymsBySumEstimation(sortedAcrs)

	return printAcronyms(sortedAcrs, amount)
}
