package main

import (
	"fmt"
	"log"
	"net/http"
	"scraper/utils"
)

func main() {

	url := utils.GetInput()

	webpage, err := http.Get(url)
	if err != nil {
		log.Fatal("Please enter a valid url.")
	}

	scrapedText := utils.ScrapeData(webpage)

	quote := utils.ParseQuote(scrapedText)

	fmt.Println(quote)
}
