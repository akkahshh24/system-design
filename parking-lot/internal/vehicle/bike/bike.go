package bike

import "parking-lot/internal/vehicle"

// Bike implements the IVechile interface
// Bike struct represents a bike vehicle
type Bike struct {
	licensePlate string // unexported field to prevent direct modification
}

func NewBike(licensePlate string) *Bike {
	return &Bike{licensePlate: licensePlate}
}

func (b *Bike) GetType() vehicle.Type {
	return vehicle.Small
}

func (b *Bike) GetLicensePlate() string {
	return b.licensePlate
}
