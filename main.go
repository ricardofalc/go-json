package main

import (
	"encoding/csv"
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

	writeGroupedOutput("grouped_output.txt", subcategories)
	writeOriginalOutput("original_output.txt", subcategories)
	writeCSVOutput("output.csv", subcategories)

	fmt.Println("Data has been successfully written to grouped_output.txt, original_output.txt, and output.csv")
}

func writeGroupedOutput(fileName string, subcategories []Subcategory) {
	// ... (keep the existing writeGroupedOutput function as is)
}

func writeOriginalOutput(fileName string, subcategories []Subcategory) {
	// ... (keep the existing writeOriginalOutput function as is)
}

func writeCSVOutput(fileName string, subcategories []Subcategory) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to create CSV file: %s", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Main Category ID", "Main Category Name", "Subcategory ID", "Subcategory Name"}
	if err := writer.Write(header); err != nil {
		log.Fatalf("Error writing CSV header: %s", err)
	}

	// Write data
	for _, subcategory := range subcategories {
		row := []string{
			subcategory.Category.ID,
			subcategory.Category.Name,
			subcategory.ID,
			subcategory.Name,
		}
		if err := writer.Write(row); err != nil {
			log.Fatalf("Error writing CSV row: %s", err)
		}
	}
}
