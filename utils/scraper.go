package utils

import (
	"scraper/structs"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeData(page *goquery.Document) ([]structs.Quote, int) {

	var scrapedQuotes []structs.Quote

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

		scrapedQuotes = append(scrapedQuotes, quote)
	})

	return scrapedQuotes, numberOfQuotes
}
