package main

import (
	"fmt"
	"happy-x-day/scraping"
)

func main() {
	fmt.Println("hello")
	events := scraping.GetBasicScrapingEvents()

	if len(events) < 3 {
		events = append(events, scraping.GetAdvancedScrapingEvents()...)
	}

	for _, event := range events {
		fmt.Println(event)
	}
}
