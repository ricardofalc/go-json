package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
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
	writeCompactOutput("raw_output.txt", subcategories)
	writeCompactGroupedOutput("grouped_output_compact.txt", subcategories)

	fmt.Println("Data has been successfully written to grouped_output.txt, original_output.txt, raw_output.txt and output.csv")
}

func writeCompactOutput(fileName string, subcategories []Subcategory) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to create output file: %s", err)
	}
	defer file.Close()

	var outputParts []string

	for i, subcategory := range subcategories {
		part := fmt.Sprintf("main category #%d: %s subcategory: %s",
			i+1, subcategory.Category.Name, subcategory.Name)
		outputParts = append(outputParts, part)
	}

	output := strings.Join(outputParts, ", ")

	_, err = file.WriteString(output)
	if err != nil {
		log.Fatalf("Failed to write to output file: %s", err)
	}
}

func writeGroupedOutput(fileName string, subcategories []Subcategory) {
	// Create and open a file to write the output
	outputFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to create output file: %s", err)
	}
	defer outputFile.Close()

	// Group subcategories by main category
	categoryMap := make(map[string][]Subcategory)
	for _, subcategory := range subcategories {
		categoryMap[subcategory.Category.Name] = append(categoryMap[subcategory.Category.Name], subcategory)
	}

	// Get sorted main category names
	var mainCategories []string
	for category := range categoryMap {
		mainCategories = append(mainCategories, category)
	}
	sort.Strings(mainCategories)

	// Write the grouped data to the file
	var count int

	for i, mainCategory := range mainCategories {
		output := fmt.Sprintf("%d. Main Category: %s\n", i+1, mainCategory)
		output += "Subcategories:\n"

		for _, subcategory := range categoryMap[mainCategory] {
			output += fmt.Sprintf("- %s\n", subcategory.Name)
			count++
		}
		output += "\n Total Subcategories: " + fmt.Sprintf("%d\n\n", count)
		count = 0

		_, err := outputFile.WriteString(output)
		if err != nil {
			log.Fatalf("Failed to write to output file: %s", err)
		}
	}
}

func writeCompactGroupedOutput(fileName string, subcategories []Subcategory) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to create output file: %s", err)
	}
	defer file.Close()

	// Group subcategories by main category
	categoryMap := make(map[string][]string)
	for _, subcategory := range subcategories {
		categoryMap[subcategory.Category.Name] = append(categoryMap[subcategory.Category.Name], subcategory.Name)
	}

	// Get sorted main category names
	var mainCategories []string
	for category := range categoryMap {
		mainCategories = append(mainCategories, category)
	}
	sort.Strings(mainCategories)

	var outputParts []string

	for i, mainCategory := range mainCategories {
		subcategoryNames := strings.Join(categoryMap[mainCategory], ", ")
		part := fmt.Sprintf("main category #%d: %s subcategories: %s",
			i+1, mainCategory, subcategoryNames)
		outputParts = append(outputParts, part)
	}

	output := strings.Join(outputParts, ", ")

	_, err = file.WriteString(output)
	if err != nil {
		log.Fatalf("Failed to write to output file: %s", err)
	}
}

func writeOriginalOutput(fileName string, subcategories []Subcategory) {
	// Create and open a file to write the output
	outputFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to create output file: %s", err)
	}
	defer outputFile.Close()

	for i, subcategory := range subcategories {
		output := fmt.Sprintf("Category #%d\n", i+1)
		output += fmt.Sprintf("Main Category: %s\n", subcategory.Category.Name)
		output += fmt.Sprintf("Subcategory: %s\n\n", subcategory.Name)

		_, err := outputFile.WriteString(output)
		if err != nil {
			log.Fatalf("Failed to write to output file: %s", err)
		}
	}
}

func writeRawOutput(fileName string, subcategories []Subcategory) {
	// Create and open a file to write the output
	outputFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to create output file: %s", err)
	}
	defer outputFile.Close()

	for i, subcategory := range subcategories {
		output := fmt.Sprintf("Category #%d\n", i+1)
		output += fmt.Sprintf("Main Category: %s\n", subcategory.Category.Name)
		output += fmt.Sprintf("Subcategory: %s\n\n", subcategory.Name)

		_, err := outputFile.WriteString(output)
		if err != nil {
			log.Fatalf("Failed to write to output file: %s", err)
		}
	}
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
