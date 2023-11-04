package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"scraper/structs"
)

func WriteToFile(file *os.File, quote []structs.Quote) {
	jsonEntry, err := json.MarshalIndent(quote, "", "    ")
	if err != nil {
		log.Fatal("Error converting entry to json")
	}
	n, err := file.Write(jsonEntry)
	if err != nil {
		log.Fatal("Error writing to file", err)
	} else {
		fmt.Println("Wrote", n, "bytes to file")
	}
}

func CreateFile(fileName string) *os.File {
	file, err := os.Create("output/" + fileName + ".json")
	if err != nil {
		log.Fatal("Could not open/create file")
	} else {
		fmt.Println("Created/accessed file:", "output/"+fileName+".json")
	}
	return file
}
