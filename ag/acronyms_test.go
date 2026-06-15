package ag_test

import (
	"testing"

	"github.com/Shangin-Leonid/acrogen/ag"

	"github.com/stretchr/testify/assert"
)

var (
	acr0 = ag.Acronym{
		Word:            "DRY",
		SumEstimation:   7,
		LetterDecodings: []string{"Don't", "Repeat", "Yourself"},
	}
	acr1 = ag.Acronym{
		Word:            "KISS",
		SumEstimation:   6,
		LetterDecodings: []string{"Keep", "It", "Simple", "Stupid"},
	}
	acr2 = ag.Acronym{
		Word:            "PRICE",
		SumEstimation:   9,
		LetterDecodings: []string{"Protect", "Rest", "Ice", "Compression", "Elevation"},
	}
	acr3 = ag.Acronym{
		Word:            "UK",
		SumEstimation:   13,
		LetterDecodings: []string{"United", "Kingdom"},
	}
)

func Test_ContainsAcronym(t *testing.T) {

	testAcrs := ag.Acronyms{acr1, acr0, acr3, acr2}

	testSuites := []struct {
		name        string
		searchWord  string
		acrs        ag.Acronyms
		indExpected int
		okExpected  bool
	}{
		{"empty collection",
			"any word",
			ag.Acronyms{},
			-1,
			false,
		},
		{"empty word",
			"",
			testAcrs,
			-1,
			false,
		},
		{"found 1",
			"UK",
			testAcrs,
			2,
			true,
		},
		{"found 2",
			"DRY",
			testAcrs,
			1,
			true,
		},
		{"not found",
			"U K",
			testAcrs,
			-1,
			false,
		},
	}

	for _, ts := range testSuites {
		t.Run(ts.name, func(t *testing.T) {

			// Test code
			ind, ok := ag.ContainsAcronym(ts.searchWord, ts.acrs)
			assert.Equal(t, ts.indExpected, ind)
			assert.Equal(t, ts.okExpected, ok)

		})
	}
}

func Test_ContainsAcronymBS(t *testing.T) {

	testAcrs := ag.Acronyms{acr0, acr1, acr2, acr3}

	testSuites := []struct {
		name        string
		searchWord  string
		acrs        ag.Acronyms
		indExpected int
		okExpected  bool
	}{
		{"empty collection",
			"any word",
			ag.Acronyms{},
			0,
			false,
		},
		{"empty word",
			"",
			testAcrs,
			0,
			false,
		},
		{"found 1",
			"UK",
			testAcrs,
			3,
			true,
		},
		{"found 2",
			"DRY",
			testAcrs,
			0,
			true,
		},
		{"not found",
			"U K",
			testAcrs,
			3,
			false,
		},
	}

	for _, ts := range testSuites {
		t.Run(ts.name, func(t *testing.T) {

			// Test code
			ind, ok := ag.ContainsAcronymBS(ts.searchWord, ts.acrs)
			assert.Equal(t, ts.indExpected, ind)
			assert.Equal(t, ts.okExpected, ok)

		})
	}
}

func Test_TakeAcronym(t *testing.T) {

	testAcrs := ag.Acronyms{acr1, acr0, acr3, acr2}

	testSuites := []struct {
		name        string
		searchWord  string
		acrs        ag.Acronyms
		acrExpected ag.Acronym
		okExpected  bool
	}{
		{"empty collection",
			"any word",
			ag.Acronyms{},
			ag.Acronym{},
			false,
		},
		{"empty word",
			"",
			testAcrs,
			ag.Acronym{},
			false,
		},
		{"found 1",
			"UK",
			testAcrs,
			acr3,
			true,
		},
		{"found 2",
			"DRY",
			testAcrs,
			acr0,
			true,
		},
		{"not found",
			"U K",
			testAcrs,
			ag.Acronym{},
			false,
		},
	}

	for _, ts := range testSuites {
		t.Run(ts.name, func(t *testing.T) {

			// Test code
			acr, ok := ag.TakeAcronym(ts.searchWord, ts.acrs)
			assert.Equal(t, ts.acrExpected, acr)
			assert.Equal(t, ts.okExpected, ok)

		})
	}
}

func Test_TakeAcronymBS(t *testing.T) {

	testAcrs := ag.Acronyms{acr0, acr1, acr2, acr3}

	testSuites := []struct {
		name        string
		searchWord  string
		acrs        ag.Acronyms
		acrExpected ag.Acronym
		okExpected  bool
	}{
		{"empty collection",
			"any word",
			ag.Acronyms{},
			ag.Acronym{},
			false,
		},
		{"empty word",
			"",
			testAcrs,
			ag.Acronym{},
			false,
		},
		{"found 1",
			"UK",
			testAcrs,
			acr3,
			true,
		},
		{"found 2",
			"DRY",
			testAcrs,
			acr0,
			true,
		},
		{"not found",
			"U K",
			testAcrs,
			ag.Acronym{},
			false,
		},
	}

	for _, ts := range testSuites {
		t.Run(ts.name, func(t *testing.T) {

			// Test code
			acr, ok := ag.TakeAcronymBS(ts.searchWord, ts.acrs)
			assert.Equal(t, ts.acrExpected, acr)
			assert.Equal(t, ts.okExpected, ok)

		})
	}
}

func Test_SortAcronymsBySumEstimation(t *testing.T) {

	testSuites := []struct {
		name         string
		acrs         ag.Acronyms
		acrsExpected ag.Acronyms
	}{
		{"empty collection",
			ag.Acronyms{},
			ag.Acronyms{},
		},
		{"one acronym",
			ag.Acronyms{acr1},
			ag.Acronyms{acr1},
		},
		{"already sorted",
			ag.Acronyms{acr3, acr2, acr0, acr1},
			ag.Acronyms{acr3, acr2, acr0, acr1},
		},
		{"need to be sorted",
			ag.Acronyms{acr1, acr0, acr2, acr3},
			ag.Acronyms{acr3, acr2, acr0, acr1},
		},
	}

	for _, ts := range testSuites {
		t.Run(ts.name, func(t *testing.T) {

			// Test code
			ag.SortAcronymsBySumEstimation(ts.acrs)
			assert.Equal(t, ts.acrsExpected, ts.acrs)

		})
	}
}

func Test_SortAcronymsByAlphabet(t *testing.T) {

	testSuites := []struct {
		name         string
		acrs         ag.Acronyms
		acrsExpected ag.Acronyms
	}{
		{"empty collection",
			ag.Acronyms{},
			ag.Acronyms{},
		},
		{"one acronym",
			ag.Acronyms{acr1},
			ag.Acronyms{acr1},
		},
		{"already sorted",
			ag.Acronyms{acr0, acr1, acr2, acr3},
			ag.Acronyms{acr0, acr1, acr2, acr3},
		},
		{"need to be sorted",
			ag.Acronyms{acr3, acr2, acr0, acr1},
			ag.Acronyms{acr0, acr1, acr2, acr3},
		},
	}

	for _, ts := range testSuites {
		t.Run(ts.name, func(t *testing.T) {

			// Test code
			ag.SortAcronymsByAlphabet(ts.acrs)
			assert.Equal(t, ts.acrsExpected, ts.acrs)

		})
	}
}
