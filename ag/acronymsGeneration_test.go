package ag_test

import (
	"testing"

	"github.com/Shangin-Leonid/acrogen/ag"
	. "github.com/Shangin-Leonid/acrogen/utils"

	"github.com/stretchr/testify/assert"
)

type LO = ag.LetterOpt
type LOS = ag.LetterOpts

var testEngDict = ag.Dict{
	"cat": Void{},
	"pet": Void{},
	"net": Void{},
	"tab": Void{},

	"you": Void{},
	"she": Void{},
	"him": Void{},

	"me": Void{},
	"I":  Void{},

	"golang":   Void{},
	"linux":    Void{},
	"postgres": Void{},
	"kafka":    Void{},
}

var testRusDict = ag.Dict{
	"КОТ": Void{},
	"ТОК": Void{},
	"ТОМ": Void{},
	"КОМ": Void{},
	"КОК": Void{},

	"ЩУП": Void{},
	"НИЦ": Void{},
	"ЗЛО": Void{},

	"ТЫ": Void{},
	"Я":  Void{},

	"ЛЮБОВЬ":      Void{},
	"ДОБРО":       Void{},
	"ДОБРОДЕТЕЛЬ": Void{},
	"БЛАГОДАТЬ":   Void{},
}

var (
	rus_1_Ы_LetterOpts = LOS{
		LO{'Ы', 1, "Ы один"},
	}
	rus_2_Я_LetterOpts = LOS{
		LO{'Я', 2, "Я два"},
	}
	rus_3_ОТ_LetterOpts = LOS{
		LO{'О', 3, "О три"},
		LO{'Т', 3, "Т три"},
	}
	rus_4_ТЧКО_LetterOpts = LOS{
		LO{'Т', 4, "Т четыре"},
		LO{'Ч', 4, "Ч четыре"},
		LO{'К', 4, "К четыре"},
		LO{'О', 4, "О четыре"},
	}
	rus_5_МОЩК_LetterOpts = LOS{
		LO{'М', 5, "М пять"},
		LO{'О', 5, "О пять"},
		LO{'Щ', 5, "Щ пять"},
		LO{'К', 5, "К пять"},
	}
	rus_6_ОР_LetterOpts = LOS{
		LO{'О', 6, "О шесть"},
		LO{'Р', 6, "Р шесть"},
	}
	rus_7_ЭЮЯ_LetterOpts = LOS{
		LO{'Э', 7, "Э семь"},
		LO{'Ю', 7, "Ю семь"},
		LO{'Я', 7, "Я семь"},
	}
)

func Test_GenerateAcronyms_Ordered_Rus(t *testing.T) {

	testSuites := []struct {
		name         string
		src          ag.Src
		acrsExpected ag.Acronyms
	}{
		{"empty src",
			ag.Src{},
			ag.Acronyms{},
		},

		{"one non existing letter",
			ag.Src{rus_1_Ы_LetterOpts},
			ag.Acronyms{},
		},

		{"one existing letter",
			ag.Src{rus_2_Я_LetterOpts},
			ag.Acronyms{
				{"Я", 2, []string{"Я два"}},
			},
		},

		{"one group one-letter word",
			ag.Src{rus_7_ЭЮЯ_LetterOpts},
			ag.Acronyms{
				{"Я", 7, []string{"Я семь"}},
			},
		},

		{"three groups no acr",
			ag.Src{rus_3_ОТ_LetterOpts, rus_5_МОЩК_LetterOpts, rus_7_ЭЮЯ_LetterOpts},
			ag.Acronyms{},
		},

		{"three groups two acrs",
			ag.Src{rus_3_ОТ_LetterOpts, rus_4_ТЧКО_LetterOpts, rus_5_МОЩК_LetterOpts},
			ag.Acronyms{
				{"ТОК", 12, []string{"Т три", "О четыре", "К пять"}},
				{"ТОМ", 12, []string{"Т три", "О четыре", "М пять"}},
			},
		},

		{"three groups four acrs",
			ag.Src{rus_4_ТЧКО_LetterOpts, rus_3_ОТ_LetterOpts, rus_5_МОЩК_LetterOpts},
			ag.Acronyms{
				{"ТОК", 12, []string{"Т четыре", "О три", "К пять"}},
				{"ТОМ", 12, []string{"Т четыре", "О три", "М пять"}},
				{"КОМ", 12, []string{"К четыре", "О три", "М пять"}},
				{"КОК", 12, []string{"К четыре", "О три", "К пять"}},
			},
		},

		{"four groups no acrs",
			ag.Src{rus_3_ОТ_LetterOpts, rus_7_ЭЮЯ_LetterOpts, rus_4_ТЧКО_LetterOpts, rus_5_МОЩК_LetterOpts},
			ag.Acronyms{},
		},
	}

	for _, ts := range testSuites {
		t.Run(ts.name, func(t *testing.T) {

			// Test code
			acrs := ag.GenerateAcronyms(ts.src, testRusDict, ag.Ordered)
			assert.Equal(t, len(ts.acrsExpected), len(acrs))
			assert.ElementsMatch(t, ts.acrsExpected, acrs)

		})
	}
}

func Test_GenerateAcronyms_NonOrdered_Rus(t *testing.T) {

	testSuites := []struct {
		name         string
		src          ag.Src
		acrsExpected ag.Acronyms
	}{
		{"empty src",
			ag.Src{},
			ag.Acronyms{},
		},

		{"one non existing letter",
			ag.Src{rus_1_Ы_LetterOpts},
			ag.Acronyms{},
		},

		{"one existing letter",
			ag.Src{rus_2_Я_LetterOpts},
			ag.Acronyms{
				{"Я", 2, []string{"Я два"}},
			},
		},

		{"one group one-letter word",
			ag.Src{rus_7_ЭЮЯ_LetterOpts},
			ag.Acronyms{
				{"Я", 7, []string{"Я семь"}},
			},
		},

		{"three groups no acr",
			ag.Src{rus_3_ОТ_LetterOpts, rus_5_МОЩК_LetterOpts, rus_7_ЭЮЯ_LetterOpts},
			ag.Acronyms{},
		},

		{"three groups first order",
			ag.Src{rus_3_ОТ_LetterOpts, rus_4_ТЧКО_LetterOpts, rus_5_МОЩК_LetterOpts},
			ag.Acronyms{
				{"КОТ", 12, []string{"К четыре", "О пять", "Т три"}},
				{"КОТ", 12, []string{"К пять", "О три", "Т четыре"}},
				{"КОТ", 12, []string{"К пять", "О четыре", "Т три"}},

				{"ТОК", 12, []string{"Т три", "О четыре", "К пять"}},
				{"ТОК", 12, []string{"Т три", "О пять", "К четыре"}},
				{"ТОК", 12, []string{"Т четыре", "О три", "К пять"}},

				{"ТОМ", 12, []string{"Т три", "О четыре", "М пять"}},
				{"ТОМ", 12, []string{"Т четыре", "О три", "М пять"}},

				{"КОМ", 12, []string{"К четыре", "О три", "М пять"}},

				{"КОК", 12, []string{"К четыре", "О три", "К пять"}},
				{"КОК", 12, []string{"К пять", "О три", "К четыре"}},
			},
		},

		{"three groups second order",
			ag.Src{rus_4_ТЧКО_LetterOpts, rus_3_ОТ_LetterOpts, rus_5_МОЩК_LetterOpts},
			ag.Acronyms{
				{"КОТ", 12, []string{"К четыре", "О пять", "Т три"}},
				{"КОТ", 12, []string{"К пять", "О три", "Т четыре"}},
				{"КОТ", 12, []string{"К пять", "О четыре", "Т три"}},

				{"ТОК", 12, []string{"Т три", "О четыре", "К пять"}},
				{"ТОК", 12, []string{"Т три", "О пять", "К четыре"}},
				{"ТОК", 12, []string{"Т четыре", "О три", "К пять"}},

				{"ТОМ", 12, []string{"Т три", "О четыре", "М пять"}},
				{"ТОМ", 12, []string{"Т четыре", "О три", "М пять"}},

				{"КОМ", 12, []string{"К четыре", "О три", "М пять"}},

				{"КОК", 12, []string{"К четыре", "О три", "К пять"}},
				{"КОК", 12, []string{"К пять", "О три", "К четыре"}},
			},
		},

		{"four groups no acrs",
			ag.Src{rus_3_ОТ_LetterOpts, rus_7_ЭЮЯ_LetterOpts, rus_4_ТЧКО_LetterOpts, rus_5_МОЩК_LetterOpts},
			ag.Acronyms{},
		},
	}

	for _, ts := range testSuites {
		t.Run(ts.name, func(t *testing.T) {

			// Test code
			acrs := ag.GenerateAcronyms(ts.src, testRusDict, ag.NonOrdered)
			assert.Equal(t, len(ts.acrsExpected), len(acrs))
			assert.ElementsMatch(t, ts.acrsExpected, acrs)

		})
	}
}
