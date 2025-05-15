package api

import (
	"encoding/json"
	"github.com/dcpanda/generate-person-data-json/pkg/person"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetupRouter(t *testing.T) {
	router := SetupRouter()

	if router == nil {
		t.Fatal("SetupRouter() returned nil")
	}
}

func TestHealthEndpoint(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Setup the router
	router := SetupRouter()

	// Create a test HTTP recorder
	w := httptest.NewRecorder()

	// Create a test request for health endpoint
	req, _ := http.NewRequest("GET", "/health", nil)

	// Serve the request
	router.ServeHTTP(w, req)

	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if status, exists := response["status"]; !exists || status != "ok" {
		t.Errorf("Expected status to be 'ok', got '%s'", status)
	}
}

func TestGetPersonsEndpoint(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Setup the router
	router := SetupRouter()

	// Test cases for testing the end points
	testCases := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "Default number of records",
			queryParam:     "",
			expectedStatus: http.StatusOK,
			expectedCount:  10,
		},
		{
			name:           "Custom number of records",
			queryParam:     "n=5",
			expectedStatus: http.StatusOK,
			expectedCount:  5,
		},
		{
			name:           "Invalid number format",
			queryParam:     "n=abc",
			expectedStatus: http.StatusBadRequest,
			expectedCount:  0,
		},
		{
			name:           "Zero records",
			queryParam:     "n=0",
			expectedStatus: http.StatusBadRequest,
			expectedCount:  0,
		},
		{
			name:           "Negative records",
			queryParam:     "n=-5",
			expectedStatus: http.StatusBadRequest,
			expectedCount:  0,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test HTTP recorder
			w := httptest.NewRecorder()

			// Create a test request
			personsURL := "/api/persons"
			if tt.queryParam != "" {
				personsURL += "?" + tt.queryParam
			}
			req, _ := http.NewRequest("GET", personsURL, nil)

			// submit request
			router.ServeHTTP(w, req)

			// Check the returned status code with the expected results
			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, w.Code)
			}

			// If response is successful, check the number of records with expectedCount
			if tt.expectedStatus == http.StatusOK {
				var personRecords []person.Person
				err := json.Unmarshal(w.Body.Bytes(), &personRecords)
				if err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if len(personRecords) != tt.expectedCount {
					t.Errorf("Expected %d records, got %d", tt.expectedCount, len(personRecords))
				}
			}
		})
	}
}
