package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Person struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zip_code"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Birthday     string `json:"birthday"`
	Age          int    `json:"age"`
}

func main() {
	numRecords := flag.Int("n", 1, "Number of fake records to generate")
	outputFilename := flag.String("o", "fake_pii_data.json", "Output filename for the JSON data")

	flag.Parse()
	fmt.Printf("Generating %d fake records into %s...\n", *numRecords, *outputFilename)

	var personRecords []Person
	for i := 0; i < *numRecords; i++ {
		person := Person{
			FirstName:    "John",
			LastName:     "Doe",
			Gender:       "Male",
			AddressLine1: "123 Main St",
			AddressLine2: "Apt 1",
			City:         "New York",
			State:        "NY",
			ZipCode:      "10001",
			Phone:        "123-456-7890",
		}
		personRecords = append(personRecords, person)
	}

	jsonData, err := json.MarshalIndent(personRecords, "", "  ")
	if err != nil {
		fmt.Printf("Error Marshalling JSON: %s\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(*outputFilename, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing JSON to file: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated %d records and saved to %s \n", *numRecords, *outputFilename)
}
