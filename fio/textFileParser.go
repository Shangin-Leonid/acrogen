package fio /* File Input Output */

import (
	"bufio"
	"errors"
	"os"
)

type StringParserFunc = func(string) error // TODO without '='

// #
// Opens file, goes through it and parses each line by 'parserFunc'.
// Returns error and amount of lines that have been parsed successfully.
// #
func ParseTextFileLineByLine(filename string, firstLineParserFunc, parserFunc StringParserFunc) (nSuccessfullyParsed int, err error) {
	if !IsTextFileNameValid(filename) {
		return 0, errors.New("некорректное название файла")
	}

	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	fs := bufio.NewScanner(file)
	if firstStringParserFunc != nil && fs.Scan() {
		err = firstLineParserFunc(fs.Text())
		nSuccessfullyParsed++
	}
	for fsSuccess := fs.Scan(); (err == nil) && fsSuccess; fsSuccess = fs.Scan() {
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
