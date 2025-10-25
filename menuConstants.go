package main

import (
	"fmt"
)

// Default filenames enumeration
const (
	DumpDefaultFilename   = "acrs_dump.txt"
	SrcDefaultFilename    = "data/src.txt"
	DictDefaultFilename   = "data/russian_words.txt"
	OutputDefaultFilename = "acrs.txt"
)

// User menu commands
type MenuCommand = string

const (
	HelpCommand        = "!H"
	ExitProgramCommand = "!Q"
	QuitModeCommand    = "!q"

	LoadAcronymsFromFileCommand       = "!1"
	GenerateAcronymsFromSourceCommand = "!2"
	PrintListOfAcronymsCommand        = "!3"
	DecodeAcronymCommand              = "!4"
	SaveAcronymsToFileCommand         = "!5"
)

type void = struct{}

// All menu commands in set
var AllMenuCommands = map[MenuCommand]void{
	HelpCommand:        void{},
	ExitProgramCommand: void{},
	QuitModeCommand:    void{},

	LoadAcronymsFromFileCommand:       void{},
	GenerateAcronymsFromSourceCommand: void{},
	PrintListOfAcronymsCommand:        void{},
	DecodeAcronymCommand:              void{},
	SaveAcronymsToFileCommand:         void{},
}

// Checks if 'str' is valid menu command
func isValidMenuCommand(str string) bool {
	_, isExisting := AllMenuCommands[str]
	return isExisting
}

// Menu messages
const (
	UserConfirmExitMes = "Are you sure about exiting?"

	UserChoiceInputFormatErrMes = "Unexpected choice (incorrect input format)."
	IncorrectNumberChoiceMes    = "Unexpected choice (a number was expected)."

	UseDefaultDumpFileChoiceMes   = "Use default file name ('" + DumpDefaultFilename + "')?"
	UseDefaultSrcFileChoiceMes    = "Use default file name ('" + SrcDefaultFilename + "')?"
	UseDefaultDictFileChoiceMes   = "Use default file name ('" + DictDefaultFilename + "')?"
	UseDefaultOutputFileChoiceMes = "Use default file name ('" + OutputDefaultFilename + "')?"

	AcrGenerationModeChoiceMes = "Must items in acronym be ordered?"

	AmountOfAcronymsToBePrintedChoiceMes = "Choose the number of acronyms for printing (0 for all)."
)

// #
// Prints a list of available menu commands and modes (with some helper info).
// #
func printMenuInfo() {
	fmt.Printf("\n>>> Menu (enter commands without quotes):\n")
	fmt.Printf("\n")
	fmt.Printf("  * Help -                          \"%s\"\n", HelpCommand)
	fmt.Printf("  * Exit 'acrogen' program -        \"%s\"\n", ExitProgramCommand)
	fmt.Printf("  * Quit current mode -             \"%s\"\n", QuitModeCommand)
	fmt.Printf("\n")
	fmt.Printf("~ Initial commands:\n")
	fmt.Printf("  * Load acronyms from file -       \"%s\"\n", LoadAcronymsFromFileCommand)
	fmt.Printf("  * Generate acronyms from source - \"%s\"\n", GenerateAcronymsFromSourceCommand)
	fmt.Printf("\n")
	fmt.Printf("~ Commands available after loading or generating:\n")
	fmt.Printf("  * Print list of acronyms -        \"%s\"\n", PrintListOfAcronymsCommand)
	fmt.Printf("  * Decode single acronym -         \"%s\"\n", DecodeAcronymCommand)
	fmt.Printf("  * Save acronyms to file -         \"%s\"\n", SaveAcronymsToFileCommand)
	fmt.Printf("\n")
}
