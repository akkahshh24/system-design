package bus

import "parking-lot/internal/vehicle"

// Bus implements the IVechile interface
// Bus struct represents a bus vehicle
type Bus struct {
	licensePlate string // unexported field to prevent direct modification
}

func NewBus(licensePlate string) *Bus {
	return &Bus{licensePlate: licensePlate}
}

func (b *Bus) GetType() vehicle.Type {
	return vehicle.Large
}

func (b *Bus) GetLicensePlate() string {
	return b.licensePlate
}
