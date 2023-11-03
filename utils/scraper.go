package utils

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeData(res *http.Response) string {
	page, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	text := page.Find("div.quoteDetails").First().Text()

	// page.Find("div.quoteDetails").Each(func(index int, element *goquery.Selection) {

	// 	text := element.Text()

	// 	fmt.Println(text)
	// 	// Do something with the textContent collected from this quoteElement
	// })

	return text
}
