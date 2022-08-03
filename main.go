package main

import (
	"encoding/json"
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
	//	... запит даних від користувача
	//result, err := FindTrains(departureStation, arrivalStation, criteria))
	//	... обробка помилки
	//	... друк result
}

func FindTrains(departureStation, arrivalStation, criteria string) (Trains, error) {
	// ... код
	return nil, nil // маєте повернути правильні значення
}
