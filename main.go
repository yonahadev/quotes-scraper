package main

import (
	"fmt"
	"log"
	"scraper/structs"
	"scraper/utils"
	"strconv"
	"time"
)

func main() {

	url := utils.GetInput("Enter a valid goodreads quote url: ")

	fileName := utils.GetInput("Enter a filename: ")

	scrapeMode := utils.GetInput("Enter scrape mode (multi/single): ")

	var quotes []structs.Quote

	if scrapeMode == "single" {
		file := utils.CreateFile(fileName)
		defer file.Close()
		page := utils.GetDocument(url)
		numberOfQuotes := utils.ScrapeData(page, &quotes)
		if numberOfQuotes != 0 {
			fmt.Println("Scraping:", numberOfQuotes)
			utils.WriteToFile(file, quotes)
		}
	} else if scrapeMode == "multi" {
		file := utils.CreateFile(fileName)
		defer file.Close()
		url, tag := utils.ParseUrl(url)
		for i := 1; i < 101; i++ {
			currentUrl := ""
			if tag == "" {
				currentUrl = url + "?page=" + strconv.Itoa(i)
			} else {
				currentUrl = url + tag + "?page=" + strconv.Itoa(i) + "&utf8=✓"
			}
			page := utils.GetDocument(currentUrl)
			numberOfQuotes := utils.ScrapeData(page, &quotes)
			if numberOfQuotes == 0 {
				fmt.Println("0 quotes found - terminating search. Url:", currentUrl)
				break
			}
			fmt.Println("Scraping:", currentUrl, numberOfQuotes, "quotes")
			time.Sleep(50 * time.Millisecond)
		}
		utils.WriteToFile(file, quotes)
	} else {
		log.Fatal("Please enter a valid scrape mode.")
	}

}
