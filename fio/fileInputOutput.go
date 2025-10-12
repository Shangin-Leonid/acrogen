package fio /* File Input Output */

import (
	"errors"
	"os"
	"path"
)

// #
// Checks validness of name of plain text file.
// TODO check file existance. If it exists then ask user if he is sure about rewriting content
// #
func isTextFileNameValid(filename string) bool {
	ext := path.Ext(filename)
	if ext != ".txt" {
		return false
	}

	return true
}

// #
// Creates new file with 'filename'.
// Writes slice elements in file in format of 'formatFunc'.
// Returns amount of successful written elements and error.
// #
func WriteSliceToFile[T any](slice []T, filename string, formatFunc func(T) string) (nSuccessfulWrites int, err error) {
	if !isTextFileNameValid(filename) {
		return 0, errors.New("incorrect name of output file")
	}

	outputFile, err := os.Create(filename)
	if err != nil {
		return 0, err
	}
	defer outputFile.Close()

	for i := range slice {
		_, err = outputFile.WriteString(formatFunc(slice[i]))
		if err != nil {
			return nSuccessfulWrites, err
		}
		nSuccessfulWrites++
	}

	return nSuccessfulWrites, nil
}
