package main

import (
	"fmt"
	"happy-x-day/scraping"
)

func main() {
	fmt.Println("hello")
	basicEvents := scraping.GetBasicScrapingEvents()
	advancedEvents := scraping.GetAdvancedScrapingEvents()

	fmt.Printf("Basic Scraping Found %d Events \n", len(basicEvents))
	for _, event := range basicEvents {
		fmt.Println(event)
	}

	fmt.Printf("Advanced Scraping Found %d Events \n", len(advancedEvents))
	for _, event := range advancedEvents {
		fmt.Println(event)
	}
}
