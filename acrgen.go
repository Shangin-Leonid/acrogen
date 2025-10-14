package main

import (
	"fmt"
	"os"
)

// TODO switch the language of messages for user
func main() {
	argsWithoutProgName := os.Args[1:]
	if len(argsWithoutProgName) != 3 {
		fmt.Println("Неверное количество входных аргументов!")
		fmt.Println("Запустите программу заново, указав названия трёх  \".txt\" файлов: входного, с существующими словами-кандидатами и выходного")
		return
	}
	srcFilename, dictFilename, outputFilename := argsWithoutProgName[0], argsWithoutProgName[1], argsWithoutProgName[2]

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

	acrs, err := generateAcronyms(src, dict)
	if err != nil {
		formatAndPrintError(err)
		return
	}

	err = exportAcronymsToFile(acrs, outputFilename)
	if err != nil {
		formatAndPrintError(err)
		return
	}

	printAcronyms(acrs)

	const DecodeChoiceMessage = "Would you like to decode any generated acronym?"
	const UserInputFormatErrMessage = "Incorrect choice (incorrect input format)."
	yesOrNo, err := giveUserAChoiceYesOrNo(DecodeChoiceMessage, UserInputFormatErrMessage)
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
	fmt.Errorf("Error: %w", err)
}
