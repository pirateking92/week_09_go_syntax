package flights

import (
	"fmt"
	"go_arrivals/flights"
	"go_arrivals/test_utils"
	"testing"
	"time"
)

func TestOutput(t *testing.T) {
	recording := test_utils.StartRecording()
	HelloWorld()
	result := test_utils.EndRecording(recording)

	expected := "Hello from the flights package!"

	if result != expected {
		t.Errorf("result is %v but %v was expected", result, expected)
	}
}

// //	func TestToString(t *testing.T) {
// //		flightNumber := "BA 341"
// //		origin := "London"
// //		scheduledTime := time.Date(2024, time.April, 1, 14, 0, 0, 0, time.UTC)
// //		flight1 := Flight{flightNumber, origin, scheduledTime}
// //		result, expected := Flight.ToString(flight1), "Flight BA 341 from London is expected at 14:00"
// //		if result != expected {
// //			t.Errorf("Result is %v when %v is expected", result, expected)
// //		}
// //	}

func TestToStringUtil(t *testing.T) {
	flight := flights.Flight{
		DueTime: time.Date(2024, time.April, 1, 14, 0, 0, 0, time.UTC),
		Origin:  "London",
		Code:    "BA 114",
	}

	recording := test_utils.StartRecording()
	defer test_utils.EndRecording(recording)

	result := flight.ToString()
	fmt.Println(result)

	capturedOutput := test_utils.EndRecording(recording)
	expected := "Flight BA 114 from London is expected at 14:00"

	if capturedOutput != expected {
		t.Errorf("result is %v but %v was expected", capturedOutput, expected)
	}
}

// func TestDisplayUtil(t *testing.T) {
// 	flights := []Flight{
// 		{DueTime: time.Date(2024, time.April, 1, 14, 0, 0, 0, time.UTC), Origin: "London", Code: "BA 114"},
// 		{DueTime: time.Date(2024, time.April, 1, 16, 0, 0, 0, time.UTC), Origin: "Berlin", Code: "LH 888"},
// 		{DueTime: time.Date(2024, time.April, 1, 18, 0, 0, 0, time.UTC), Origin: "Tokyo", Code: "JA 903"},
// 	}
// 	b := Board{Flights: flights}
// 	recording := test_utils.StartRecording()
// 	defer test_utils.EndRecording(recording)

// 	result := b.Display()
// 	fmt.Println(result)

// 	capturedOutput := test_utils.EndRecording(recording)
// 	expected := `Time From Code
// 14:00 London BA 114
// 16:00 Berlin LH 888
// 18:00 Tokyo JA 903
// 	`

// 	// trimming whitespace at the end and beginning
// 	if strings.TrimSpace(capturedOutput) != strings.TrimSpace(expected) {
// 		t.Errorf("Display() = %q, expected %q", capturedOutput, expected)
// 	}
// }

// // func TestToStringUtil(t *testing.T) {
// // 	recording := test_utils.StartRecording()
// // 	flight2 := []Flight{
// // 		{DueTime: time.Date(2024, time.April, 1, 14, 0, 0, 0, time.UTC), Origin: "London", Code: "BA 114"},
// // 		{DueTime: time.Date(2024, time.April, 1, 16, 0, 0, 0, time.UTC), Origin: "Berlin", Code: "LH 888"},
// // 		{DueTime: time.Date(2024, time.April, 1, 18, 0, 0, 0, time.UTC), Origin: "Tokyo", Code: "JA 903"},
// // 	}
// // 	Flight.ToString(flight2)
// // 	result := test_utils.EndRecording(recording)

// // 	expected := "Flight BA 341 from London is expected at 14:00"

// // 	if result != expected {
// // 		t.Errorf("result is %v but %v was expected", result, expected)
// // 	}
// // }

// func TestDisplay(t *testing.T) {
// 	flights := []Flight{
// 		{DueTime: time.Date(2024, time.April, 1, 14, 0, 0, 0, time.UTC), Origin: "London", Code: "BA 114"},
// 		{DueTime: time.Date(2024, time.April, 1, 16, 0, 0, 0, time.UTC), Origin: "Berlin", Code: "LH 888"},
// 		{DueTime: time.Date(2024, time.April, 1, 18, 0, 0, 0, time.UTC), Origin: "Tokyo", Code: "JA 903"},
// 	}
// 	b := Board{Flights: flights}

// 	expected := `Time From Code
// 14:00 London BA 114
// 16:00 Berlin LH 888
// 18:00 Tokyo JA 903
// `
// 	result := b.Display()

// 	if result != expected {
// 		t.Errorf("Display() = %q, expected %q", result, expected)
// 	}
// }
