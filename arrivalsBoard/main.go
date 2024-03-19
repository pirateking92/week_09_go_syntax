package main

import (
	"fmt"
	"go_arrivals/board"
)

func center(s string, w int) string {
	return fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s))
}

func main() {

	fmt.Println(">>>", center("MattPort", 20), "<<<")
	board.DisplayAndBoard()

	// commented this next line because if you dont
	// pass a cmd line arg into it, then it gives back
	// an err. i didnt add any err checking, which
	// is probably why.
	// json.PrintCodeReadFromJSON()

	// flight1 := flights.NewFlight("JA123", "Vancouver", time.Now())

	// flight2 := flights.NewFlight("Ma77", "Tokyo", time.Date(2023, time.April, 1, 16, 0, 0, 0, time.UTC))

	// flight3 := flights.NewFlight("JA 123", "Kyoto", time.Date(2023, time.April, 1, 12, 0, 0, 0, time.UTC))

	// fmt.Println(flight1.ToString())
	// fmt.Println(flight2.ToString())
	// fmt.Println(flight3.ToString())

	// // create a flight slice
	// flightSlice := flights.Flights{flight1, flight2, flight3}

	// // creating a board instance that can be used by main

	// myBoard := board.NewBoard(flightSlice)
	// fmt.Println(myBoard.Display())

	// for _, f := range myBoard.Flights {
	// 	formattedTime := f.DueTime.Format("15:04")
	// 	tbl.AddRow(f.Code, f.Origin, formattedTime)
	// }

	// fmt.Println("\n   ", center("Arrivals", 20), "   \n")

	// flts, err := json.ReadFromJSON("flightData.json")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// b := board.NewBoard(flts)
	// displayResult := b.Display()
	// fmt.Println(displayResult)

	// a := board.NewBoard(flts)
	// boardResult := a.Board()
	// fmt.Println(boardResult)

}
