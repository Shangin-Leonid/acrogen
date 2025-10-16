package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"acrgen/fio"
)

const LetterOptsSeparator = ""
const LetterOptSeparator = " -- "

// #
// Parses source data file and import its content.
// #
func importSrcFromFile(srcFilename string) (Src, error) {
	src := make(Src, 0, 10)
	src = append(src, make(LetterOpts, 0, 10))

	var parseSrcFileLine fio.StringParserFunc = func(line string) error {
		if line == LetterOptsSeparator {
			if len(src[len(src)-1]) == 0 {
				return errors.New("incorrect format of input file: first (initial) or multiple consecutive blank lines are prohibited")
			}
			src = append(src, make(LetterOpts, 0, 10))
			return nil
		}

		splittedLine := strings.Split(line, LetterOptSeparator)
		if len(splittedLine) != 3 {
			return errors.New("incorrect format of input file: unexpected data format error during reading the file")
		}

		letterToken := []rune(splittedLine[0])
		if len(letterToken) != 1 || !unicode.IsLetter(letterToken[0]) {
			return errors.New("incorrect format of input file. first token is not a letter")
		}
		letter := unicode.ToLower(letterToken[0])

		estimation, err := strconv.Atoi(splittedLine[1])
		if err != nil {
			return errors.New("incorrect format of input file. second token is not a number or incorrect number")
		}

		decoding := splittedLine[2]

		src[len(src)-1] = append(src[len(src)-1], LetterOpt{letter, estimation, decoding})
		return nil
	}

	_, err := fio.ParseTextFileLineByLine(srcFilename, nil, parseSrcFileLine)

	if err != nil {
		return nil, err
	}

	if len(src[len(src)-1]) == 0 {
		src = src[:len(src)-1]
	}

	return src, nil
}

// #
// Parses dictionary file (list of valid words) and import its content.
// #
func importDictionaryFromFile(dictFilename string, expectedWordsAmount uint64) (Dict, error) {
	dict := make(Dict, expectedWordsAmount)

	var parseWordFromFileLine fio.StringParserFunc = func(line string) error {
		dict[line] = struct{}{}
		return nil
	}

	_, err := fio.ParseTextFileLineByLine(dictFilename, nil, parseWordFromFileLine)

	if err != nil {
		return nil, err
	}

	return dict, nil
}

// Enumeration represents mode of acronyms file export.
type ExportModeT int

const (
	FullFormat ExportModeT = iota + 1
	OnelineFormat
)

// #
// Export acronyms to output file in short format (without letters decoding, but each acronym is on new line).
// #
func exportAcronymsToFile(acrs Acronyms, outputFilename string, mode ExportModeT) error {
	var formatFunc func(acr Acronym) string

	if mode == FullFormat {
		formatFunc = func(acr Acronym) string {
			outp := acr.word + LetterOptSeparator + strconv.Itoa(acr.sumEstimation) + "\n"
			// TODO optimize by switching from string to []rune
			for i, letter := range []rune(acr.word) {
				outp += string(letter) + LetterOptSeparator + acr.letterDecodings[i] + "\n"
			}
			outp += LetterOptsSeparator + "\n"
			return outp
		}
	} else if mode == OnelineFormat {
		formatFunc = func(acr Acronym) string {
			return acr.word + "\n"
		}
	}

	_, err := fio.WriteSliceToTextFile(acrs, outputFilename, formatFunc)
	return err
}

// #
// Prints acronyms in console in poor format (acroonym only, without any decoding info).
// #
func printAcronyms(acrs Acronyms) {
	fmt.Printf("\nList of acronyms:\n")
	for i := range acrs {
		fmt.Println(acrs[i].word)
	}
	fmt.Printf("\n")
}
