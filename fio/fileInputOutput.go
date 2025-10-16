package fio /* File Input Output */

import (
	"errors"
	"os"
)

// #
// Creates new file with 'filename'.
// Writes slice elements in file in format of 'formatFunc'.
// Returns amount of successful written elements and error.
// #
func WriteSliceToTextFile[T any](slice []T, filename string, formatFunc func(T) string) (nSuccessfulWrites int, err error) {
	if !IsTextFileNameValid(filename) {
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
