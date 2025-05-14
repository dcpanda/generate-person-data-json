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

func ParseFlags() (int, string) {
	numRecords := flag.Int("n", 100, "Number of fake records to generate")
	outputFilename := flag.String("o", "fake_pii_data.json", "Output filename for the JSON data")
	flag.Parse()
	return *numRecords, *outputFilename
}

func CalculateAge(birthday string) (int, error) {
	birthDate, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return 0, fmt.Errorf("error parsing birthday: %w", err)
	}
	age := math.Ceil(time.Since(birthDate).Hours() / 24 / 365)
	return int(age), nil
}

func GeneratePerson() (Person, error) {
	gofakeit.Seed(0)
	address := gofakeit.Address()
	birthdate := gofakeit.DateRange(time.Now().AddDate(-100, 0, 0), time.Now()).Format("2006-01-02")

	age, err := CalculateAge(birthdate)
	if err != nil {
		return Person{}, err
	}

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
		Age:          age,
	}

	return person, nil
}

func GeneratePersonRecords(numRecords int) ([]Person, error) {
	var personRecords []Person
	for i := 0; i < numRecords; i++ {
		person, err := GeneratePerson()
		if err != nil {
			return nil, fmt.Errorf("error generating person record: %w", err)
		}
		personRecords = append(personRecords, person)
	}
	return personRecords, nil
}

func MarshalToJSON(personRecords []Person) ([]byte, error) {
	jsonData, err := json.MarshalIndent(personRecords, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON: %w", err)
	}
	return jsonData, nil
}

func WriteJSONToFile(filename string, jsonData []byte) error {
	err := os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing JSON to file: %w", err)
	}
	return nil
}

func main() {
	numRecords, outputFilename := ParseFlags()
	fmt.Printf("Generating %d fake records into %s...\n", numRecords, outputFilename)

	personRecords, err := GeneratePersonRecords(numRecords)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	jsonData, err := MarshalToJSON(personRecords)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	err = WriteJSONToFile(outputFilename, jsonData)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated %d records and saved to %s \n", numRecords, outputFilename)
}
