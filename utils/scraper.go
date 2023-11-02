package utils

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeData(res *http.Response) int {
	page, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	page.Find("div.quoteDetails").Each(func(index int, element *goquery.Selection) {

		text := element.Text()

		fmt.Println(text)
		// Do something with the textContent collected from this quoteElement
	})

	return 0
}
