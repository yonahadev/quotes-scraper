package utils

import (
	"fmt"
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

	page.Find("div.quoteDetails").Each(func(index int, element *goquery.Selection) {

		quote := ParseQuote(element.Text())
		author := element.Find("span.authorOrTitle").Text()
		source := element.Find("span").Find("a.authorOrTitle").Text()
		tags := element.Find("div.quoteFooter").Find("div.greyText").Text()
		likes := element.Find("div.quoteFooter").Find("div.right").Text()
		if source == "" {
			source = "source is unknown"
		}

		fmt.Println(quote, author, source, tags, likes)
		// Do something with the textContent collected from this quoteElement
	})

	return text
}
