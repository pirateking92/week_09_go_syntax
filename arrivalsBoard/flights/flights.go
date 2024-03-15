package flights

import (
	"fmt"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello from the flights package!")
}

type flight struct {
	Code    string
	Origin  string
	DueTime time.Time
}

func NewFlight(code, origin string, dueTime time.Time) flight {
	return flight{
		Code:    code,
		Origin:  origin,
		DueTime: dueTime,
	}
}

func (f flight) ToString() string {
	formattedTime := f.DueTime.Format("15:04")
	return fmt.Sprintf(("Flight %s from %s is expected at %s"), f.Code, f.Origin, formattedTime)
}

// Define a type alias for the slice of flight structs
type Flights []flight

type board struct {
	Flights []flight
}

func NewBoard(flights []flight) board {
	return board{
		Flights: flights,
	}
}

func (b board) Display() string {
	var result string
	result += "Time From Code\n"
	for _, f := range b.Flights {
		formattedTime := f.DueTime.Format("15:04")
		result += fmt.Sprintf("%s %s %s\n", formattedTime, f.Origin, f.Code)
	}
	return result
}
