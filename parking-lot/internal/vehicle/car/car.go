package car

import "parking-lot/internal/vehicle"

// Car implements the IVechile interface
// Car struct represents a car vehicle
type Car struct {
	licensePlate string // unexported field to prevent direct modification
}

func NewCar(licensePlate string) *Car {
	return &Car{licensePlate: licensePlate}
}

func (c *Car) GetType() vehicle.Type {
	return vehicle.Medium
}

func (c *Car) GetLicensePlate() string {
	return c.licensePlate
}
