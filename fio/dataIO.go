package fio /* File Input Output */

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/Shangin-Leonid/acrogen/ag"
	"github.com/Shangin-Leonid/acrogen/utils"
)

// # LoadSrcFromFile parses source data file and load its content.
//
// # Params:
//
//   - source file name
//
// # Returns:
//
//   - 'ag.Src' collection
//   - error
func LoadSrcFromFile(srcFilename string) (ag.Src, error) {
	src := make(ag.Src, 0, 10)
	src = append(src, make(ag.LetterOpts, 0, 10))

	var parseSrcFileLine StringParserFunc = func(line string) error {
		if line == LineSeparator {
			if len(src[len(src)-1]) == 0 {
				return utils.NewSTError("incorrect format of input file: first (initial) or multiple consecutive blank lines are prohibited")
			}
			src = append(src, make(ag.LetterOpts, 0, 10))
			return nil
		}

		splittedLine := strings.Split(line, TokenSeparator)
		if len(splittedLine) != 3 {
			return utils.NewSTError("incorrect format of input file: unexpected data format error during reading the file")
		}

		letterToken := []rune(splittedLine[0])
		if len(letterToken) != 1 || !unicode.IsLetter(letterToken[0]) {
			return utils.NewSTError("incorrect format of input file. first token is not a letter")
		}
		letter := unicode.ToLower(letterToken[0])

		estimation, err := strconv.Atoi(splittedLine[1])
		if err != nil {
			return utils.NewSTError("incorrect format of input file. second token is not a number or incorrect number")
		}

		decoding := splittedLine[2]

		src[len(src)-1] = append(src[len(src)-1], ag.LetterOpt{letter, estimation, decoding})
		return nil
	}

	_, err := parseTextFileLineByLine(srcFilename, nil, parseSrcFileLine)

	if err != nil {
		return nil, err
	}

	if len(src[len(src)-1]) == 0 {
		src = src[:len(src)-1]
	}

	return src, nil
}

// # LoadDictionaryFromFile parses dictionary file (list of valid words) and load its content.
//
// # Params:
//
//   - dictionary file name
//   - expected amount of words
//
// # Returns:
//
//   - 'ag.Dict' collection
//   - error
func LoadDictionaryFromFile(dictFilename string, expectedWordsAmount uint64) (ag.Dict, error) {
	dict := make(ag.Dict, expectedWordsAmount)

	var parseWordFromFileLine StringParserFunc = func(line string) error {
		dict[line] = struct{}{}
		return nil
	}

	_, err := parseTextFileLineByLine(dictFilename, nil, parseWordFromFileLine)

	if err != nil {
		return nil, err
	}

	return dict, nil
}

type ExportModeT int

// Enumeration represents mode of acronyms file export.
const (
	_ ExportModeT = iota
	FullFormat
	OnelineFormat
)

// # SaveAcronymsToFile saves (exports) acronyms to output file.
//
// # Params:
//
//   - acronyms to save
//   - destination file name
//   - mode of export (full or oneline, see fio.ExportModeT)
//
// # Returns:
//
//   - error
//
// # Description:
//
// The format is the third parameter passed to function.
func SaveAcronymsToFile(acrs ag.Acronyms, outputFilename string, mode ExportModeT) error {

	var formatFunc func(acr ag.Acronym) string
	if mode == FullFormat {
		formatFunc = func(acr ag.Acronym) string {
			outp := acr.Word + TokenSeparator + strconv.Itoa(acr.SumEstimation) + "\n"
			for i, letter := range []rune(acr.Word) {
				outp += string(letter) + TokenSeparator + acr.LetterDecodings[i] + "\n"
			}
			outp += LineSeparator + "\n"
			return outp
		}
	} else if mode == OnelineFormat {
		formatFunc = func(acr ag.Acronym) string {
			return acr.Word + "\n"
		}
	}

	_, err := writeSliceToTextFile(acrs, outputFilename, true, formatFunc)
	return err
}

// # LoadAcronymsFromFile imports acronyms from file with 'FullFormat'.
//
// # Params:
//
//   - source file name
//
// # Returns:
//
//   - acronyms have been read
//   - error
//
// # TODOs:
//   - Refactor, simplify (see 'parseAcronymsInFile' inside)
func LoadAcronymsFromFile(filename string) (acrs ag.Acronyms, err error) {

	var parseFirstLineInFile StringParserFunc = func(line string) error {
		if len(line) == 0 {
			return utils.NewSTError("unexpected empty first line in file")
		}
		capacity, err := strconv.Atoi(line)
		if err != nil {
			return utils.NewSTError(err.Error())
		}

		acrs = make(ag.Acronyms, 0, capacity)

		return nil
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
	var parseAcronymsInFile StringParserFunc = func(line string) error {
		lineRunes := []rune(line)

		if line == LineSeparator {
			cur = Empty
			if !isCurLineTypeCorrect(cur, prev) {
				return utils.NewSTError("incorrect data/format: unexpected empty line")
			}
			if prev != First {
				lastLD := acrs[len(acrs)-1].LetterDecodings
				if prev == Letter && len(lastLD) != cap(lastLD) {
					return utils.NewSTError("incorrect data/format: unexpected empty line within acronyms description")
				}
			}
		} else if len(lineRunes) < 2 {
			return utils.NewSTError("incorrect data/format: some short line with no unexpected meaning")
		} else if unicode.IsLetter(lineRunes[0]) && !unicode.IsLetter(lineRunes[1]) {
			cur = Letter
			if !isCurLineTypeCorrect(cur, prev) {
				return utils.NewSTError("incorrect data/format: maybe unexpected letter decoding line")
			}

			if !unicode.IsLower(lineRunes[0]) {
				return utils.NewSTError("incorrect format: the letter is not lowercase")
			}

			curAcrAsRune := []rune(acrs[len(acrs)-1].Word)
			curLetterInd := len(acrs[len(acrs)-1].LetterDecodings)
			if lineRunes[0] != curAcrAsRune[curLetterInd] {
				return utils.NewSTError("incorrect data: the letter is not the same as in the acronym")
			}

			tsInd := strings.Index(line, TokenSeparator)
			if tsInd == -1 {
				return utils.NewSTError("incorrect data/format: incorrect format of letter decoding line (no token separator between letter and decoding)")
			}
			decodingInd := tsInd + len(TokenSeparator)
			if len(acrs[len(acrs)-1].LetterDecodings) == cap(acrs[len(acrs)-1].LetterDecodings) {
				return utils.NewSTError("unexpected error: maybe unexpected (extra) letter in acronym")
			}
			acrs[len(acrs)-1].LetterDecodings = append(acrs[len(acrs)-1].LetterDecodings, line[decodingInd:])

			if curLetterInd == len(curAcrAsRune)-1 {
				cur = LastAcrLetter
			}
		} else {
			cur = Acr
			if !isCurLineTypeCorrect(cur, prev) {
				return utils.NewSTError("incorrect data/format: unexpected acronym line or smth else")
			}

			tsInd := strings.Index(line, TokenSeparator)
			if tsInd == -1 {
				return utils.NewSTError("incorrect data/format: incorrect format of acronym line (no token separator between word and estimation)")
			}
			estInd := tsInd + len(TokenSeparator)
			est, err := strconv.Atoi(line[estInd:])
			if err != nil {
				return utils.NewSTError("incorrect data/format: incorrect summary estimation of acronym")
			}

			nLetters := 0
			for _, l := range line[:tsInd] {
				if !unicode.IsLetter(l) || !unicode.IsLower(l) {
					return utils.NewSTError("incorrect data/format: not letters or upper case letters in acronym")
				}
				nLetters++
			}
			acrs = append(acrs, ag.Acronym{line[:tsInd], est, make([]string, 0, nLetters)})
		}

		prev = cur
		return nil
	}

	_, err = parseTextFileLineByLine(filename, parseFirstLineInFile, parseAcronymsInFile)

	return acrs, err
}
