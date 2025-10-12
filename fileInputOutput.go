package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"

	"acrgen/tfp"
)

// #
// Parses source data file and import its content.
// #
func importSrcFromFile(srcFilename string) (Src, error) {
	src := make(Src, 0, 10)
	src = append(src, make(LetterOpts, 0, 10))

	var parseSrcFileLine tfp.LineParserFunc = func(line string) error {
		const LetterOptsSeparator = ""
		const LetterOptSeparator = " -- "

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

	_, err := tfp.ParseFileLineByLine(srcFilename, parseSrcFileLine)

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

	var parseWordFromFileLine tfp.LineParserFunc = func(line string) error {
		dict[line] = struct{}{}
		return nil
	}

	_, err := tfp.ParseFileLineByLine(dictFilename, parseWordFromFileLine)

	if err != nil {
		return nil, err
	}

	return dict, nil
}
