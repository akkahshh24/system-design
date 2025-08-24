package large

import (
	parkingspot "parking-lot/internal/parking_spot"
	"parking-lot/internal/vehicle"
)

// Small implements the IParkingSpot interface
// Small struct represents a small parking spot
type LargeParkingSpot struct {
	spotID  string
	vehicle vehicle.IVechile // nil if the spot is available
}

func NewLargeParkingSpot(spotID string) *LargeParkingSpot {
	return &LargeParkingSpot{
		spotID:  spotID,
		vehicle: nil, // spot is initially available
	}
}

func (l *LargeParkingSpot) GetType() vehicle.Type {
	return vehicle.Small
}

func (l *LargeParkingSpot) GetID() string {
	return l.spotID
}

func (l *LargeParkingSpot) IsAvailable() bool {
	return l.vehicle == nil
}

func (l *LargeParkingSpot) Occupy(v vehicle.IVechile) error {
	// Check if the spot is already occupied
	if !l.IsAvailable() {
		return parkingspot.ErrSpotOccupied
	}

	// Check if the vehicle type is compatible with the spot type
	if v.GetType() != vehicle.Small {
		return parkingspot.ErrInvalidVehicleType
	}

	// Assign the vehicle to the spot
	l.vehicle = v
	return nil
}

func (l *LargeParkingSpot) Vacate() error {
	// Check if the spot is already available
	if l.IsAvailable() {
		return parkingspot.ErrSpotAlreadyAvailable
	}

	// Vacate the spot
	l.vehicle = nil
	return nil
}
