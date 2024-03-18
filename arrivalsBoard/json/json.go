package json

import (
	"encoding/json"
	"go_arrivals/flights"
	"os"
	"time"
)

// importing at the beginning isnt always
// necessary as the imports will add automatically
// later on when you save

type jsonData struct {
	Flights []struct {
		Origin           string `json:"from"`
		Code             string `json:"code"`
		ScheduledArrival string `json:"scheduled_arrival"` // these are json maps to what they correspond to in the their respective json file
	} `json:"flights"`
}

func ReadFromJSON(filePath string) (flights.Flights, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data jsonData
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, err
	}

	var flts flights.Flights
	for _, f := range data.Flights {
		ScheduledArrival, _ := time.Parse(time.RFC3339, f.ScheduledArrival)
		flts = append(flts, flights.NewFlight(f.Code, f.Origin, ScheduledArrival))
	}

	return flts, nil
}
