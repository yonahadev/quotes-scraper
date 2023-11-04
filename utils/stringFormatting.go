package utils

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var quoteExpression = regexp.MustCompile(`“(.*?)”`) //.*? means non greedy - finds the first matching
var urlExpression = regexp.MustCompile(`\?page=\d+`)

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

func ParseUrl(url string) string {
	url = urlExpression.ReplaceAllLiteralString(url, "")
	fmt.Println("Checking all entries under:", url)
	return url
}
