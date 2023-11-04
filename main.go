package main

import (
	"fmt"
	"log"
	"scraper/utils"
	"strconv"
	"time"
)

func main() {

	url := utils.GetInput("Enter a valid goodreads quote url:")

	fileName := utils.GetInput("Enter a filename: ")

	scrapeMode := utils.GetInput("Enter scrape mode (multi/single): ")

	if scrapeMode == "single" {
		file := utils.CreateFile(fileName)
		defer file.Close()
		page := utils.GetDocument(url)
		quotes, numberOfQuotes := utils.ScrapeData(page)
		if numberOfQuotes != 0 {
			fmt.Println("Scraping:", numberOfQuotes)
			utils.WriteToFile(file, quotes)
		}
	} else if scrapeMode == "multi" {
		file := utils.CreateFile(fileName)
		defer file.Close()
		baseUrl := utils.ParseUrl(url)
		for i := 1; i < 101; i++ {
			currentUrl := baseUrl + "?page=" + strconv.Itoa(i)
			page := utils.GetDocument(currentUrl)
			quotes, numberOfQuotes := utils.ScrapeData(page)
			if numberOfQuotes == 0 {
				fmt.Println("0 quotes found - terminating search.")
				break
			}
			fmt.Println("Scraping:", currentUrl, numberOfQuotes, "quotes")
			utils.WriteToFile(file, quotes)
			time.Sleep(50 * time.Millisecond)
		}
	} else {
		log.Fatal("Please enter a valid scrape mode.")
	}

}
