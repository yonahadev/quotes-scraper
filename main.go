package main

import (
	"scraper/utils"
)

func main() {

	// url := utils.GetInput("Enter a valid goodreads quote url:")

	fileName := utils.GetInput("Enter a filename:")

	page := utils.ProcessUrl("https://www.goodreads.com/quotes/tag/philosophy?page=1")

	quotes := utils.ScrapeData(page)

	file := utils.CreateFile(fileName)
	defer file.Close()

	utils.WriteToFile(file, quotes)
}
