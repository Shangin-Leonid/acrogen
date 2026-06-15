package utils /* Utils */

import (
	"path"
	"path/filepath"
	"unicode/utf8"
)

// # IsTextFileNameValid checks validness of name of plain text file.
//
// # Params:
//
//   - filename
//
// # Returns:
//
//   - flag if valid
//
// # TODOs:
//
//   - Check file existance. If it exists then ask user if he is sure about rewriting content.
func IsTextFileNameValid(filename string) bool {

	if !utf8.ValidString(filename) {
		return false
	}

	ext := path.Ext(filename)
	if ext != ".txt" {
		return false
	}

	if filename == ext {
		return false
	}

	return true
}

// # WithoutExt trims file extension from filename.
func WithoutExt(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}
