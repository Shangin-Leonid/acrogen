package fio /* File Input Output */

import (
	"bufio"
	"errors"
	"os"
	"strconv"
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
	if firstLineParserFunc != nil && fs.Scan() {
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

// #
// Creates new file with 'filename'.
// Writes slice elements in file in format of 'formatFunc'.
// Returns amount of successful written elements and error.
// #
func WriteSliceToTextFile[T any](slice []T, filename string, needWriteLen bool, formatFunc func(T) string) (nSuccessfulWrites int, err error) {
	if !IsTextFileNameValid(filename) {
		return 0, errors.New("incorrect name of output file")
	}

	outputFile, err := os.Create(filename)
	if err != nil {
		return 0, err
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(strconv.Itoa(len(slice)) + "\n\n")
	if err != nil {
		return 0, err
	}
	nSuccessfulWrites++
	for i := range slice {
		_, err = outputFile.WriteString(formatFunc(slice[i]))
		if err != nil {
			return nSuccessfulWrites, err
		}
		nSuccessfulWrites++
	}

	return nSuccessfulWrites, nil
}
