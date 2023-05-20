package strings

import (
	"fmt"
	"testing"
)

var samples = []string{
	"oneword",
	"oneword123",
	"twowords",
	"twowords456",
	"twowordsabc",
	"twoabcwords",
	"ONEWORD",
	"ONEWORD123",
	"TWOWORDS",
	"TWOWORDS456",
	"TWOWORDSABC",
	"TWOABCWORDS",
	"twoWords",
	"twoWords456",
	"twoWordsABC",
	"twoABCWords",
	"Oneword",
	"Oneword123",
	"TwoWords",
	"TwoWords456",
	"TwoWordsABC",
	"TwoABCWords",
	"two_words",
	"two_words456",
	"two_words_abc",
	"two_abc_words",
	"TWO_WORDS",
	"TWO_WORDS456",
	"TWO_WORDS_ABC",
	"TWO_ABC_WORDS",
	"Two_Words",
	"Two_Words456",
	"Two_Words_ABC",
	"Two_ABC_Words",
	"two-words",
	"two-words456",
	"two-words-abc",
	"two-abc-words",
	"TWO-WORDS",
	"TWO-WORDS456",
	"TWO-WORDS-ABC",
	"TWO-ABC-WORDS",
	"Two-Words",
	"Two-Words456",
	"Two-Words-ABC",
	"Two-ABC-Words",
}

func TestCaseOf(t *testing.T) {
	result := make(map[Case][]string)
	for _, sample := range samples {
		result[CaseOf(sample)] = append(result[CaseOf(sample)], sample)
	}
	for _, c := range allCases {
		fmt.Println(c)
		fmt.Println(result[c])
		fmt.Println()
	}
	fmt.Println(Unknown)
	fmt.Println(result[Unknown])
}

func TestBreak(t *testing.T) {
	for _, sample := range samples {
		fmt.Println(sample)
		fmt.Println(Break(sample))
		fmt.Println()
	}
}

func TestConvertCase(t *testing.T) {
	s := "HTML_CSS_JS"
	for _, c := range allCases {
		fmt.Println(c)
		fmt.Println(ConvertCase(s, c))
		fmt.Println()
	}
}

func TestCamelWords(t *testing.T) {
	words := []string{"", "", ""}
	camelWords(words, true, false)
	fmt.Println(words)
}
