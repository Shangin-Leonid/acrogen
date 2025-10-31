package ui /* User Interface */

import (
	"github.com/fatih/color"
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

// Prefix before menu message
const (
	MenuPrefix    = ">>>"
	MessagePrefix = ">"
)

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

// Color themes
var MenuColor *color.Color = color.New(color.FgYellow, color.Bold, color.Faint)
var SuccessColor *color.Color = color.New(color.FgGreen, color.Bold)
var WarningColor *color.Color = color.RGB(255, 165, 0).Add(color.Bold) // Orange
var ErrorColor *color.Color = color.New(color.FgRed, color.Bold)

// #
// Prints a list of available menu commands and modes (with some helper info).
// #
func printMenuInfo() {
	MenuColor.Printf("\n%s Menu (enter commands without quotes):\n", MenuPrefix)
	MenuColor.Printf("\n")
	MenuColor.Printf("  * Help -                          \"%s\"\n", HelpCommand)
	MenuColor.Printf("  * Exit 'acrogen' program -        \"%s\"\n", ExitProgramCommand)
	MenuColor.Printf("  * Quit current mode -             \"%s\"\n", QuitModeCommand)
	MenuColor.Printf("\n")
	MenuColor.Printf("~ Initial commands:\n")
	MenuColor.Printf("  * Load acronyms from file -       \"%s\"\n", LoadAcronymsFromFileCommand)
	MenuColor.Printf("  * Generate acronyms from source - \"%s\"\n", GenerateAcronymsFromSourceCommand)
	MenuColor.Printf("\n")
	MenuColor.Printf("~ Commands available after loading or generating:\n")
	MenuColor.Printf("  * Print list of acronyms -        \"%s\"\n", PrintListOfAcronymsCommand)
	MenuColor.Printf("  * Decode single acronym -         \"%s\"\n", DecodeAcronymCommand)
	MenuColor.Printf("  * Save acronyms to file -         \"%s\"\n", SaveAcronymsToFileCommand)
	MenuColor.Printf("\n")
}

// Data format in console input-output
const TokenSeparator = " -- "
const LineSeparator = ""
