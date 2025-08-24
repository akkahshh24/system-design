package parkingspot

import "errors"

var (
	ErrSpotOccupied         = errors.New("parking spot is already occupied")
	ErrInvalidVehicleType   = errors.New("invalid vehicle type for this parking spot")
	ErrSpotAlreadyAvailable = errors.New("parking spot is already available")
)
