package main

import (
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

	utils.ScrapeData(webpage)

}
