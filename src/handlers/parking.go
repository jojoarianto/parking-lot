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
