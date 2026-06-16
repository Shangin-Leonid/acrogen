package fio /* File Input Output */

import (
	"bufio"
	"os"
	"strconv"

	"github.com/Shangin-Leonid/acrogen/utils"
)

type StringParserFunc func(string) error

// # parseTextFileLineByLine opens file, goes through it and parses each line by 'parserFunc'.
//
// # Params:
//
//   - filename
//   - function for parsing first line (maybe nil)
//   - function for parsing other lines
//
// # Returns:
//
//   - amount of successfully parsed lines
//   - error
func parseTextFileLineByLine(filename string, firstLineParserFunc, parserFunc StringParserFunc) (nSuccessfullyParsed int, err error) {
	if !utils.IsTextFileNameValid(filename) {
		return 0, utils.NewSTError("invalid name of text file")
	}

	file, err := os.Open(filename)
	if err != nil {
		return 0, utils.NewSTError(err.Error())
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
		return nSuccessfullyParsed, utils.NewSTError(err.Error())
	}

	return nSuccessfullyParsed, nil
}

// # writeSliceToTextFile creates new file and write slice elements to it using formatting func.
//
// # Params:
//
//   - slice
//   - destination file name
//   - flag if need to write len of slice at first line
//   - func for formatting elements
//
// # Returns:
//
//   - amount of successfully written lines
//   - error
func writeSliceToTextFile[T any](slice []T, filename string, needWriteLen bool, formatFunc func(T) string) (nSuccessfulWrites int, err error) {
	if !utils.IsTextFileNameValid(filename) {
		return 0, utils.NewSTError("incorrect name of output file")
	}

	outputFile, err := os.Create(filename)
	if err != nil {
		return 0, utils.NewSTError(err.Error())
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(strconv.Itoa(len(slice)) + "\n\n")
	if err != nil {
		return 0, utils.NewSTError(err.Error())
	}
	nSuccessfulWrites++
	for i := range slice {
		_, err = outputFile.WriteString(formatFunc(slice[i]))
		if err != nil {
			return nSuccessfulWrites, utils.NewSTError(err.Error())
		}
		nSuccessfulWrites++
	}

	return nSuccessfulWrites, nil
}
