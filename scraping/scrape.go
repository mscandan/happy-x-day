package scraping

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var URL_TO_SCRAPE = "https://www.timeanddate.com/on-this-day/"

func GetTodaysEvents() {
	res, err := http.Get(URL_TO_SCRAPE)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	doc.Find(".otd-detail").Each(func(_ int, eventList *goquery.Selection) {
		eventList.Find(".otd-title").Each(func(idx int, event *goquery.Selection) {
			txt := event.Text()
			// filter the text with the unwanted words
			fmt.Printf("Event %d: %s\n\n", idx+1, txt)
		})
	})
}
