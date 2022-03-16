package utils

import (
	"regexp"
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
	"nuclear",
	"fire",
	"military",
	"violation",
	"raid",
	"dead",
	"terror",
	"beat",
	"battle",
	"poison",
}

func FilterNegativeEvents(events []string) []string {
	var filtered []string

	for _, event := range events {
		isAllGood := true
		for _, negativeWord := range NegativeSentences {
			searchRegex := "(?i)" + negativeWord
			r := regexp.MustCompile(searchRegex)
			if r.MatchString(event) {
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
