package main

import (
	"log"
	"scraper/structs"
	"scraper/utils"
)

func main() {

	url := utils.GetInput("Enter a valid goodreads quote url: ")

	fileName := utils.GetInput("Enter a filename: ")

	scrapeMode := utils.GetInput("Enter scrape mode (multi/single): ")

	var quotes []structs.Quote

	if scrapeMode == "single" {
		utils.SingleScrape(fileName, url, quotes)
	} else if scrapeMode == "multi" {
		utils.MultiScrape(fileName, url, quotes)
	} else {
		log.Fatal("Please enter a valid scrape mode.")
	}

}
