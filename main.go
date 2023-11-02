package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://www.goodreads.com/quotes"
	text := fmt.Sprintf("Parsing webpage: %s", url)
	fmt.Println(text)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	page, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	page.Find("div.quoteDetails").Each(func(i int, element *goquery.Selection) {
		text := element.Find("div.quoteText").Find("span.authorOrTitle").Text()
		fmt.Println(text)
	})
}
