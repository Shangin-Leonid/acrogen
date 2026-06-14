package ag /* Acronyms Generation */

// # convertToAcronym
//
// # TODO:
//   - docs
func convertToAcronym(los LetterOpts) Acronym {
	sumEstimation := 0
	letterDecodings := []string{}

	for i := range los {
		sumEstimation += los[i].Estimation
		letterDecodings = append(letterDecodings, los[i].Decoding)
	}

	return Acronym{asWord(los), sumEstimation, letterDecodings}
}

// # asWord
//
// # TODO:
//   - docs
func asWord(los LetterOpts) string {
	word := make([]rune, 0, len(los))
	for i := range los {
		word = append(word, los[i].Letter)
	}
	return string(word)
}
