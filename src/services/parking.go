package services

import (
	"errors"
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

// StatusParkingSpace is to return string of view state parking lot
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

// CarParking is method to parking a car
func (p *ParkingSpace) CarParking(car models.Car) (models.Parking, string, error) {

	var newParkingSlot []models.ParkingSlot
	var parkAtSlot string
	isFoundEmptySlot := false

	// no parking lot create
	if len(p.Parking.ParkingSlots) < 1 {
		return p.Parking, parkAtSlot, errors.New("Sorry, parking lot is unavailable")
	}

	// searching empty slot && filtering using copy slice
	for _, slot := range p.Parking.ParkingSlots {

		if !isFoundEmptySlot && slot.ParkedCar.CarID == 0 {

			// empty slot is found
			isFoundEmptySlot = true
			slot.ParkedCar = car
			parkAtSlot = slot.NameSlot
		}

		newParkingSlot = append(newParkingSlot, slot)
	}

	// handling not found slot
	if !isFoundEmptySlot {
		return p.Parking, parkAtSlot, errors.New("Sorry, parking lot is full")
	}

	p.Parking.ParkingSlots = newParkingSlot
	return p.Parking, parkAtSlot, nil
}

// CarLeaveParking is method for car leave parking area
func (p *ParkingSpace) CarLeaveParking(slotNo int) (models.Parking, error) {

	var newParkingSlot []models.ParkingSlot

	// searching slot_no && filtering using copy slice
	for _, slot := range p.Parking.ParkingSlots {

		if slot.SlotID == int64(slotNo) {

			// re init slot
			slot = models.ParkingSlot{
				SlotID:    slot.SlotID,
				NameSlot:  slot.NameSlot,
				ParkedCar: models.Car{},
			}
		}

		newParkingSlot = append(newParkingSlot, slot)
	}

	p.Parking.ParkingSlots = newParkingSlot
	return p.Parking, nil
}
