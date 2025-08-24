package medium

import (
	parkingspot "parking-lot/internal/parking_spot"
	"parking-lot/internal/vehicle"
)

// Medium implements the IParkingSpot interface
// Medium struct represents a medium parking spot
type MediumParkingSpot struct {
	spotID  string
	vehicle vehicle.IVechile // nil if the spot is available
}

func NewMediumParkingSpot(spotID string) *MediumParkingSpot {
	return &MediumParkingSpot{
		spotID:  spotID,
		vehicle: nil, // spot is initially available
	}
}

func (m *MediumParkingSpot) GetType() vehicle.Type {
	return vehicle.Small
}

func (m *MediumParkingSpot) GetID() string {
	return m.spotID
}

func (m *MediumParkingSpot) IsAvailable() bool {
	return m.vehicle == nil
}

func (m *MediumParkingSpot) Occupy(v vehicle.IVechile) error {
	// Check if the spot is already occupied
	if !m.IsAvailable() {
		return parkingspot.ErrSpotOccupied
	}

	// Check if the vehicle type is compatible with the spot type
	if v.GetType() != vehicle.Small {
		return parkingspot.ErrInvalidVehicleType
	}

	// Assign the vehicle to the spot
	m.vehicle = v
	return nil
}

func (m *MediumParkingSpot) Vacate() error {
	// Check if the spot is already available
	if m.IsAvailable() {
		return parkingspot.ErrSpotAlreadyAvailable
	}

	// Vacate the spot
	m.vehicle = nil
	return nil
}
