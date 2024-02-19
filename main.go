package main

import (
	"fmt"

	"github.com/GabDewraj/textSearch/search"
)

func main() {
	// Test examples
	testCases := []struct {
		InputText string
		Subtext   string
	}{
		{
			InputText: "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
			Subtext:   "Peter",
		},
		{
			InputText: "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
			Subtext:   "peter",
		},
		{
			InputText: "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
			Subtext:   "pick",
		},
		{
			InputText: "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
			Subtext:   "pi",
		},
		{
			InputText: "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
			Subtext:   "z",
		},
		{
			InputText: "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
			Subtext:   "Peterz",
		},
	}
	for _, test := range testCases {
		result := search.TextSearch(test.InputText, test.Subtext)
		fmt.Println(result)
	}

}
