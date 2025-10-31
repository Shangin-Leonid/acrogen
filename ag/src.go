package ag /* Acronyms Generation */

// #
// Describes one source file entry (line), that represents a variant of acronym letter, its estimation and decoding (description).
// #
type LetterOpt struct {
	Letter     rune
	Estimation int
	Decoding   string
}
type LetterOpts = []LetterOpt
type Src = []LetterOpts
