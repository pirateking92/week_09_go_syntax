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
// this will link calls outside of this package
// to this, which then links to the flight struct
type Flights []flight
