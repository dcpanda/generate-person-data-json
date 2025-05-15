package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/dcpanda/generate-person-data-json/pkg/person"
)

func TestGenerateJSONFile(t *testing.T) {
	// Create a temporary directory for test output
	tempDir, err := os.MkdirTemp("", "generate-json-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after the test

	// Test Cases
	testCases := []struct {
		name       string
		numRecords int
	}{
		{
			name:       "Generate 1 record",
			numRecords: 1,
		},
		{
			name:       "Generate 5 records",
			numRecords: 5,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary output file
			outputFile := filepath.Join(tempDir, "test_output.json")

			personRecords, err := person.GeneratePersonRecords(tt.numRecords)
			if err != nil {
				t.Fatalf("Failed to generate person records: %v", err)
			}

			jsonData, err := person.MarshalToJSON(personRecords)
			if err != nil {
				t.Fatalf("Failed to marshal JSON: %v", err)
			}

			// Write to file
			err = os.WriteFile(outputFile, jsonData, 0644)
			if err != nil {
				t.Fatalf("Failed to write JSON to file: %v", err)
			}

			// Verify the file exists
			if _, err := os.Stat(outputFile); os.IsNotExist(err) {
				t.Fatalf("Output file was not created: %v", err)
			}

			// Read the file back
			fileData, err := os.ReadFile(outputFile)
			if err != nil {
				t.Fatalf("Failed to read output file: %v", err)
			}

			var unmarshaledRecords []person.Person
			err = json.Unmarshal(fileData, &unmarshaledRecords)
			if err != nil {
				t.Fatalf("Failed to unmarshal JSON from file: %v", err)
			}

			// Verify the number of records
			if len(unmarshaledRecords) != tt.numRecords {
				t.Errorf("Expected %d records, got %d", tt.numRecords, len(unmarshaledRecords))
			}
		})
	}
}
