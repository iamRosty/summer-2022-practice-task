package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Trains []Train

type Train struct {
	TrainID            int
	DepartureStationID int
	ArrivalStationID   int
	Price              float32
	ArrivalTime        time.Time
	DepartureTime      time.Time
}

// TrainID            int       `json: "trainId"`
// DepartureStationID int       `json: "departureStationID"`
// ArrivalStationID   int       `json: "arrivalStationID"`
// Price              float32   `json: "price"`
// ArrivalTime        time.Time `json: "arrivalTime"`
// DepartureTime      time.Time `json: "departureTime"`

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

	// var departureStation, arrivalStation, criteria string

	// fmt.Print("Type the station of departure and press enter: ")
	// fmt.Scan(&departureStation)
	// fmt.Print("Type the station of arrival and press enter: ")
	// fmt.Scan(&arrivalStation)
	// fmt.Print("Type the criteria for selecting trains that suit you.\n Use ONE of these keywords:'price', 'arrival-time' and  'departure-time'.\n Price- cheaper first, arrival-time - first those that arrive earlier, departure-time - first those that depart earlier: ")
	// fmt.Scan(&criteria)

	jsonFile, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var trainsBase Trains
	createTrainsBase(jsonData, &trainsBase)

	//result, err := FindTrains(departureStation, arrivalStation, criteria))
	//	... обробка помилки
	//	... друк result
	fmt.Println(trainsBase)
}

func FindTrains(departureStation, arrivalStation, criteria string) (Trains, error) {

	return nil, nil // маєте повернути правильні значення
}
func createTrainsBase(data []byte, t *Trains) {
	if !json.Valid(data) {
		fmt.Println("invalid json!")
	}
	err := json.Unmarshal(data, &t)

	if err != nil {
		log.Fatal(err)
	}
}
