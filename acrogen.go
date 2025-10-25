package main

import (
	"errors"
	"fmt"
	"os"

	"acrogen/cio"
)

func main() {

	runApp()

	return
}

// #
// Runs the 'acrogen' application.
// The inner point of the program.
// #
func runApp() {
	var err error
	var acrs Acronyms

	printMenuInfo()
	var userInp string
	for {
		fmt.Printf("> Enter a command:\n")

		_, err = fmt.Scanf("%s", &userInp)
		if err != nil {
			formatAndPrintError(err)
			return
		}

		switch userInp {
		case HelpCommand:
			runHelpMode()
		case ExitProgramCommand:
			fmt.Printf("\n\"Acrogen\" (\"%s\") finished with success.\n", os.Args[0])
			return
		case QuitModeCommand:
			fmt.Printf("You are in the menu - nothing to quit from.\n")
		case LoadAcronymsFromFileCommand:
			acrs = runLoadingAcronymsFromFileMode()
		case GenerateAcronymsFromSourceCommand:
			acrs = runGeneratingAcronymsFromSourceMode()
		case PrintListOfAcronymsCommand:
			runListOfAcronymsPrintingMode(acrs)
		case DecodeAcronymCommand:
			runAcronymsDecodingMode(acrs)
		case SaveAcronymsToFileCommand:
			// HERE
			runSavingAcronymsToFileMode(acrs)
		default:
			processIncorrectUserMenuCommand(userInp)
		}

		fmt.Printf("\n")
	}
}

// #
// Formats and prints error in console in 'stderr'.
// #
func formatAndPrintError(err error) {
	fmt.Fprintln(os.Stderr, fmt.Errorf("> Error: %w.", err))
}

// #
// Helps a user to apply the 'acrogen'.
// #
func runHelpMode() {
	// TODO
	printMenuInfo()
}

// #
// TODO docs
// #
func runLoadingAcronymsFromFileMode() Acronyms {
	fmt.Println("> Loading acronyms from file:")

	// Give a choice of input file
	yesOrNo, err := cio.GiveUserYesOrNoChoice(UseDefaultDumpFileChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return nil
	}
	var filename string
	if yesOrNo == cio.Yes {
		filename = DumpDefaultFilename
	} else if yesOrNo == cio.No {
		filename, err = cio.GiveUserChoiceOfFilename("> Enter a name of file:\n")
		if err != nil {
			formatAndPrintError(err)
			return nil
		}
	}

	// Load acronyms from file
	acrs, err := loadAcronymsFromFile(filename)
	if err != nil {
		formatAndPrintError(err)
		return nil
	}
	fmt.Printf("\n> %d acronyms have been successfully loaded from '%s'.\n", len(acrs), filename)
	return acrs
}

// #
// TODO docs
// #
func runGeneratingAcronymsFromSourceMode() Acronyms {
	fmt.Println("> Generating acronyms from source:")

	// Give a choice of source file
	yesOrNo, err := cio.GiveUserYesOrNoChoice(UseDefaultSrcFileChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return nil
	}
	var srcFilename string
	if yesOrNo == cio.Yes {
		srcFilename = SrcDefaultFilename
	} else if yesOrNo == cio.No {
		srcFilename, err = cio.GiveUserChoiceOfFilename("> Enter a name of source file:\n")
		if err != nil {
			formatAndPrintError(err)
			return nil
		}
	}

	// Load source data from file.
	src, err := loadSrcFromFile(srcFilename)
	if err != nil {
		formatAndPrintError(err)
		return nil
	}

	// Give a choice of dictionary file
	yesOrNo, err = cio.GiveUserYesOrNoChoice(UseDefaultDictFileChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return nil
	}
	var dictFilename string
	if yesOrNo == cio.Yes {
		dictFilename = DictDefaultFilename
	} else if yesOrNo == cio.No {
		dictFilename, err = cio.GiveUserChoiceOfFilename("> Enter a name of dictionary file:\n")
		if err != nil {
			formatAndPrintError(err)
			return nil
		}
	}

	// Load dictionary from file.
	const ExpectedWordsAmount = 1532570 // 1'532'568 = amount of russian words in my collection
	dict, err := loadDictionaryFromFile(dictFilename, ExpectedWordsAmount)
	if err != nil {
		formatAndPrintError(err)
		return nil
	}

	// Give a choice of program mode: generate acronyms with or without strict order.
	yesOrNo, err = cio.GiveUserYesOrNoChoice(AcrGenerationModeChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return nil
	}
	var mode AcrGeneratorMode
	if yesOrNo == cio.Yes {
		mode = Ordered
	} else if yesOrNo == cio.No {
		mode = NonOrdered
	}

	// Generate and sort acronyms.
	acrs := generateAcronyms(src, dict, mode)
	SortAcronymsByAlphabet(acrs)
	fmt.Printf("\n> %d acronyms were successfully generated and sorted by alphabet.\n", len(acrs))
	return acrs
}

// #
// TODO docs
// #
func runListOfAcronymsPrintingMode(acrs Acronyms) {
	fmt.Println("> Printing acronyms in console:")

	if acrs == nil {
		formatAndPrintError(errors.New("unexpected empty acronym collection"))
		return
	}

	amount, err := cio.GiveUserNumberChoice(AmountOfAcronymsToBePrintedChoiceMes, IncorrectNumberChoiceMes)
	if err != nil {
		formatAndPrintError(err)
		return
	}

	err = printMostSuitableAcronyms(acrs, amount)
	if err != nil {
		formatAndPrintError(err)
		return
	}
}

// #
// TODO docs
// #
func runAcronymsDecodingMode(acrs Acronyms) {
	invitingLine := fmt.Sprintf("> Acronyms decoding (use \"%s\" to quit from this mode):\n", QuitModeCommand)

	if acrs == nil {
		formatAndPrintError(errors.New("unexpected empty acronym collection"))
		return
	}

	containsAcronymWrap := func(userInp string) (bool, error) {
		_, ok := containsAcronymBS(userInp, acrs)
		return ok, nil
	}
	takeAndPrintAcronym := func(userInp string) error {
		// TODO maybe reuse index that was found in 'containsAcronymWrap'
		acr, _ := takeAcronymBS(userInp, acrs) // No need to check 'ok': we've just checked that acronym is in collection
		printAcronymInDetail(acr)
		fmt.Printf("\n")
		return nil
	}

	err, _ := cio.ProcessUserInputUntilExitCommand(
		invitingLine,
		"Please, enter an acronym:",
		"No such acronym was found.\n",
		QuitModeCommand,
		containsAcronymWrap,
		takeAndPrintAcronym)

	if err != nil {
		formatAndPrintError(err)
		return
	}
}

// #
// TODO docs
// #
func runSavingAcronymsToFileMode(acrs Acronyms) {
	fmt.Println("> Saving acronyms to file:")

	// Give a choice of output file
	yesOrNo, err := cio.GiveUserYesOrNoChoice(UseDefaultOutputFileChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return
	}
	var filename string
	if yesOrNo == cio.Yes {
		filename = OutputDefaultFilename
	} else if yesOrNo == cio.No {
		filename, err = cio.GiveUserChoiceOfFilename("> Enter a name of output file:\n")
		if err != nil {
			formatAndPrintError(err)
			return
		}
	}

	err = exportAcronymsToFile(acrs, filename, FullFormat)
	if err != nil {
		formatAndPrintError(err)
		return
	}
}

// #
// Processes incorrect user menu command and tries to help or guess the meaning of the user's input.
// #
func processIncorrectUserMenuCommand(userInp string) {
	if userInp == "!h" {
		fmt.Printf("> Maybe you mean \"%s\"? Try again.\n", HelpCommand)
	} else {
		fmt.Println("Incorrect input!")
	}
}
