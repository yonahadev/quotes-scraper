package utils

import (
	"regexp"
)

var re = regexp.MustCompile(`“(.*?)”`) //.*? means non greedy - finds the first matching

func ParseQuote(text string) string {
	slice := re.FindStringSubmatch(text)
	return slice[0]
}
