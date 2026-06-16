package ag /* Acronyms Generation */

// # convertToAcronym converts letter options to acronym.
//
// # Params:
//
//   - a collection of letter options
//
// # Returns:
//
//   - acronym based on the input
//
// # Description:
//
// Example:
// {{'D', 1, "Don't"}, {'R', 2, "Repeat"}, {'Y', 3, "Yourself"}} ->
// -> {"DRY", 6, {"Don't", "Repeat", "Yourself"}}
func convertToAcronym(los LetterOpts) Acronym {
	sumEstimation := 0
	letterDecodings := []string{}

	for i := range los {
		sumEstimation += los[i].Estimation
		letterDecodings = append(letterDecodings, los[i].Decoding)
	}

	return Acronym{asWord(los), sumEstimation, letterDecodings}
}

// # asWord builds a word from letter options.
//
// # Params:
//
//   - letter options
//
// # Returns:
//
//   - word build from input
//
// # Description:
//
// Example:
// {{'D', 1, "Don't"}, {'R', 2, "Repeat"}, {'Y', 3, "Yourself"}} -> "DRY"
func asWord(los LetterOpts) string {
	word := make([]rune, 0, len(los))
	for i := range los {
		word = append(word, los[i].Letter)
	}
	return string(word)
}
