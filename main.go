package main

import (
	"fmt"
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

func main() {
	//	... запит даних від користувача
	var departureStation, arrivalStation, criteria string

	fmt.Print("Type the station of departure and press enter: ")
	fmt.Scan(&departureStation)
	fmt.Print("Type the station of arrival and press enter: ")
	fmt.Scan(&arrivalStation)
	fmt.Print("Type the criteria for selecting trains that suit you.\n Use ONE of these keywords:'price', 'arrival-time' and  'departure-time'.\n Price- cheaper first, arrival-time - first those that arrive earlier, departure-time - first those that depart earlier: ")
	fmt.Scan(&criteria)
	//result, err := FindTrains(departureStation, arrivalStation, criteria))
	//	... обробка помилки
	//	... друк result
}

func FindTrains(departureStation, arrivalStation, criteria string) (Trains, error) {
	// ... код
	return nil, nil // маєте повернути правильні значення
}
