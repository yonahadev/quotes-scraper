package utils

import (
	"fmt"
)

func GetInput() string {
	fmt.Println("Enter a valid goodreads quotes url:")

	var userInput string
	fmt.Scan(&userInput)

	return userInput
}
