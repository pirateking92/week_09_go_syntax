package main

import (
	"fmt"
	"go_arrivals/flights"
	"time"

	"github.com/rodaine/table"
)

func main() {
	fmt.Println("Soon to be Arrivals!!")

	flight1 := flights.NewFlight("JA123", "Vancouver", time.Now())

	flight2 := flights.NewFlight("Ma77", "Tokyo", time.Date(2023, time.April, 1, 16, 0, 0, 0, time.UTC))

	flight3 := flights.NewFlight("JA 123", "Kyoto", time.Date(2023, time.April, 1, 12, 0, 0, 0, time.UTC))

	fmt.Println(flight1.ToString())
	fmt.Println(flight2.ToString())
	fmt.Println(flight3.ToString())

	// create a flight slice
	flightSlice := flights.Flights{flight1, flight2, flight3}

	// creating a board instance that can be used by main

	myBoard := flights.NewBoard(flightSlice)
	fmt.Println(myBoard.Display())

	tbl := table.New("Code", "Origin", "Time")

	for _, f := range myBoard.Flights {
		formattedTime := f.DueTime.Format("15:04")
		tbl.AddRow(f.Code, f.Origin, formattedTime)
	}

	tbl.Print()
	// flights.HelloWorld()
}
