package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"acrgen/tfp"
)

func main() {
	argsWithoutProgName := os.Args[1:]
	if len(argsWithoutProgName) != 3 {
		fmt.Println("Неверное количество входных аргументов!")
		fmt.Println("Запустите программу заново, указав названия трёх  \".txt\" файлов: входного, с существующими словами-кандидатами и выходного")
		return
	}
	srcFilename := argsWithoutProgName[0]

	_, err := importSrcFromFile(srcFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// #
// Describes one source file entry (line), that represents a variant of acronym letter, its estimation and decoding (description).
// #
type LetterOpt struct {
	letter     rune
	estimation int
	decoding   string
}
type LetterOpts []LetterOpt
type Src []LetterOpts

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
		letter := letterToken[0]

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
