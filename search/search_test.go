package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextToSearch(t *testing.T) {
	assertWithTest := assert.New(t)

	testCases := []struct {
		InputText      string
		Subtext        string
		ExpectedResult string
		Description    string
	}{
		{
			InputText:      "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
			Subtext:        "Peter",
			ExpectedResult: "Peter \"1, 20, 75\"",
			Description:    "Case insensitive matching is correct",
		},
		{
			InputText:      "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
			Subtext:        "peter",
			ExpectedResult: "peter \"1, 20, 75\"",
			Description:    "Case insensitive matching is correct",
		},
		{
			InputText:      "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
			Subtext:        "pick",
			ExpectedResult: "pick \"30, 58\"",
			Description:    "extra case, match the word pick",
		},
		{
			InputText:      "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
			Subtext:        "pi",
			ExpectedResult: "pi \"30, 37, 43, 51, 58\"",
			Description:    "match subsections of a word",
		},
		{
			InputText:      "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
			Subtext:        "z",
			ExpectedResult: "z \"<No Output>\"",
			Description:    "no matches for a letter are found",
		},
		{
			InputText:      "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
			Subtext:        "Peterz",
			ExpectedResult: "Peterz \"<No Output>\"",
			Description:    "no matches for a letter are found",
		},
		{
			InputText:      "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
			Subtext:        "PEteRz",
			ExpectedResult: "PEteRz \"<No Output>\"",
			Description:    "mixed case lettering,not found",
		},
		{
			InputText:      "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
			Subtext:        "PEteR",
			ExpectedResult: "PEteR \"1, 20, 75\"",
			Description:    "mixed case lettering found",
		},
	}
	for _, test := range testCases {
		result := TextSearch(test.InputText, test.Subtext)
		assertWithTest.Equal(test.ExpectedResult, result, test.Description)
	}
}
