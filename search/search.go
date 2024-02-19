package search

import (
	"fmt"
	"strconv"
	"strings"
)

func TextSearch(textToSearch string, subText string) string {
	// Make the search case insensitive
	lowerCaseTextToSearch := strings.ToLower(textToSearch)
	lowerCaseSubtext := strings.ToLower(subText)
	// Interate through string to find the index of match
	// At the index position of the first character
	var positions []string
	// We want to only iterate up the point where the remaining
	// text is of equal length of the subtext but anything smaller
	// will not create a match
	for i := 0; i <= len(lowerCaseTextToSearch)-len(lowerCaseSubtext); i++ {
		// Reference the string from position i up to length of the subtext
		if lowerCaseTextToSearch[i:i+len(lowerCaseSubtext)] == lowerCaseSubtext {
			// Offset index by one
			positions = append(positions, strconv.Itoa(i+1))
		}
	}
	// Use comma and tab to separate the output positions
	stringedPositions := strings.Join(positions, ", ")
	if len(stringedPositions) == 0 {
		return fmt.Sprintf("%s \"<No Output>\"", subText)
	}
	return fmt.Sprintf("%s \"%s\"", subText, stringedPositions)
}
