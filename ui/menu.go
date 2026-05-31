package ui /* User Interface */

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"

	"github.com/Shangin-Leonid/acrogen/ag"
	"github.com/Shangin-Leonid/acrogen/fio"
	"github.com/Shangin-Leonid/acrogen/utils"
)

// # RunConsoleApp runs the whole 'acrogen' application (the inner point).
//
// # Params: -
//
// # Returns: -
//
// # Description:
//
// Covers panic, handles errors.
//
// # TODOs:
//
//   - '!Q' in decoding mode must exit the program.
//   - Ask user if he is sure before exit the program.
//   - '!H' in decoding mode must work.
func RunConsoleApp() {

	defer func() {
		if p := recover(); p != nil {
			defer os.Exit(1)
			log.Printf("PANIC: %v\n", p)
			log.Printf("Stack trace:\n%s", debug.Stack())
			ErrorColor.Printf("PANIC: %v\n", p)
		}
	}()

	var err error
	var acrs ag.Acronyms

	printMenuInfo()

	var userInp string
	var needExit bool
	for {
		MenuColor.Println(AskUserToEnterAnyCommandMes)

		_, err = fmt.Scanf("%s", &userInp)
		if err != nil {
			panic(err.Error())
		}

		switch userInp {
		case HelpCommand:
			runHelpMode()
		case ExitProgramCommand:
			needExit, err = runTryOfExiting()
		case QuitModeCommand:
			WarningColor.Println(WrongAttemptOfExitingFromMenuMes)
		case LoadAcronymsFromFileCommand:
			acrs, err = runLoadingAcronymsFromFileMode()
		case GenerateAcronymsFromSourceCommand:
			acrs, err = runGeneratingAcronymsFromSourceMode()
		case PrintListOfAcronymsCommand:
			err = runListOfAcronymsPrintingMode(acrs)
		case DecodeAcronymCommand:
			err = runAcronymsDecodingMode(acrs)
		case SaveAcronymsToFileCommand:
			err = runSavingAcronymsToFileMode(acrs)
		default:
			processInvalidUserMenuCommand(userInp)
		}

		printAndLogErrorIfExists(err)

		if needExit {
			return
		}

		MenuColor.Println()
	}
}

// # printAndLogErrorIfExists formats, logs and prints error if it exists.
//
// # Params: -
//
// # Returns: -
//
// # Description:
func printAndLogErrorIfExists(err error) {
	if err == nil {
		return
	}

	log.Printf("Error: %v.\n", err)
	if ste := utils.NewSTError(""); errors.As(err, &ste) {
		log.Printf("Stack trace of error:\n %s\n\n", ste.StackTrace())
	}

	ErrorColor.Fprintf(os.Stdout, "%s Error: %v.", MenuPrefix, err)
}

// # runHelpMode prints a hint of menu.
//
// # Params: -
//
// # Returns: -
//
// # Description:
//
// Helps a user to apply the 'acrogen'.
func runHelpMode() {
	printMenuInfo()
}

// # runTryOfExiting asks user to confirm exiting and returns result of his input.
//
// # Params: -
//
// # Returns:
//
//   - flag if need to exit application
//   - error
//
// # Description:
//
// Prints "bye-bye" message if need.
func runTryOfExiting() (needExit bool, _ error) {
	yesOrNo, err := giveUserYesOrNoChoice(UserConfirmExitMes, UserChoiceInputFormatErrMes)
	if err != nil {
		return true, err
	}

	if yesOrNo == No {
		return false, nil
	}

	SuccessColor.Printf("\n%s \"Acrogen\" (\"%s\") finished with success.\n\n", MenuPrefix, os.Args[0])
	return true, nil
}

// # runLoadingAcronymsFromFileMode tries to load acronyms from file (default or user's).
//
// # Params: -
//
// # Returns:
//
//   - acronyms collection (nil if error)
//   - error
func runLoadingAcronymsFromFileMode() (ag.Acronyms, error) {
	MenuColor.Printf("\n%s Loading acronyms from file:\n", MenuPrefix)

	// Give a choice of input file
	yesOrNo, err := giveUserYesOrNoChoice(UseDefaultDumpFileChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		return nil, err
	}
	var filename string
	if yesOrNo == Yes {
		filename = DumpDefaultFilename
	} else if yesOrNo == No {
		filename, err = giveUserChoiceOfFilename("Enter a name of file:")
		if err != nil {
			return nil, err
		}
	}

	// Load acronyms from file
	acrs, err1 := fio.LoadAcronymsFromFile(filename)
	fmt.Println(err1)
	if err1 != nil {
		return nil, err
	}
	SuccessColor.Printf("\n%s %d acronyms have been successfully loaded from '%s'.\n", MessagePrefix, len(acrs), filename)
	return acrs, nil
}

// # runGeneratingAcronymsFromSourceMode tries to generate acronyms from source and dictionary files (default or user's)
//
// # Params: -
//
// # Returns:
//
//   - acronyms collection (nil if error)
//   - error
func runGeneratingAcronymsFromSourceMode() (ag.Acronyms, error) {
	MenuColor.Printf("\n%s Generating acronyms from source:\n", MenuPrefix)

	// Give a choice of source file
	yesOrNo, err := giveUserYesOrNoChoice(UseDefaultSrcFileChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		return nil, err
	}
	var srcFilename string
	if yesOrNo == Yes {
		srcFilename = SrcDefaultFilename
	} else if yesOrNo == No {
		srcFilename, err = giveUserChoiceOfFilename("Enter a name of source file:")
		if err != nil {
			return nil, err
		}
	}

	// Load source data from file.
	src, err := fio.LoadSrcFromFile(srcFilename)
	if err != nil {
		return nil, err
	}

	// Give a choice of dictionary file
	yesOrNo, err = giveUserYesOrNoChoice(UseDefaultDictFileChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		return nil, err
	}
	var dictFilename string
	if yesOrNo == Yes {
		dictFilename = DictDefaultFilename
	} else if yesOrNo == No {
		dictFilename, err = giveUserChoiceOfFilename("Enter a name of dictionary file:")
		if err != nil {
			return nil, err
		}
	}

	// Load dictionary from file.
	const ExpectedWordsAmount = 1532570 // 1'532'568 = amount of russian words in my collection
	dict, err := fio.LoadDictionaryFromFile(dictFilename, ExpectedWordsAmount)
	if err != nil {
		return nil, err
	}

	// Give a choice of program mode: generate acronyms with or without strict order.
	yesOrNo, err = giveUserYesOrNoChoice(AcrGenerationModeChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		return nil, err
	}
	mode := utils.TerOp(yesOrNo == Yes, ag.Ordered, ag.NonOrdered)

	// Generate and sort acronyms.
	acrs := ag.GenerateAcronyms(src, dict, mode)
	ag.SortAcronymsByAlphabet(acrs)
	SuccessColor.Printf("\n%s %d acronyms were successfully generated and sorted by alphabet.\n", MessagePrefix, len(acrs))
	return acrs, nil
}

// # runListOfAcronymsPrintingMode prints acronyms.
//
// # Params:
//
//   - collection of acronyms to print
//
// # Returns:
//
//   - error
//
// # Description:
//
// Gives a user a choice of acronyms amount.
func runListOfAcronymsPrintingMode(acrs ag.Acronyms) error {
	MenuColor.Printf("\n%s Printing acronyms in console:\n", MenuPrefix)

	if acrs == nil {
		return utils.NewSTError("unexpected empty acronym collection")
	} else if len(acrs) == 0 {
		WarningColor.Printf("\n%s No acronyms were found\n", MenuPrefix)
		return nil
	}

	if acrs == nil {
		return utils.NewSTError("unexpected empty acronym collection")
	}

	amount, err := giveUserNumberChoice(AmountOfAcronymsToBePrintedChoiceMes, IncorrectNumberChoiceMes)
	if err != nil {
		return err
	}

	err = printMostSuitableAcronyms(acrs, amount)
	if err != nil {
		return err
	}

	return nil
}

// # runAcronymsDecodingMode decodes acronyms.
//
// # Params:
//
//   - collection of acronyms to decode
//
// # Returns:
//
//   - error
//
// # TODOs:
//
//   - Maybe reuse index that was found in 'containsAcronymWrap' for 'takeAndPrintAcronym'
func runAcronymsDecodingMode(acrs ag.Acronyms) error {
	invitingLine := fmt.Sprintf("\n%s Acronyms decoding (use \"%s\" to quit from this mode):\n", MenuPrefix, QuitModeCommand)

	if acrs == nil {
		return utils.NewSTError("unexpected empty acronym collection")
	} else if len(acrs) == 0 {
		WarningColor.Printf("\n%s No acronyms were found\n", MenuPrefix)
		return nil
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
		return err
	}

	return nil
}

// # runSavingAcronymsToFileMode saves acronyms to output file (default or user's).
//
// # Params:
//
//   - collection of acronyms to save
//
// # Returns:
//
//   - error
func runSavingAcronymsToFileMode(acrs ag.Acronyms) error {
	MenuColor.Printf("\n%s Saving acronyms to file:\n", MenuPrefix)

	if acrs == nil {
		return utils.NewSTError("unexpected empty acronym collection")
	} else if len(acrs) == 0 {
		WarningColor.Printf("\n%s No acronyms were found\n", MenuPrefix)
		return nil
	}

	// Give a choice of output file
	yesOrNo, err := giveUserYesOrNoChoice(UseDefaultOutputFileChoiceMes, UserChoiceInputFormatErrMes)
	if err != nil {
		return err
	}
	var filename string
	if yesOrNo == Yes {
		filename = OutputDefaultFilename
	} else if yesOrNo == No {
		filename, err = giveUserChoiceOfFilename("Enter a name of output file:")
		if err != nil {
			return err
		}
	}

	// Save acronyms to file
	err = fio.SaveAcronymsToFile(acrs, filename, fio.FullFormat)
	if err != nil {
		return err
	}

	SuccessColor.Printf("%s Acronyms have been successfully saved to file '%s'.\n", MessagePrefix, filename)

	return nil
}

// # processInvalidUserMenuCommand processes incorrect user menu command and tries to help.
//
// # Params:
//
//   - user's input as string
//
// # Returns: -
//
// # Description:
//
// Tries to guess the meaning of the user's incorrect input.
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
