package fio /* File Input Output */

import (
	"bufio"
	"errors"
	"os"
)

type LineParserFunc = func(string) error // TODO without '='

// #
// Opens file, goes through it and parses each line by 'parserFunc'.
// Returns:
//   - amount of lines that have been parsed successfully
//   - error
//
// #
func ParseFileLineByLine(filename string, parserFunc LineParserFunc) (nSuccessfullyParsed int, err error) {
	if !IsTextFileNameValid(filename) {
		return 0, errors.New("некорректное название файла")
	}

	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	fs := bufio.NewScanner(file)
	for fsSuccess := fs.Scan(); fsSuccess && (err == nil); fsSuccess = fs.Scan() {
		err = parserFunc(fs.Text())
		nSuccessfullyParsed++
	}
	if err != nil {
		return nSuccessfullyParsed - 1, err
	}
	if fsErr := fs.Err(); fsErr != nil {
		return nSuccessfullyParsed, fsErr
	}

	return nSuccessfullyParsed, nil
}
