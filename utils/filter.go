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
	"genocide",
	"fight",
	"gross",
	"guilty",
	"harmful",
	"hostile",
	"ignorant",
	"immigrant",
	"malicious",
	"threatening",
	"threat",
	"global warming",
	"accident",
	"kill",
	"massacre",
}

func FilterNegativeEvents(events []string) []string {
	var filtered []string

	for _, event := range events {
		isAllGood := true
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
