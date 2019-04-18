package handlers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jojoarianto/parking-lot/src/models"
	"github.com/jojoarianto/parking-lot/src/services"
)

var parkingSpace models.Parking
var carCounter int

// CreateParkingHandler is method to handler input
// Parameter Example =>> "create_parking_lot 6"
func CreateParkingHandler(arrCommandStr []string) error {

	if len(arrCommandStr) < 2 {
		return errors.New("At least 2 arguments (create_parking_lot n_slot)")
	}

	n, err := strconv.ParseInt(arrCommandStr[1], 10, 64)
	if err != nil {
		return errors.New("Parameter must be numeric")
	}

	// init parking service
	parkingService := services.NewParkingService(parkingSpace)

	// call create parking parkingSpace with n slot
	parkingSpace, _ = parkingService.CreateParkingSpace(int(n))
	msg := fmt.Sprintf("Created a parking lot with %d slots", n)
	fmt.Println(msg)

	return nil
}

// StatusParkingSlotHandler is method to handler input
// Parameter Example =>> "status"
func StatusParkingSlotHandler() error {

	// init parking service
	parkingService := services.NewParkingService(parkingSpace)

	// call status parking (return in string)
	resp, _ := parkingService.StatusParkingSpace()
	fmt.Println(resp)

	return nil
}

// CarParkingHandler is method for park a car
// Parameter Example =>> "park KA-01-HH-3141 Black"
func CarParkingHandler(arrCommandStr []string) error {

	if len(arrCommandStr) < 3 {
		return errors.New("At least 3 arguments (park car_platno car_colour)")
	}

	// init parking service
	parkingService := services.NewParkingService(parkingSpace)

	carCounter++

	// create a car
	newcar := models.Car{
		CarID:  int64(carCounter),
		PlatNo: arrCommandStr[1],
		Colour: arrCommandStr[2],
	}

	var err error
	var parkAtSlot string

	// call car parking method with car param
	parkingSpace, parkAtSlot, err = parkingService.CarParking(newcar)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("Allocated slot number: %s", parkAtSlot)
	fmt.Println(msg)

	return nil
}
