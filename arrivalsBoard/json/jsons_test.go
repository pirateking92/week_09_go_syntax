package json

import (
	"go_arrivals/flights"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

func TestReadFromJSON(t *testing.T) {
	testCases := []struct {
		name     string
		filePath string
		expected flights.Flights
	}{
		{
			name:     "Read flights from JSON file",
			filePath: filepath.Join("testdata", "flights.json"),
			expected: flights.Flights{
				flights.NewFlight("BA1393", "Manchester", time.Date(2024, 3, 14, 13, 5, 0, 0, time.UTC)),
				flights.NewFlight("BA276", "Hyderabad", time.Date(2024, 3, 14, 12, 45, 0, 0, time.UTC)),
				flights.NewFlight("WS18", "Calgary", time.Date(2024, 3, 14, 12, 50, 0, 0, time.UTC)),
				flights.NewFlight("LX332", "Zurich", time.Date(2024, 3, 14, 13, 00, 0, 0, time.UTC)),
				flights.NewFlight("QR7", "Doha", time.Date(2024, 3, 14, 13, 20, 0, 0, time.UTC)),
				flights.NewFlight("VS20", "San Francisco", time.Date(2024, 3, 14, 12, 55, 0, 0, time.UTC)),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			flights, err := ReadFromJSON(tc.filePath)
			if err != nil {
				t.Errorf("ReadFromJSON(%q) returned an error: %v", tc.filePath, err)
				return
			}

			if !reflect.DeepEqual(flights, tc.expected) {
				t.Errorf("ReadFromJSON(%q) returned unexpected flights:\nGot: %v\nExpected: %v", tc.filePath, flights, tc.expected)
			}
		})
	}
}
