package main

import (
	"fmt"
	"scraper/utils"
)

func main() {

	url := utils.GetInput()

	page := utils.ProcessUrl(url)

	quotes := utils.ScrapeData(page)

	fmt.Println(quotes)
}
