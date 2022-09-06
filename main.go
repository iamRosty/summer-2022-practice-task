// Package main implements a utility to search for suitable trains,
// selected by user criteria: departure/arrival station, departure/arrival time and fare.
// The data of the trains are in the file data.json
// Meaning of keyword criteria:'price': cheaper first, 'arrival-time': first those that arrive earlier,
// 'departure-time': first those that depart earlier.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"time"
)

const jsonFile string = "./data.json"
const layout string = "15:04:05"
const outputQty int = 3

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

	result, err := FindTrains(departureStation, arrivalStation, criteria)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range result {
		fmt.Printf("%+v\n", v)
	}
}

func FindTrains(departureStation, arrivalStation, criteria string) (Trains, error) {
	if departureStation == "" {
		return nil, errors.New("empty departure station")
	}

	if intVar, err := strconv.Atoi(departureStation); intVar <= 0 || err != nil {
		return nil, errors.New("bad departure station input")
	}

	if arrivalStation == "" {
		return nil, errors.New("empty arrival station")
	}

	if intVar, err := strconv.Atoi(arrivalStation); intVar <= 0 || err != nil {
		return nil, errors.New("bad arrival station input")
	}

	if criteria != "price" && criteria != "arrival-time" && criteria != "departure-time" {
		return nil, errors.New("unsupported criteria")
	}

	jsonData := readJson(jsonFile)

	var trainsBase Trains
	err := json.Unmarshal(jsonData, &trainsBase)
	if err != nil {
		log.Fatal(err)
	}

	var selectedTrains Trains
	for _, v := range trainsBase {
		if departureStation == strconv.Itoa(v.DepartureStationID) && arrivalStation == strconv.Itoa(v.ArrivalStationID) {
			selectedTrains = append(selectedTrains, v)
		}
	}

	if len(selectedTrains) == 0 {
		return nil, nil
	}

	sortTrains(selectedTrains, criteria)

	if len(selectedTrains) < outputQty {
		return selectedTrains, nil
	}

	return selectedTrains[:outputQty], nil
}

func sortTrains(t Trains, criteria string) {
	switch criteria {
	case "price":
		sort.SliceStable(t, func(i, j int) bool {
			return t[i].Price < t[j].Price
		})
	case "arrival-time":
		sort.SliceStable(t, func(i, j int) bool {
			return t[i].ArrivalTime.Before(t[j].ArrivalTime)
		})
	case "departure-time":
		sort.SliceStable(t, func(i, j int) bool {
			return t[i].DepartureTime.Before(t[j].DepartureTime)
		})
	}
}

func readJson(jsonFile string) []byte {
	jsonData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	if !json.Valid(jsonData) {
		log.Fatal("invalid json!")
	}
	return jsonData
}
