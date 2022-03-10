package main

import (
	"fmt"
	"happy-x-day/scraping"
)

func main() {
	fmt.Println("hello")
	events := scraping.GetTodaysEvents()

	for _, event := range events {
		fmt.Println(event)
	}
}
