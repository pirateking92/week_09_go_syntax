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
	result += "Time From Code\n"
	for _, f := range b.Flights {
		formattedTime := f.DueTime.Format("15:04")
		result += fmt.Sprintf("%s %s %s\n", formattedTime, f.Origin, f.Code)
	}
	return result
}
