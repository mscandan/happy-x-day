package scraping

import (
	"happy-x-day/utils"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	BASIC_SCRAPE_URL    = "https://www.timeanddate.com/on-this-day/"
	ADVANCED_SCRAPE_URL = "https://www.onthisday.com/today/events.php"
)

func GetBasicScrapingEvents() []string {
	res, err := http.Get(BASIC_SCRAPE_URL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	allEvents := []string{}

	doc.Find(".otd-detail").Each(func(_ int, eventList *goquery.Selection) {
		eventList.Find(".otd-title").Each(func(_ int, event *goquery.Selection) {
			txt := event.Text()
			allEvents = append(allEvents, txt)
		})
	})

	filteredEvents := utils.FilterNegativeEvents(allEvents)

	return filteredEvents
}

func GetAdvancedScrapingEvents() []string {
	res, err := http.Get(ADVANCED_SCRAPE_URL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	allEvents := []string{}

	doc.Find(".event").Each(func(_ int, event *goquery.Selection) {
		txt := event.Text()
		allEvents = append(allEvents, txt)
	})

	filteredEvents := utils.FilterNegativeEvents(allEvents)

	return filteredEvents
}
