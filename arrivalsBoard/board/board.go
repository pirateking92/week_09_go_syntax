package board

import (
	"bytes"
	"fmt"
	"go_arrivals/flights"

	"go_arrivals/json"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func center(s string, w int) string {
	return fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s))
}

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
			result += fmt.Sprintf("%s %s %s %s\n", formattedDueTime, f.Origin, f.Code, color.RedString("Cancelled"))
		} else if f.Arrived.IsZero() {
			result += fmt.Sprintf("%s %s %s %s %s\n", formattedDueTime, f.Origin, f.Code, color.GreenString("Expected"), formattedExpectedAt)
		} else {
			result += fmt.Sprintf("%s %s %s %s %s\n", formattedDueTime, f.Origin, f.Code, color.YellowString("Landed"), formattedArrival)
		}
	}
	return result
}

func (b board) Board() string {

	// SETTING UP THE TABLE WITH THE COLUMNS WANTED
	tbl := table.New("Code", "Origin", "Time", "Arrived", "Expected", "Cancelled")

	flts, err := json.ReadFromJSON("flightData.json")
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, f := range flts {
		formattedTime := f.DueTime.Format("15:04")
		var formattedArrival, formattedExpectedAt string
		if !f.Arrived.IsZero() {
			formattedArrival = f.Arrived.Format("15:04")
		}

		if !f.ExpectedAt.IsZero() {
			formattedExpectedAt = f.ExpectedAt.Format("15:04")
		}
		tbl.AddRow(f.Code, f.Origin, formattedTime, formattedArrival, formattedExpectedAt, f.Cancelled)
	}

	var buf bytes.Buffer
	tbl.Print()
	return buf.String()
}

func DisplayAndBoard() {
	fmt.Println("\n   ", center("Arrivals", 20), "   \n")

	flts, err := json.ReadFromJSON("flightData.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	b := NewBoard(flts)
	displayResult := b.Display()
	fmt.Println(displayResult)

	a := NewBoard(flts)
	boardResult := a.Board()
	fmt.Println(boardResult)

}
