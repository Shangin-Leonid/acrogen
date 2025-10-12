package fio /* File Input Output */

import (
	"path"
)

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
