package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Subcategory struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Category Category `json:"category"`
}

func main() {
	// Open and read the JSON file
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Read the file content into a byte slice
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	// Unmarshal the JSON data into a slice of Subcategory structs
	var subcategories []Subcategory
	err = json.Unmarshal(bytes, &subcategories)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %s", err)
	}

	// Create and open a file to write the output
	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Failed to create output file: %s", err)
	}
	defer outputFile.Close()

	// Write the extracted data to the file
	for i, subcategory := range subcategories {
		output := fmt.Sprintf("Category #%d\n", i+1)
		output += fmt.Sprintf("Main Category: %s\n", subcategory.Category.Name)
		output += fmt.Sprintf("Subcategory: %s\n\n", subcategory.Name)

		_, err := outputFile.WriteString(output)
		if err != nil {
			log.Fatalf("Failed to write to output file: %s", err)
		}
	}

	fmt.Println("Data has been successfully written to output.txt")
}
