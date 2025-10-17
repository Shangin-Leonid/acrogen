package main

import (
	"errors"
	"fmt"
	"math"
	"os"

	"acrgen/fio"
)

func main() {
	argsWithoutProgName := os.Args[1:]
	if len(argsWithoutProgName) != 4 {
		fmt.Println("Incorrect number of program input arguments!")
		fmt.Println("Restart the program with passing 4 names of \".txt\" files: for input, with real existing words (dictionary) and for output.")
		return
	}
	srcFilename, dictFilename := argsWithoutProgName[0], argsWithoutProgName[1]
	dumpFilename, outputFilename := argsWithoutProgName[2], argsWithoutProgName[3]

	var acrs Acronyms

	const LoadDumpChoiceMes = "Would you like to load generated acronyms from dump file? Else they will be generated from source."
	const UserChoiceInputFormatErrMes = "Unexpected choice (incorrect input format)."
	yesOrNo, err := giveUserYesOrNoChoice(LoadDumpChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return
	}
	if yesOrNo == Yes {
		acrs, err = LoadAcronymsFromFile(dumpFilename)
		if err != nil {
			formatAndPrintError(err)
			return
		}

		fmt.Printf("\n%d acronyms were successfully loaded from '%s'.\n", len(acrs), dumpFilename)
	} else if yesOrNo == No {
		src, err := loadSrcFromFile(srcFilename)
		if err != nil {
			formatAndPrintError(err)
			return
		}

		const ExpectedWordsAmount = 1532570 // 1'532'568 = amount of russian words in my collection
		dict, err := loadDictionaryFromFile(dictFilename, ExpectedWordsAmount)
		if err != nil {
			formatAndPrintError(err)
			return
		}

		acrs, err = generateAcronyms(src, dict)
		if err != nil {
			formatAndPrintError(err)
			return
		}
		SortAcronymsBySumEstimation(acrs)

		fmt.Printf("\n%d acronyms were successfully generated and sorted by their estimation.\n", len(acrs))

		const dumpFilenameSuffix = "_dump"
		dumpOutputFilename := fio.GetWithoutExt(outputFilename) + dumpFilenameSuffix + ".txt"
		err = exportAcronymsToFile(acrs, dumpOutputFilename, FullFormat)
		if err != nil {
			formatAndPrintError(err)
			return
		}
	}

	const AcrConsolePrintChoiceMes = "Would you like to print acronyms in console?"
	yesOrNo, err = giveUserYesOrNoChoice(AcrConsolePrintChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return
	}
	if yesOrNo == Yes {
		const AmountOfAcronymsChoiceMes = "Choose number of acronyms for console printing (0 for all)."
		const IncorrectNumberMes = "Unexpected choice (a number was expected)."
		amount, err := giveUserNumberChoice(AmountOfAcronymsChoiceMes, IncorrectNumberMes)
		if err != nil {
			formatAndPrintError(err)
			return
		}
		if amount > len(acrs) {
			formatAndPrintError(errors.New("too many acronyms are requested to print"))
			return
		}
		if amount == 0 {
			amount = math.MaxInt
		}
		printAcronyms(acrs, amount)
	} else if yesOrNo == No {
		return
	}

	const DecodeChoiceMes = "Would you like to decode any generated acronym?"
	yesOrNo, err = giveUserYesOrNoChoice(DecodeChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return
	}

	if yesOrNo == Yes {
		containsAcronymWrap := func(userInp string) (bool, error) {
			return containsAcronym(userInp, acrs), nil
		}
		takeAndPrintAcronym := func(userInp string) error {
			acr, _ := takeAcronym(userInp, acrs) // No need to check 'ok': we've just checked that the acronym is in the collection
			printAcronymInDetail(acr)
			fmt.Printf("\n")
			return nil
		}

		err, _ := processUserInputUntilExitCommand(
			"",
			"\nPlease, enter an acronym:",
			"No such acronym was found.\n",
			containsAcronymWrap,
			takeAndPrintAcronym)
		if err != nil {
			formatAndPrintError(err)
			return
		}
	}

	fmt.Println("\n\"Acrgen\" finished with success.")
	return
}

func formatAndPrintError(err error) {
	fmt.Println(fmt.Errorf("Error: %w", err))
}
