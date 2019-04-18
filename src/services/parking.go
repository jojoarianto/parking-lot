package services

import (
	"strconv"

	"github.com/jojoarianto/parking-lot/src/models"
)

// ParkingSpace is object of parking
type ParkingSpace struct {
	Parking models.Parking
}

// NewParkingService constructor for service
func NewParkingService(p models.Parking) *ParkingSpace {
	return &ParkingSpace{
		Parking: p,
	}
}

// CreateParkingSpace is method to create parking space
func (p *ParkingSpace) CreateParkingSpace(nSlot int) (models.Parking, error) {

	var parkingSlot []models.ParkingSlot

	// init parking space with n slot
	for id := 1; id <= nSlot; id++ {
		name := strconv.Itoa(id)
		newParkingSlot := models.ParkingSlot{
			SlotID:    int64(id),
			NameSlot:  name,
			ParkedCar: models.Car{},
		}
		parkingSlot = append(parkingSlot, newParkingSlot)
	}

	p.Parking.ParkingSlots = parkingSlot
	return p.Parking, nil
}
