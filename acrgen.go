package main

import (
	"errors"
	"fmt"
	"math"
	"os"

	"acrgen/cio"
	"acrgen/fio"
)

func main() {
	argsWithoutProgName := os.Args[1:]
	if len(argsWithoutProgName) != 4 {
		fmt.Println("Incorrect number of program input arguments!")
		fmt.Println("Restart the program with passing 4 names of \".txt\" files: for input, with real existing words (dictionary), for acronyms dump and for output.")
		return
	}
	srcFilename, dictFilename := argsWithoutProgName[0], argsWithoutProgName[1]
	dumpFilename, outputFilename := argsWithoutProgName[2], argsWithoutProgName[3]

	var acrs Acronyms

	const LoadDumpChoiceMes = "Would you like to load generated acronyms from dump file? Else they will be generated from source."
	const UserChoiceInputFormatErrMes = "Unexpected choice (incorrect input format)."
	yesOrNo, err := cio.GiveUserYesOrNoChoice(LoadDumpChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return
	}
	if yesOrNo == cio.Yes {
		acrs, err = LoadAcronymsFromFile(dumpFilename)
		if err != nil {
			formatAndPrintError(err)
			return
		}

		fmt.Printf("\n%d acronyms were successfully loaded from '%s'.\n", len(acrs), dumpFilename)
	} else if yesOrNo == cio.No {
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

		acrs = generateAcronyms(src, dict, NonOrdered)
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
	yesOrNo, err = cio.GiveUserYesOrNoChoice(AcrConsolePrintChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return
	}
	if yesOrNo == cio.Yes {
		const AmountOfAcronymsChoiceMes = "Choose number of acronyms for console printing (0 for all)."
		const IncorrectNumberMes = "Unexpected choice (a number was expected)."
		amount, err := cio.GiveUserNumberChoice(AmountOfAcronymsChoiceMes, IncorrectNumberMes)
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
	} else if yesOrNo == cio.No {
		return
	}

	const DecodeChoiceMes = "Would you like to decode any generated acronym?"
	yesOrNo, err = cio.GiveUserYesOrNoChoice(DecodeChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return
	}

	if yesOrNo == cio.Yes {
		containsAcronymWrap := func(userInp string) (bool, error) {
			return containsAcronym(userInp, acrs), nil
		}
		takeAndPrintAcronym := func(userInp string) error {
			acr, _ := takeAcronym(userInp, acrs) // No need to check 'ok': we've just checked that the acronym is in the collection
			printAcronymInDetail(acr)
			fmt.Printf("\n")
			return nil
		}

		err, _ := cio.ProcessUserInputUntilExitCommand(
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
