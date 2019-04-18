package services

import (
	"fmt"
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

// StatusParkingSpace is method to create parking space
func (p *ParkingSpace) StatusParkingSpace() (string, error) {

	// print header
	str := "Slot No.    Registration No    Colour\n"

	// fill with status per parking slot
	for _, slot := range p.Parking.ParkingSlots {

		if slot.ParkedCar.CarID == 0 {
			// skip empty slot
			continue
		}

		// Adjustment colomn slot no with max 10 char
		col1 := fmt.Sprintf("%d                            ", slot.SlotID)
		col1 = col1[:12]

		// Adjustment colomn plat no with max 20 char
		col2 := fmt.Sprintf("%s                            ", slot.ParkedCar.PlatNo)
		col2 = col2[:19]

		// Adjustment colomn car colour with max 10 char
		col3 := fmt.Sprintf("%s", slot.ParkedCar.Colour)
		// col3 = col3[:15]

		str += fmt.Sprintf("%s%s%s\n",
			col1,
			col2,
			col3,
		)
	}

	return str, nil
}
