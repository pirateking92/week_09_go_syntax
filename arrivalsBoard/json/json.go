package json

import (
	"encoding/json"
	"fmt"
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
		Status           struct {
			Arrived    string `json:"arrived"`
			ExpectedAt string `json:"expected_at"`
			Cancelled  bool   `json:"cancelled"`
		} `json:"status"`
	} `json:"flights"`
	Airports struct {
		DOH string `json:"Doha`
		HYD string `json:"Hyderabad`
		LHR string `json:"London"`
		MAN string `json:"Manchester"`
		SFO string `json:"San Francisco`
		YYC string `json:"Calgary"`
		ZRH string `json:"Zurich"`
	} `json:"airports"`
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
		Arrived, _ := time.Parse(time.RFC3339, f.Status.Arrived)
		ExpectedAt, _ := time.Parse(time.RFC3339, f.Status.ExpectedAt)
		flts = append(flts, flights.NewFlight(f.Code, f.Origin, ScheduledArrival, Arrived, ExpectedAt, f.Status.Cancelled))
	}

	return flts, nil
}

func getAirportName(code string, airports jsonData) string {
	switch code {
	case "DOH":
		return airports.Airports.DOH
	case "HYD":
		return airports.Airports.HYD
	case "LHR":
		return airports.Airports.LHR
	case "MAN":
		return airports.Airports.MAN
	case "SFO":
		return airports.Airports.SFO
	case "YYC":
		return airports.Airports.YYC
	case "ZRH":
		return airports.Airports.ZRH
	default:
		return ""
	}
}
func CodeReadFromJSON(filePath string, airportCode string) (flights.Flights, error) {
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
	airportName := getAirportName(airportCode, data)
	if airportName == "" {
		return nil, fmt.Errorf("invalid airport code: %s", airportCode)
	}

	for _, f := range data.Flights {
		if f.Origin == airportName {
			ScheduledArrival, _ := time.Parse(time.RFC3339, f.ScheduledArrival)
			Arrived, _ := time.Parse(time.RFC3339, f.Status.Arrived)
			ExpectedAt, _ := time.Parse(time.RFC3339, f.Status.ExpectedAt)
			flts = append(flts, flights.NewFlight(f.Code, f.Origin, ScheduledArrival, Arrived, ExpectedAt, f.Status.Cancelled))
		}
	}

	return flts, nil
}

func PrintCodeReadFromJSON() {
	airportCode := os.Args[1]
	flights, err := CodeReadFromJSON("flightData.json", airportCode)
	if err != nil {
		fmt.Println("Error reading flights: %v\n", err)
		return
	}
	fmt.Printf("Flights for given airport code: %v\n", airportCode)
	for _, flight := range flights {
		fmt.Println(flight.ToString(flight))
	}
}
