package ui /* User Interface */

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"acrogen/ag"
	"acrogen/fio"
)

// #
// Runs the 'acrogen' application.
// The inner point of the program.
// TODO '!Q' in decoding mode must exit the program
// TODO ask user if he is sure before exit the program
// TODO '!H' in decoding mode must work
// #
func RunConsoleApp() {
	var err error
	var acrs ag.Acronyms

	printMenuInfo()
	var userInp string
	for {
		MenuColor.Printf("%s Enter a command:\n", MessagePrefix)

		_, err = fmt.Scanf("%s", &userInp)
		if err != nil {
			formatAndPrintError(err)
			return
		}

		switch userInp {
		case HelpCommand:
			runHelpMode()
		case ExitProgramCommand:
			if runTryOfExiting() {
				return
			}
		case QuitModeCommand:
			WarningColor.Printf("%s You are in the menu - nothing to quit from.\n", MessagePrefix)
		case LoadAcronymsFromFileCommand:
			acrs = runLoadingAcronymsFromFileMode()
		case GenerateAcronymsFromSourceCommand:
			acrs = runGeneratingAcronymsFromSourceMode()
		case PrintListOfAcronymsCommand:
			runListOfAcronymsPrintingMode(acrs)
		case DecodeAcronymCommand:
			runAcronymsDecodingMode(acrs)
		case SaveAcronymsToFileCommand:
			runSavingAcronymsToFileMode(acrs)
		default:
			processInvalidUserMenuCommand(userInp)
		}

		MenuColor.Printf("\n")
	}
}

// #
// Formats and prints error in console in 'stderr'.
// #
func formatAndPrintError(err error) {
	ErrorColor.Fprintln(os.Stderr, fmt.Errorf("%s Error: %w.", MenuPrefix, err))
}

// #
// Helps a user to apply the 'acrogen'.
// #
func runHelpMode() {
	// TODO
	printMenuInfo()
}

// #
// Asks user to confirm exiting. Prints something before return positive exiting flag.
// Returns 'true' if need to exit program.
// #
func runTryOfExiting() (needExit bool) {
	yesOrNo, err := giveUserYesOrNoChoice(UserConfirmExitMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return true
	}

	if yesOrNo == No {

		return false
	}

	SuccessColor.Printf("\n%s \"Acrogen\" (\"%s\") finished with success.\n\n", MenuPrefix, os.Args[0])
	return true
}

// #
// TODO docs
// #
func runLoadingAcronymsFromFileMode() ag.Acronyms {
	MenuColor.Printf("%s Loading acronyms from file:\n", MenuPrefix)

	// Give a choice of input file
	yesOrNo, err := giveUserYesOrNoChoice(UseDefaultDumpFileChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return nil
	}
	var filename string
	if yesOrNo == Yes {
		filename = DumpDefaultFilename
	} else if yesOrNo == No {
		filename, err = giveUserChoiceOfFilename("Enter a name of file:")
		if err != nil {
			formatAndPrintError(err)
			return nil
		}
	}

	// Load acronyms from file
	acrs, err := fio.LoadAcronymsFromFile(filename)
	if err != nil {
		formatAndPrintError(err)
		return nil
	}
	SuccessColor.Printf("\n%s %d acronyms have been successfully loaded from '%s'.\n", MessagePrefix, len(acrs), filename)
	return acrs
}

// #
// TODO docs
// #
func runGeneratingAcronymsFromSourceMode() ag.Acronyms {
	MenuColor.Printf("%s Generating acronyms from source:\n", MenuPrefix)

	// Give a choice of source file
	yesOrNo, err := giveUserYesOrNoChoice(UseDefaultSrcFileChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return nil
	}
	var srcFilename string
	if yesOrNo == Yes {
		srcFilename = SrcDefaultFilename
	} else if yesOrNo == No {
		srcFilename, err = giveUserChoiceOfFilename("Enter a name of source file:")
		if err != nil {
			formatAndPrintError(err)
			return nil
		}
	}

	// Load source data from file.
	src, err := fio.LoadSrcFromFile(srcFilename)
	if err != nil {
		formatAndPrintError(err)
		return nil
	}

	// Give a choice of dictionary file
	yesOrNo, err = giveUserYesOrNoChoice(UseDefaultDictFileChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return nil
	}
	var dictFilename string
	if yesOrNo == Yes {
		dictFilename = DictDefaultFilename
	} else if yesOrNo == No {
		dictFilename, err = giveUserChoiceOfFilename("Enter a name of dictionary file:")
		if err != nil {
			formatAndPrintError(err)
			return nil
		}
	}

	// Load dictionary from file.
	const ExpectedWordsAmount = 1532570 // 1'532'568 = amount of russian words in my collection
	dict, err := fio.LoadDictionaryFromFile(dictFilename, ExpectedWordsAmount)
	if err != nil {
		formatAndPrintError(err)
		return nil
	}

	// Give a choice of program mode: generate acronyms with or without strict order.
	yesOrNo, err = giveUserYesOrNoChoice(AcrGenerationModeChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return nil
	}
	var mode ag.AcrGeneratorMode
	if yesOrNo == Yes {
		mode = ag.Ordered
	} else if yesOrNo == No {
		mode = ag.NonOrdered
	}

	// Generate and sort acronyms.
	acrs := ag.GenerateAcronyms(src, dict, mode)
	ag.SortAcronymsByAlphabet(acrs)
	SuccessColor.Printf("\n%s %d acronyms were successfully generated and sorted by alphabet.\n", MessagePrefix, len(acrs))
	return acrs
}

// #
// TODO docs
// #
func runListOfAcronymsPrintingMode(acrs ag.Acronyms) {
	MenuColor.Printf("%s Printing acronyms in console:\n", MenuPrefix)

	if acrs == nil {
		formatAndPrintError(errors.New("unexpected empty acronym collection"))
		return
	}

	amount, err := giveUserNumberChoice(AmountOfAcronymsToBePrintedChoiceMes, IncorrectNumberChoiceMes)
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
func runAcronymsDecodingMode(acrs ag.Acronyms) {
	invitingLine := fmt.Sprintf("%s Acronyms decoding (use \"%s\" to quit from this mode):\n", MenuPrefix, QuitModeCommand)

	if acrs == nil {
		formatAndPrintError(errors.New("unexpected empty acronym collection"))
		return
	}

	containsAcronymWrap := func(userInp string) (bool, error) {
		_, ok := ag.ContainsAcronymBS(userInp, acrs)
		return ok, nil
	}
	takeAndPrintAcronym := func(userInp string) error {
		// TODO maybe reuse index that was found in 'containsAcronymWrap'
		acr, _ := ag.TakeAcronymBS(userInp, acrs) // No need to check 'ok': we've just checked that acronym is in collection
		printAcronymInDetail(acr)
		fmt.Printf("\n")
		return nil
	}

	err, _ := processUserInputUntilExitCommand(
		QuitModeCommand,
		invitingLine,
		fmt.Sprintf("%s Please, enter an acronym:", MessagePrefix),
		fmt.Sprintf("%s No such acronym was found.\n", MessagePrefix),
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
func runSavingAcronymsToFileMode(acrs ag.Acronyms) {
	MenuColor.Printf("%s Saving acronyms to file:\n", MenuPrefix)

	// Give a choice of output file
	yesOrNo, err := giveUserYesOrNoChoice(UseDefaultOutputFileChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		formatAndPrintError(err)
		return
	}
	var filename string
	if yesOrNo == Yes {
		filename = OutputDefaultFilename
	} else if yesOrNo == No {
		filename, err = giveUserChoiceOfFilename("Enter a name of output file:")
		if err != nil {
			formatAndPrintError(err)
			return
		}
	}

	err = fio.ExportAcronymsToFile(acrs, filename, fio.FullFormat)
	if err != nil {
		formatAndPrintError(err)
		return
	}
}

// #
// Processes incorrect user menu command and tries to help or guess the meaning of the user's input.
// #
func processInvalidUserMenuCommand(userInp string) {
	spaceFreeUserInp := strings.ReplaceAll(userInp, " ", "")

	if isValidMenuCommand(spaceFreeUserInp) {
		WarningColor.Printf("%s Unexpected spaces. Maybe you mean \"%s\"? Try again.\n", MessagePrefix, spaceFreeUserInp)
	} else if isValidMenuCommand("!" + spaceFreeUserInp) {
		WarningColor.Printf("%s Maybe you mean \"%s\"? Try again.\n", MessagePrefix, "!"+spaceFreeUserInp)
	} else if userInp == "!h" {
		WarningColor.Printf("%s Maybe you mean \"%s\"? Try again.\n", MessagePrefix, HelpCommand)
	} else {
		WarningColor.Printf("%s Incorrect command. Try again.\n", MessagePrefix)
	}
}
