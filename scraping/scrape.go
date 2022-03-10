package scraping

import (
	"happy-x-day/utils"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var URL_TO_SCRAPE = "https://www.timeanddate.com/on-this-day/"

func GetTodaysEvents() []string {
	res, err := http.Get(URL_TO_SCRAPE)
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
