package ag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convertToAcronym(t *testing.T) {
	type LO = LetterOpt
	testSuites := []struct {
		name        string
		losParam    LetterOpts
		acrExpected Acronym
	}{
		{"first test suite",
			LetterOpts{
				LO{'D', 1, "Don't"},
				LO{'R', 5, "Repeat"},
				LO{'Y', 3, "Yourself"},
			},
			Acronym{
				Word:            "DRY",
				SumEstimation:   9,
				LetterDecodings: []string{"Don't", "Repeat", "Yourself"},
			},
		},
	}

	for _, ts := range testSuites {
		t.Run(ts.name, func(t *testing.T) {

			// Test code
			res := convertToAcronym(ts.losParam)
			assert.Equal(t, res, ts.acrExpected)

		})
	}
}

func Test_asWord(t *testing.T) {
	type LO = LetterOpt
	testSuites := []struct {
		name           string
		losParam       LetterOpts
		asWordExpected string
	}{
		{"first test suite",
			LetterOpts{
				LO{'D', 1, "Don't"},
				LO{'R', 5, "Repeat"},
				LO{'Y', 3, "Yourself"},
			},
			"DRY",
		},

		{"second test suite",
			LetterOpts{
				LO{'D', 1, "Don't"},
				LO{'R', 5, "Repeat"},
				LO{'Y', 3, "Yourself"},
				LO{'P', 6, "Please"},
			},
			"DRYP",
		},
	}

	for _, ts := range testSuites {
		t.Run(ts.name, func(t *testing.T) {

			// Test code
			res := asWord(ts.losParam)
			assert.Equal(t, res, ts.asWordExpected)

		})
	}
}
