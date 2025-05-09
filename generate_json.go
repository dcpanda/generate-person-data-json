package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"math"
	"os"
	"time"
)

type Person struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	State        string `json:"state_code"`
	ZipCode      string `json:"zip_code"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Birthday     string `json:"birthday"`
	Age          int    `json:"age"`
}

func CalculateAge(birthday string) int {
	birthDate, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		fmt.Printf("Error parsing birthday: %s\n", err)
		os.Exit(1)
	}
	age := math.Ceil(time.Since(birthDate).Hours() / 24 / 365)
	return int(age)
}
func main() {
	numRecords := flag.Int("n", 100, "Number of fake records to generate")
	outputFilename := flag.String("o", "fake_pii_data.json", "Output filename for the JSON data")

	flag.Parse()
	fmt.Printf("Generating %d fake records into %s...\n", *numRecords, *outputFilename)

	var personRecords []Person
	for i := 0; i < *numRecords; i++ {
		gofakeit.Seed(0)
		address := gofakeit.Address()
		birthdate := gofakeit.DateRange(time.Now().AddDate(-100, 0, 0), time.Now()).Format("2006-01-02")
		person := Person{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			Gender:       gofakeit.Gender(),
			AddressLine1: address.Street,
			AddressLine2: "",
			City:         address.City,
			State:        gofakeit.StateAbr(),
			ZipCode:      address.Zip,
			Phone:        gofakeit.Phone(),
			Email:        gofakeit.Email(),
			Birthday:     birthdate,
			Age:          CalculateAge(birthdate),
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
