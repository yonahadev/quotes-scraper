package utils

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ProcessUrl(url string) *goquery.Document {
	webpage, err := http.Get(url)
	if err != nil {
		log.Fatal("Please enter a valid url.")
	}
	page, err := goquery.NewDocumentFromReader(webpage.Body)
	if err != nil {
		log.Fatal(err)
	}
	return page
}
