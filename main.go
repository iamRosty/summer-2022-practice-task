package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

const jsonFile string = "./data.json"

type Trains []Train

type Train struct {
	TrainID            int
	DepartureStationID int
	ArrivalStationID   int
	Price              float32
	ArrivalTime        time.Time
	DepartureTime      time.Time
}

func (t *Train) UnmarshalJSON(data []byte) error {
	type tempStruct struct {
		TrainID            int
		DepartureStationID int
		ArrivalStationID   int
		Price              float32
		ArrivalTime        string
		DepartureTime      string
	}

	var target tempStruct
	err := json.Unmarshal(data, &target)
	if err != nil {
		return err
	}

	layout := "15:04:05"
	arrivalTime, err := time.Parse(layout, target.ArrivalTime)
	if err != nil {
		return err
	}
	departureTime, err := time.Parse(layout, target.DepartureTime)
	if err != nil {
		return err
	}

	t.TrainID = target.TrainID
	t.DepartureStationID = target.DepartureStationID
	t.ArrivalStationID = target.ArrivalStationID
	t.Price = target.Price
	t.DepartureTime = departureTime
	t.ArrivalTime = arrivalTime

	return nil
}

func main() {
	var departureStation, arrivalStation, criteria string

	fmt.Print("Type the station of departure and press enter: ")
	fmt.Scan(&departureStation)
	fmt.Print("Type the station of arrival and press enter: ")
	fmt.Scan(&arrivalStation)
	fmt.Print("Use ONE of these keywords:'price': cheaper first, 'arrival-time': first those that arrive earlier, 'departure-time': first those that depart earlier.\nType the criteria for selecting trains that suit you: ")
	fmt.Scan(&criteria)

}

func FindTrains(departureStation, arrivalStation, criteria string) (Trains, error) {
	// ... код
	return nil, nil // маєте повернути правильні значення
}

func readJson(s string) []byte {
	jsonData, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	if !json.Valid(jsonData) {
		log.Fatal("invalid json!")
	}
	return jsonData
}
