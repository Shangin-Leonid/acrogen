package main

import (
	"fmt"
	"os"

	"acrogen/cio"
	"acrogen/fio"
)

func main() {
	// Read and check command line arguments.
	argsWithoutProgName := os.Args[1:]
	if len(argsWithoutProgName) != 4 {
		fmt.Println("Error: incorrect number of program input arguments!")
		fmt.Println("Restart the program with passing 4 names of \".txt\" files: for input, with real existing words (dictionary), for acronyms dump and for output.")
		return
	}
	srcFilename, dictFilename := argsWithoutProgName[0], argsWithoutProgName[1]
	dumpFilename, outputFilename := argsWithoutProgName[2], argsWithoutProgName[3]
	if dumpFilename == outputFilename {
		fmt.Println("Error: dump and output filenames must be different.")
	}

	// Give a choice of program mode: load existing (generated earlier) acronyms from dump file or build them from source.
	var acrs Acronyms
	const LoadDumpChoiceMes = "Load generated acronyms from dump file? Else they will be generated from source."
	const UserChoiceInputFormatErrMes = "Unexpected choice (incorrect input format)."
	yesOrNo, err := cio.GiveUserYesOrNoChoice(LoadDumpChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return
	}
	if yesOrNo == cio.Yes {
		// Load existing (generated earlier) acronyms from dump file.
		acrs, err = loadAcronymsFromFile(dumpFilename)
		if err != nil {
			formatAndPrintError(err)
			return
		}
		fmt.Printf("\n%d acronyms have been successfully loaded from '%s'.\n", len(acrs), dumpFilename)
	} else if yesOrNo == cio.No {
		// Generate acronyms from source.

		// Load source data from file.
		src, err := loadSrcFromFile(srcFilename)
		if err != nil {
			formatAndPrintError(err)
			return
		}

		// Load dictionary from file.
		const ExpectedWordsAmount = 1532570 // 1'532'568 = amount of russian words in my collection
		dict, err := loadDictionaryFromFile(dictFilename, ExpectedWordsAmount)
		if err != nil {
			formatAndPrintError(err)
			return
		}

		// Give a choice of program mode: generate acronyms with or without strict order.
		const AcrGenerationModeChoiceMes = "Does the order of items in acronym matter?"
		yesOrNo, err = cio.GiveUserYesOrNoChoice(AcrGenerationModeChoiceMes, UserChoiceInputFormatErrMes)
		if err != nil {
			formatAndPrintError(err)
			return
		}
		var mode AcrGeneratorMode
		if yesOrNo == cio.Yes {
			mode = Ordered
		} else if yesOrNo == cio.No {
			mode = NonOrdered
		}

		// Generate and sort acronyms.
		acrs = generateAcronyms(src, dict, mode)
		if len(acrs) == 0 {
			return
		}
		SortAcronymsBySumEstimation(acrs)
		fmt.Printf("\n%d acronyms were successfully generated and sorted by their estimation.\n", len(acrs))

		// Export generated acronyms to the dump file.
		const dumpFilenameSuffix = "_dump"
		dumpOutputFilename := fio.GetWithoutExt(outputFilename) + dumpFilenameSuffix + ".txt"
		err = exportAcronymsToFile(acrs, dumpOutputFilename, FullFormat)
		if err != nil {
			formatAndPrintError(err)
			return
		}
	}

	// Process generated acronyms.

	// Give a choice and maybe print some acronyms to console.
	const AcrConsolePrintChoiceMes = "Print most suitable acronyms in console?"
	yesOrNo, err = cio.GiveUserYesOrNoChoice(AcrConsolePrintChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return
	}
	if yesOrNo == cio.Yes {
		const AmountOfAcronymsChoiceMes = "Choose the number of acronyms for printing (0 for all)."
		const IncorrectNumberMes = "Unexpected choice (a number was expected)."
		amount, err := cio.GiveUserNumberChoice(AmountOfAcronymsChoiceMes, IncorrectNumberMes)
		if err != nil {
			formatAndPrintError(err)
			return
		}

		err = printMostSuitableAcronyms(acrs, amount)
		if err != nil {
			formatAndPrintError(err)
			return
		}
	} else if yesOrNo == cio.No {
		/* Do nothing */
	}

	// Maybe decode some acronyms.
	const DecodeChoiceMes = "Decode some generated acronyms?"
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

	fmt.Printf("\n\"Acrogen\" (\"%s\") finished with success.\n", os.Args[0])
	return
}

// #
// Formats and prints error in console in 'stderr'.
// #
func formatAndPrintError(err error) {
	fmt.Fprintln(os.Stderr, fmt.Errorf("Error: %w.", err))
}
