package utils

import (
	"regexp"
)

func ParseQuote(text string) string {
	re := regexp.MustCompile(`“(.*?)”`) //.*? means non greedy - finds the first matching
	slice := re.FindStringSubmatch(text)
	return slice[0]
}
