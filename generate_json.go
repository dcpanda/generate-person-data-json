package main

import (
	"fmt"
)

type Person struct {
	FirstName    string
	LastName     string
	Gender       string
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	ZipCode      string
	Phone        string
	Email        string
	Birthday     string
	Age          int
}

func main() {
	numRecords := 100
	outputFilename := "fake_data.json"

	fmt.Printf("Generating %d fake records into %s...\n", numRecords, outputFilename)

}
