package parkingspot

import "parking-lot/internal/vehicle"

type IParkingSpot interface {
	GetType() vehicle.Type
	GetID() string
	IsAvailable() bool
	Occupy(vehicle vehicle.IVechile) error
	Vacate() error
}
