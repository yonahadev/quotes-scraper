package utils

import (
	"fmt"
	"scraper/structs"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeData(page *goquery.Document, quoteList *[]structs.Quote) int {

	numberOfQuotes := page.Find("div.quoteDetails").Length()

	page.Find("div.quoteDetails").Each(func(index int, element *goquery.Selection) {

		//find relevant dom text
		text := element.Text()
		author := element.Find("span.authorOrTitle").Text()
		source := element.Find("span").Find("a.authorOrTitle").Text()
		tagString := element.Find("div.quoteFooter").Find("div.greyText").Text()
		likes := element.Find("div.quoteFooter").Find("div.right").Text()

		//format  text correctly
		text = ParseQuote(text)
		source = ParseSource(source)
		author = ParseAuthor(author)
		tags := ParseTags(tagString)
		likeCount := ParseLikes(likes)

		quote := structs.Quote{Text: text, Author: author, Source: source, Tags: tags, Likes: likeCount}

		*quoteList = append(*quoteList, quote)
	})

	return numberOfQuotes
}

func SingleScrape(fileName string, url string, quotes []structs.Quote) {
	file := CreateFile(fileName)
	defer file.Close()
	page := GetDocument(url)
	numberOfQuotes := ScrapeData(page, &quotes)
	fmt.Println(numberOfQuotes)
	if numberOfQuotes != 0 {
		fmt.Println("Scraping:", numberOfQuotes)
		WriteToFile(file, quotes)
	} else {
		fmt.Println("0 quotes found - terminating search.")
	}
}

func MultiScrape(fileName string, url string, quotes []structs.Quote) {
	total := 0
	file := CreateFile(fileName)
	defer file.Close()
	url, tag := ParseUrl(url)
	for i := 1; i < 101; i++ {
		currentUrl := ""
		if tag == "" {
			currentUrl = url + "?page=" + strconv.Itoa(i)
		} else {
			currentUrl = url + tag + "?page=" + strconv.Itoa(i) + "&utf8=✓"
		}
		page := GetDocument(currentUrl)
		numberOfQuotes := ScrapeData(page, &quotes)
		if numberOfQuotes == 0 {
			fmt.Println(total, "quotes found - terminating search.")
			break
		}
		total += numberOfQuotes
		fmt.Println("Scraping:", currentUrl, numberOfQuotes, "quotes")
		time.Sleep(50 * time.Millisecond)
	}
	WriteToFile(file, quotes)
}
