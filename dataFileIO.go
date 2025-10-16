package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"acrgen/fio"
)

const TokenSeparator = " -- "
const LineSeparator = ""

// #
// Parses source data file and import its content.
// #
func importSrcFromFile(srcFilename string) (Src, error) {
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
			outp := acr.word + TokenSeparator + strconv.Itoa(acr.sumEstimation) + "\n"
			// TODO optimize by switching from string to []rune
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
// Import acronyms from dump file with 'FullFormat'.
// #
func importAcronymsFromFile(dumpFilename string) (acrs Acronyms, err error) {

	var parseFirstLineInDumpFile fio.StringParserFunc = func(line string) error {
		if len(line) == 0 {
			return errors.New("unexpected empty first line in dump file")
		}
		size, err := strconv.Atoi(line[:len(line)-1]) // Exclude '\n' ending
		if err == nil {
			acrs = make(Acronyms, 0, size)
		}
		return err
	}

	// Enumeration represents the type of previous parsed line.
	type PrevLineT int
	const (
		First PrevLineT = iota + 1
		Empty
		Acr
		NotLastLetter
		LastLetterInCurAcr
	)
	prev := First
	// TODO optimize strings and runes. Maybe use new type to represent strings
	// TODO last line must be empty
	var parseAcronymsInDumpFile fio.StringParserFunc = func(line string) error {
		var cur PrevLineT
		lineAsRune := []rune(line)

		if line == LineSeparator+"\n" || line == "\n" || line == "" {
			switch prev {
			case Empty, NotLastLetter, Acr:
				return errors.New("data/format error: unexpected empty line")
			}

			lastLD := acrs[len(acrs)-1].letterDecodings
			if prev == NotLastLetter && len(lastLD) != cap(lastLD) {
				return errors.New("data/format error: unexpected empty line within acronyms description")
			}
			cur = Empty
		} else if len(lineAsRune) < 2 {
			return errors.New("data/format error: some short line with no unexpected meaning")
		} else if unicode.IsLetter(lineAsRune[0]) && !unicode.IsLetter(lineAsRune[1]) {
			// TODO check or delete last '\n' if exist ?
			// TODO check upper or lower case letter
			if prev != NotLastLetter && prev != Acr {
				return errors.New("data/format error: maybe unexpected letter decoding line")
			}

			curAcrAsRune := []rune(acrs[len(acrs)-1].word)
			lastLD := acrs[len(acrs)-1].letterDecodings
			curLetterInd := len(lastLD)

			if lineAsRune[0] != curAcrAsRune[curLetterInd] {
				return errors.New("incorrect data: the letter is not the same as in the acronym")
			}

			tsInd := strings.Index(line, TokenSeparator)
			if tsInd == -1 {
				return errors.New("data/format error: incorrect format of letter decoding line (no token separator between letter and decoding)")
			}
			decodingInd := tsInd + len(TokenSeparator)

			if len(lastLD) == cap(lastLD) {
				return errors.New("unexpected error: 'len(lD) == cap(lD)'")
			}
			lastLD = append(lastLD, line[decodingInd:])

			if len(lastLD) == cap(lastLD) {
				cur = LastLetterInCurAcr
			} else {
				cur = NotLastLetter
			}
		} else {
			if prev != Empty {
				return errors.New("data/format error: unexpected acronym line or smth else")
			}

			tsInd := strings.Index(line, TokenSeparator)
			if tsInd == -1 {
				return errors.New("data/format error: incorrect format of acronym line (no token separator between word and estimation)")
			}
			estInd := tsInd + len(TokenSeparator)
			est, err := strconv.Atoi(line[estInd:])
			if err != nil {
				return errors.New("data/format error: incorrect summary estimation of acronym")
			}

			for _, l := range line[:tsInd] {
				if !unicode.IsLetter(l) || !!unicode.IsLower(l) {
					return errors.New("data/format error: not letters or upper case letters in acronym")
				}
			}
			acrs = append(acrs, Acronym{line[:tsInd], est, make([]string, 0, tsInd)})

			cur = Acr
		}

		prev = cur
		return nil
	}

	_, err = fio.ParseTextFileLineByLine(dumpFilename, parseFirstLineInDumpFile, parseAcronymsInDumpFile)

	return acrs, err
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
