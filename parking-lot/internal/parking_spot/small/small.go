package small

import (
	parkingspot "parking-lot/internal/parking_spot"
	"parking-lot/internal/vehicle"
)

// Small implements the IParkingSpot interface
// Small struct represents a small parking spot
type SmallParkingSpot struct {
	spotID  string
	vehicle vehicle.IVechile // nil if the spot is available
}

func NewSmallParkingSpot(spotID string) *SmallParkingSpot {
	return &SmallParkingSpot{
		spotID:  spotID,
		vehicle: nil, // spot is initially available
	}
}

func (s *SmallParkingSpot) GetType() vehicle.Type {
	return vehicle.Small
}

func (s *SmallParkingSpot) GetID() string {
	return s.spotID
}

func (s *SmallParkingSpot) IsAvailable() bool {
	return s.vehicle == nil
}

func (s *SmallParkingSpot) Occupy(v vehicle.IVechile) error {
	// Check if the spot is already occupied
	if !s.IsAvailable() {
		return parkingspot.ErrSpotOccupied
	}

	// Check if the vehicle type is compatible with the spot type
	if v.GetType() != vehicle.Small {
		return parkingspot.ErrInvalidVehicleType
	}

	// Assign the vehicle to the spot
	s.vehicle = v
	return nil
}

func (s *SmallParkingSpot) Vacate() error {
	// Check if the spot is already available
	if s.IsAvailable() {
		return parkingspot.ErrSpotAlreadyAvailable
	}

	// Vacate the spot
	s.vehicle = nil
	return nil
}
