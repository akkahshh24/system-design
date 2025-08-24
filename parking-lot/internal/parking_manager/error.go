package parkingmanager

import "errors"

var (
	ErrNoAvailableSpots = errors.New("no available spots for this vehicle type")
	ErrVehicleNotParked = errors.New("vehicle is not parked")
)
