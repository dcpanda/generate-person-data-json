package person

import (
	"encoding/json"
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {

	//The test cases uses time.Now() to ensure we always get the same age.
	testCases := []struct {
		name     string
		birthday string
		want     int
		wantErr  bool
	}{
		{
			name:     "Valid birthday - 30 years ago",
			birthday: time.Now().AddDate(-30, 0, 0).Format("2006-01-02"),
			want:     31, //rounded
			wantErr:  false,
		},
		{
			name:     "Valid birthday - 10 years ago",
			birthday: time.Now().AddDate(-10, 0, 0).Format("2006-01-02"),
			want:     11, //rounded
			wantErr:  false,
		},
		{
			name:     "Invalid birthday format",
			birthday: "01/01/2000",
			want:     0,
			wantErr:  true,
		},
		{
			name:     "Empty birthday",
			birthday: "",
			want:     0,
			wantErr:  true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateAge(tt.birthday)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateAge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("CalculateAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeneratePerson(t *testing.T) {
	person, err := GeneratePerson()
	if err != nil {
		t.Fatalf("GeneratePerson() error = %v", err)
	}

	// Check that all fields are populated. This is more to test the package than the code.
	if person.FirstName == "" {
		t.Error("GeneratePerson() FirstName is empty")
	}
	if person.LastName == "" {
		t.Error("GeneratePerson() LastName is empty")
	}
	if person.Gender == "" {
		t.Error("GeneratePerson() Gender is empty")
	}
	if person.AddressLine1 == "" {
		t.Error("GeneratePerson() AddressLine1 is empty")
	}
	if person.City == "" {
		t.Error("GeneratePerson() City is empty")
	}
	if person.State == "" {
		t.Error("GeneratePerson() State is empty")
	}
	if person.ZipCode == "" {
		t.Error("GeneratePerson() ZipCode is empty")
	}
	if person.Phone == "" {
		t.Error("GeneratePerson() Phone is empty")
	}
	if person.Email == "" {
		t.Error("GeneratePerson() Email is empty")
	}
	if person.Birthday == "" {
		t.Error("GeneratePerson() Birthday is empty")
	}
	if person.Age <= 0 {
		t.Errorf("GeneratePerson() Age = %v, want > 0", person.Age)
	}
}

func TestGeneratePersonRecords(t *testing.T) {

	// You can add 100000 as numRecords for large volume scenario,
	// but this is a unit testing and not load testing
	testCases := []struct {
		name       string
		numRecords int
		wantErr    bool
	}{
		{
			name:       "Generate 1 record",
			numRecords: 1,
			wantErr:    false,
		},
		{
			name:       "Generate 5 records",
			numRecords: 5,
			wantErr:    false,
		},
		{
			name:       "Generate 0 records",
			numRecords: 0,
			wantErr:    false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GeneratePersonRecords(tt.numRecords)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratePersonRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) != tt.numRecords {
				t.Errorf("GeneratePersonRecords() returned %v records, want %v", len(got), tt.numRecords)
			}
		})
	}
}

func TestMarshalToJSON(t *testing.T) {

	person := Person{
		FirstName:    "James",
		LastName:     "Bond",
		Gender:       "Male",
		AddressLine1: "123 Main St",
		AddressLine2: "Apt 007",
		City:         "Anytown",
		State:        "CA",
		ZipCode:      "12345",
		Phone:        "555-123-4007",
		Email:        "james.bond@mi5.com",
		Birthday:     "1990-01-01",
		Age:          33,
	}

	t.Run("Single person", func(t *testing.T) {
		personRecords := []Person{person}
		jsonData, err := MarshalToJSON(personRecords)
		if err != nil {
			t.Fatalf("MarshalToJSON() error = %v", err)
		}

		// Unmarshal the JSON to verify it's valid
		var unmarshaledRecords []Person
		err = json.Unmarshal(jsonData, &unmarshaledRecords)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON: %v", err)
		}

		if len(unmarshaledRecords) != 1 {
			t.Errorf("Expected 1 record, got %d", len(unmarshaledRecords))
		}

		// Verify the unmarshaled data matches the original
		if unmarshaledRecords[0].FirstName != person.FirstName {
			t.Errorf("FirstName = %v, want %v", unmarshaledRecords[0].FirstName, person.FirstName)
		}
	})

}
