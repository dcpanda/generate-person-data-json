package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dcpanda/generate-person-data-json/pkg/api"
	"github.com/dcpanda/generate-person-data-json/pkg/person"
)

func main() {
	// Define command-line flags
	numRecords := flag.Int("n", 100, "Number of fake records to generate")
	outputFilename := flag.String("o", "fake_pii_data.json", "Output filename for the JSON data")
	serverMode := flag.Bool("server", false, "Run as HTTP server")
	port := flag.String("port", ":8080", "Port for HTTP server")

	flag.Parse()

	// Run in server mode if specified
	if *serverMode {
		fmt.Printf("Starting HTTP server on %s...\n", *port)
		fmt.Println("API endpoints:")
		fmt.Println("  GET /health - Health check")
		fmt.Println("  GET /api/persons?n=<number> - Generate person data (default n=10)")

		err := api.StartServer(*port)
		if err != nil {
			fmt.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
		return
	}

	// Otherwise, run in CLI mode (backward compatibility)
	fmt.Printf("Generating %d fake records into %s...\n", *numRecords, *outputFilename)

	// Generate person records
	personRecords, err := person.GeneratePersonRecords(*numRecords)
	if err != nil {
		fmt.Printf("Error generating person records: %s\n", err)
		os.Exit(1)
	}

	// Convert to JSON
	jsonData, err := person.MarshalToJSON(personRecords)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %s\n", err)
		os.Exit(1)
	}

	// Write to file
	err = os.WriteFile(*outputFilename, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing JSON to file: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated %d records and saved to %s\n", *numRecords, *outputFilename)
}
