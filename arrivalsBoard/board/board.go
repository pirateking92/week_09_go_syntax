package board

import (
	"fmt"
	"go_arrivals/flights"
)

type board struct {
	Flights flights.Flights
}

func NewBoard(flights flights.Flights) board {
	return board{
		Flights: flights,
	}
}

func (b board) Display() string {
	var result string
	result += "Time Code From Status\n"
	for _, f := range b.Flights {
		formattedDueTime := f.DueTime.Format("15:04")
		formattedArrival := f.Arrived.Format("15:04")
		formattedExpectedAt := f.ExpectedAt.Format("15:04")
		if f.Cancelled {
			result += fmt.Sprintf("%s %s %s cancelled\n", formattedDueTime, f.Origin, f.Code)
		} else if f.Arrived.IsZero() {
			result += fmt.Sprintf("%s %s %s expected %s\n", formattedDueTime, f.Origin, f.Code, formattedExpectedAt)
		} else {
			result += fmt.Sprintf("%s %s %s landed %s\n", formattedDueTime, f.Origin, f.Code, formattedArrival)
		}
	}
	return result
}

// func Board() string {
// 	flts, err := json.ReadFromJSON("flightData.json")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return ""
// 	}
// 	// SETTING UP THE TABLE WITH THE COLUMNS WANTED
// 	tbl := table.New("Code", "Origin", "Time", "Arrived", "Expected", "Cancelled")

// 	for _, f := range flts {
// 		formattedTime := f.DueTime.Format("15:04")
// 		var formattedArrival, formattedExpectedAt string
// 		if !f.Arrived.IsZero() {
// 			formattedArrival = f.Arrived.Format("15:04")
// 		}

// 		if !f.ExpectedAt.IsZero() {
// 			formattedExpectedAt = f.ExpectedAt.Format("15:04")
// 		}
// 		tbl.AddRow(f.Code, f.Origin, formattedTime, formattedArrival, formattedExpectedAt, f.Cancelled)
// 	}

// 	var buf bytes.Buffer
// 	tbl.Print(&buf)
// 	return buf.String()
// }
