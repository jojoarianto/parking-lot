package models

// Parking is collection of parking slot
type Parking struct {
	ParkingSlots []ParkingSlot
}

// ParkingSlot is object for parking space
type ParkingSlot struct {
	SlotID    int64
	NameSlot  string
	ParkedCar Car
}
