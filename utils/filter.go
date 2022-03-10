package utils

import (
	"strings"
)

var NegativeSentences = []string{
	"war",
	"death",
	"attack",
	"bomb",
	"coup",
}

func FilterNegativeEvents(events []string) []string {
	var filtered []string

	isAllGood := true
	for _, event := range events {
		for _, negativeWord := range NegativeSentences {
			if strings.Contains(event, negativeWord) {
				isAllGood = false
			}
		}
		if isAllGood {
			filtered = append(filtered, event)
		}
		isAllGood = true
	}

	return filtered
}
