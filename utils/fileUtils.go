package utils /* Utils */

import (
	"path"
	"path/filepath"
)

// #
// Checks validness of name of plain text file.
// TODO check file existance. If it exists then ask user if he is sure about rewriting content
// #
func IsTextFileNameValid(filename string) bool {
	ext := path.Ext(filename)
	if ext != ".txt" {
		return false
	}

	return true
}

// TODO docs
func GetWithoutExt(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}
