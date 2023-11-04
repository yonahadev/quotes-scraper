package utils

import (
	"fmt"
)

func GetInput(prompt string) string {
	fmt.Print(prompt)

	var userInput string
	fmt.Scan(&userInput)

	return userInput
}
