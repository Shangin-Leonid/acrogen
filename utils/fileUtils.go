package utils /* Utils */

import (
	"path"
	"path/filepath"
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
	ext := path.Ext(filename)
	if ext != ".txt" {
		return false
	}

	return true
}

// # GetWithoutExt trims file extension from filename.
func GetWithoutExt(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}
