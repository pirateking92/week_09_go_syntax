package flights

import (
	"fmt"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello from the flights package!")
}

type flight struct {
	Code       string
	Origin     string
	DueTime    time.Time
	Arrived    time.Time
	ExpectedAt time.Time
	Cancelled  bool
}

func NewFlight(code, origin string, dueTime, arrived, expectedAt time.Time, cancelled bool) flight {
	return flight{
		Code:       code,
		Origin:     origin,
		DueTime:    dueTime,
		Arrived:    arrived,
		ExpectedAt: expectedAt,
		Cancelled:  cancelled,
	}
}

// func NewFlight(code, origin string, dueTime, arrived, expectedAt time.Time, cancelled bool) flight {
// 	var arrivedTime, expectedAtTime time.Time

// 	if arrived.IsZero() {
// 		arrivedTime = time.Time{}
// 	} else {
// 		arrivedTime = arrived
// 	}

// 	if expectedAt.IsZero() {
// 		expectedAtTime = time.Time{}
// 	} else {
// 		expectedAtTime = expectedAt
// 	}

// 	return flight{
// 		Code:       code,
// 		Origin:     origin,
// 		DueTime:    dueTime,
// 		Arrived:    arrivedTime,
// 		ExpectedAt: expectedAtTime,
// 		Cancelled:  cancelled,
// 	}
// }

// func (f flight) ToString() string {
// 	formattedDueTime := f.DueTime.Format("15:04")
// 	formattedArrival := f.Arrived.Format("15:04")
// 	formattedExpectedAt := f.ExpectedAt.Format("15:04")
// 	return fmt.Sprintf(("Flight %s from %s is expected at %s"), f.Code, f.Origin, formattedDueTime, formattedArrival, formattedExpectedAt, f.Cancelled)
// }

// Define a type alias for the slice of flight structs
// this will link calls outside of this package
// to this, which then links to the flight struct
type Flights []flight
