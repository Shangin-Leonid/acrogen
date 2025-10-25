package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"acrogen/fio"
)

const TokenSeparator = " -- "
const LineSeparator = ""

// #
// Parses source data file and load its content.
// Returns 'Src' collection and error.
// #
func loadSrcFromFile(srcFilename string) (Src, error) {
	src := make(Src, 0, 10)
	src = append(src, make(LetterOpts, 0, 10))

	var parseSrcFileLine fio.StringParserFunc = func(line string) error {
		if line == LineSeparator {
			if len(src[len(src)-1]) == 0 {
				return errors.New("incorrect format of input file: first (initial) or multiple consecutive blank lines are prohibited")
			}
			src = append(src, make(LetterOpts, 0, 10))
			return nil
		}

		splittedLine := strings.Split(line, TokenSeparator)
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
// Parses dictionary file (list of valid words) and load its content.
// Returns dictionary and error.
// #
func loadDictionaryFromFile(dictFilename string, expectedWordsAmount uint64) (Dict, error) {
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
			outp := acr.word + TokenSeparator + strconv.Itoa(acr.sumEstimation) + "\n"
			for i, letter := range []rune(acr.word) {
				outp += string(letter) + TokenSeparator + acr.letterDecodings[i] + "\n"
			}
			outp += LineSeparator + "\n"
			return outp
		}
	} else if mode == OnelineFormat {
		formatFunc = func(acr Acronym) string {
			return acr.word + "\n"
		}
	}

	_, err := fio.WriteSliceToTextFile(acrs, outputFilename, true, formatFunc)
	return err
}

// #
// Import acronyms from file with 'FullFormat'.
// #
func loadAcronymsFromFile(filename string) (acrs Acronyms, err error) {

	var parseFirstLineInFile fio.StringParserFunc = func(line string) error {
		if len(line) == 0 {
			return errors.New("unexpected empty first line in file")
		}
		capacity, err := strconv.Atoi(line)
		if err == nil {
			acrs = make(Acronyms, 0, capacity)
		}
		return err
	}

	// Enumeration represents the type of previous parsed line.
	type LineT int
	const (
		First         LineT = 1
		Empty               = 2
		Acr                 = 3
		Letter              = 4
		LastAcrLetter       = 5
	)
	isCurLineTypeCorrect := func(cur, prev LineT) bool {
		switch cur {
		case Empty:
			return prev == First || prev == LastAcrLetter
		case Acr:
			return prev == Empty
		case Letter:
			return prev == Acr || prev == Letter
		case LastAcrLetter:
			return prev == Letter
		}
		return false
	}

	var prev LineT = First
	var cur LineT
	// TODO refactor, simplify
	var parseAcronymsInFile fio.StringParserFunc = func(line string) error {
		lineRunes := []rune(line)

		if line == LineSeparator {
			cur = Empty
			if !isCurLineTypeCorrect(cur, prev) {
				return errors.New("incorrect data/format: unexpected empty line")
			}
			if prev != First {
				lastLD := acrs[len(acrs)-1].letterDecodings
				if prev == Letter && len(lastLD) != cap(lastLD) {
					return errors.New("incorrect data/format: unexpected empty line within acronyms description")
				}
			}
		} else if len(lineRunes) < 2 {
			return errors.New("incorrect data/format: some short line with no unexpected meaning")
		} else if unicode.IsLetter(lineRunes[0]) && !unicode.IsLetter(lineRunes[1]) {
			cur = Letter
			if !isCurLineTypeCorrect(cur, prev) {
				return errors.New("incorrect data/format: maybe unexpected letter decoding line")
			}

			if !unicode.IsLower(lineRunes[0]) {
				return errors.New("incorrect format: the letter is not lowercase")
			}

			curAcrAsRune := []rune(acrs[len(acrs)-1].word)
			curLetterInd := len(acrs[len(acrs)-1].letterDecodings)
			if lineRunes[0] != curAcrAsRune[curLetterInd] {
				return errors.New("incorrect data: the letter is not the same as in the acronym")
			}

			tsInd := strings.Index(line, TokenSeparator)
			if tsInd == -1 {
				return errors.New("incorrect data/format: incorrect format of letter decoding line (no token separator between letter and decoding)")
			}
			decodingInd := tsInd + len(TokenSeparator)
			if len(acrs[len(acrs)-1].letterDecodings) == cap(acrs[len(acrs)-1].letterDecodings) {
				return errors.New("unexpected error: maybe unexpected (extra) letter in acronym")
			}
			acrs[len(acrs)-1].letterDecodings = append(acrs[len(acrs)-1].letterDecodings, line[decodingInd:])

			if curLetterInd == len(curAcrAsRune)-1 {
				cur = LastAcrLetter
			}
		} else {
			cur = Acr
			if !isCurLineTypeCorrect(cur, prev) {
				return errors.New("incorrect data/format: unexpected acronym line or smth else")
			}

			tsInd := strings.Index(line, TokenSeparator)
			if tsInd == -1 {
				return errors.New("incorrect data/format: incorrect format of acronym line (no token separator between word and estimation)")
			}
			estInd := tsInd + len(TokenSeparator)
			est, err := strconv.Atoi(line[estInd:])
			if err != nil {
				return errors.New("incorrect data/format: incorrect summary estimation of acronym")
			}

			nLetters := 0
			for _, l := range line[:tsInd] {
				if !unicode.IsLetter(l) || !unicode.IsLower(l) {
					return errors.New("incorrect data/format: not letters or upper case letters in acronym")
				}
				nLetters++
			}
			acrs = append(acrs, Acronym{line[:tsInd], est, make([]string, 0, nLetters)})
		}

		prev = cur
		return nil
	}

	_, err = fio.ParseTextFileLineByLine(filename, parseFirstLineInFile, parseAcronymsInFile)

	return acrs, err
}

// #
// Prints acronym in console in detailed format (decodes each letter).
// #
func printAcronymInDetail(acr Acronym) {
	fmt.Printf("%s%s%d\n", acr.word, TokenSeparator, acr.sumEstimation)
	for i, letter := range []rune(acr.word) {
		fmt.Printf("%c -- %s\n", letter, acr.letterDecodings[i])
	}
}

// #
// Prints acronyms in console in poor format (acronym only, without any decoding info).
// 'amount' == 0 means printing all acronyms.
// #
func printAcronyms(acrs Acronyms, amount int) error {

	switch {
	case amount < 0:
		return errors.New("incorrect (negative) amount of acronyms")
	case amount == 0:
		amount = len(acrs)
	case amount > len(acrs):
		return errors.New("too many acronyms are requested to print")
	}

	fmt.Printf("\nList of acronyms:\n")
	for i := 0; i < amount; i++ {
		fmt.Printf("%s%s%d", acrs[i].word, TokenSeparator, acrs[i].sumEstimation)
	}
	fmt.Printf("\n")

	return nil
}

// #
// Prints most suitable (by sumEstimation) acronyms in console in poor format (acronym only, without any decoding info).
// #
func printMostSuitableAcronyms(acrs Acronyms, amount int) error {

	// TODO optimize by not sorting all elements, but by taking the amount of best ones
	sortedAcrs := make(Acronyms, len(acrs))
	copy(sortedAcrs, acrs)
	SortAcronymsBySumEstimation(sortedAcrs)

	return printAcronyms(sortedAcrs, amount)
}
