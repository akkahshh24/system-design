package parkingmanager

import (
	parkingspot "parking-lot/internal/parking_spot"
	"parking-lot/internal/vehicle"
)

type IParkingManager interface {
	ParkVehicle(v vehicle.IVechile) (parkingspot.IParkingSpot, error)
	UnparkVehicle(v vehicle.IVechile) error
}
