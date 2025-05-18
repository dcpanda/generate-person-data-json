package person

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"math"
	"time"
)

// Person represents a person with various attributes.
// There is gofakeit struct. See https://github.com/brianvoe/gofakeit?tab=readme-ov-file#struct
// I have seen the more complex the data generation, the fake tag becomes more difficult to use.
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

// CalculateAge calculates a person's age based on their birthday
// Returns the age and an error if the birthday cannot be parsed
func CalculateAge(birthday string) (int, error) {
	birthDate, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return 0, fmt.Errorf("error parsing birthday: %w", err)
	}
	age := math.Ceil(time.Since(birthDate).Hours() / 24 / 365)
	return int(age), nil
}

// GeneratePerson creates a single fake person record
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

// GeneratePersonRecords generates a specified number of fake person records
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

// MarshalToJSON converts person records to JSON
func MarshalToJSON(personRecords []Person) ([]byte, error) {
	jsonData, err := json.MarshalIndent(personRecords, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON: %w", err)
	}
	return jsonData, nil
}
