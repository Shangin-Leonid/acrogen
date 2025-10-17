package main

import (
	"fmt"
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

	const LoadDumpChoiceMes = "Would you like to load generated acronyms from dump file?"
	const UserChoiceInputFormatErrMes = "Incorrect choice (incorrect input format)."
	yesOrNo, err := giveUserAChoiceYesOrNo(LoadDumpChoiceMes, UserChoiceInputFormatErrMes)
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
		src, err := importSrcFromFile(srcFilename)
		if err != nil {
			formatAndPrintError(err)
			return
		}

		const ExpectedWordsAmount = 1532570 // 1'532'568 = amount of russian words in my collection
		dict, err := importDictionaryFromFile(dictFilename, ExpectedWordsAmount)
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

	const AcrConsolePrintChoiceMes = "Would you like to print all generated acronyms in console?"
	yesOrNo, err = giveUserAChoiceYesOrNo(AcrConsolePrintChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return
	}
	if yesOrNo == Yes {
		printAcronyms(acrs)
	} else if yesOrNo == No {
		return
	}

	const DecodeChoiceMes = "Would you like to decode any generated acronym?"
	yesOrNo, err = giveUserAChoiceYesOrNo(DecodeChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return
	}

	if yesOrNo == Yes {
		containsAcronymWrap := func(userInp string) (bool, error) {
			return containsAcronym(userInp, acrs), nil
		}
		takeAndPrintAcronym := func(userInp string) error {
			acr, _ := takeAcronym(userInp, acrs) // No error can be, we've just checked that the acronym is in the collection
			printAcronymInDetail(acr)
			return nil
		}

		err, _ := processUserInputUntilExitCommand(
			"",
			"\nPlease, enter an acronym:",
			"No such acronym was found.",
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
