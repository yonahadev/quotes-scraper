package utils

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var quoteExpression = regexp.MustCompile(`“(.*?)”`) //.*? means non greedy - finds the first matching
var urlExpression = regexp.MustCompile(`(\?page=\d+)|(#.+)`)
var utf8Tag = regexp.MustCompile(`tag\?.*=`)
var utf8Paginated = regexp.MustCompile(`&utf8=.*`)

func ParseQuote(text string) string {
	slice := quoteExpression.FindStringSubmatch(text)
	return slice[0]
}

func ParseLikes(text string) int {
	text = strings.Replace(text, "likes", "", -1)
	text = strings.TrimSpace(text)
	likes, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}
	return likes
}

func ParseAuthor(text string) string {
	text = strings.Replace(text, ",", "", -1)
	text = strings.TrimSpace(text)
	return text
}

func ParseSource(text string) string {
	if text == "" {
		text = "unknown"
	} else {
		text = strings.TrimSpace(text)
	}
	return text
}

func ParseTags(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.Replace(text, "tags:", "", -1)
	slice := strings.Split(text, ",")
	for i := 0; i < len(slice); i++ {
		slice[i] = strings.TrimSpace(slice[i])
	}
	return slice
}

func ParseUrl(url string) (string, string) {
	url = urlExpression.ReplaceAllLiteralString(url, "")
	if len(utf8Paginated.FindAllString(url, -1)) > 0 {
		fmt.Println("pagination match")
		url = utf8Paginated.ReplaceAllLiteralString(url, "")
		tag := strings.Replace(url, "https://www.goodreads.com/quotes/tag/", "", -1)
		url = "https://www.goodreads.com/quotes/tag/"
		fmt.Println("Checking entries under custom tag:", tag)
		return url, tag
	} else if len(utf8Tag.FindAllString(url, -1)) > 0 {
		url = utf8Tag.ReplaceAllLiteralString(url, "")
		tag := strings.Replace(url, "https://www.goodreads.com/quotes/", "", -1)
		url = "https://www.goodreads.com/quotes/tag/"
		fmt.Println("Checking entries under custom tag:", tag)
		return url, tag
	}
	return url, ""
}
