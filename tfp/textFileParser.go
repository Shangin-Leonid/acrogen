package tfp /*Text File Parser */

import (
	"bufio"
	"errors"
	"os"
	"path"
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
	if !isTextFileNameValid(filename) {
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

// #
// Checks validness of name of plain text file.
// #
func isTextFileNameValid(filename string) bool {
	ext := path.Ext(filename)
	if ext != ".txt" {
		return false
	}

	return true
}
