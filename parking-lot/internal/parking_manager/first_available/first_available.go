package demo

import (
	parkingmanager "parking-lot/internal/parking_manager"
	parkingspot "parking-lot/internal/parking_spot"
	"parking-lot/internal/vehicle"
)

// DemoParkingManager is a demo implementation of the parking manager interface.
type FirstAvailableParkingManager struct {
	// Map to keep track of available spots by vehicle type
	// Provides O(1) access to available spots for each vehicle type
	availableSpots map[vehicle.Type][]parkingspot.IParkingSpot

	// Map to keep track of which vehicle is parked in which spot
	// Provides O(1) access to the parking spot for a given vehicle
	vehicleToSpotMap map[vehicle.IVechile]parkingspot.IParkingSpot

	// Map to keep track of which spot occupies which vehicle
	spotToVehicleMap map[parkingspot.IParkingSpot]vehicle.IVechile
}

func NewFirstAvailableParkingManager(availableSpots map[vehicle.Type][]parkingspot.IParkingSpot) *FirstAvailableParkingManager {
	return &FirstAvailableParkingManager{
		availableSpots:   availableSpots,
		vehicleToSpotMap: make(map[vehicle.IVechile]parkingspot.IParkingSpot),
	}
}

func (pm *FirstAvailableParkingManager) findSpotForVehicle(v vehicle.IVechile) (parkingspot.IParkingSpot, error) {
	// Get the vehicle type
	vehicleType := v.GetType()

	// Check for available spots of the required type
	spots, exists := pm.availableSpots[vehicleType]
	if !exists || len(spots) == 0 {
		return nil, parkingmanager.ErrNoAvailableSpots // No available spots for this vehicle type
	}

	// Return the first available spot
	return spots[0], nil
}

func (pm *FirstAvailableParkingManager) ParkVehicle(vehicle vehicle.IVechile) (parkingspot.IParkingSpot, error) {
	// Find an available spot for the vehicle
	spot, err := pm.findSpotForVehicle(vehicle)
	if err != nil {
		return nil, err
	}

	err = spot.Occupy(vehicle)
	if err != nil {
		return nil, err
	}

	// Map the vehicle to the spot
	pm.vehicleToSpotMap[vehicle] = spot
	pm.spotToVehicleMap[spot] = vehicle

	// Remove the spot from available spots
	vehicleType := vehicle.GetType()
	currentSpots := pm.availableSpots[vehicleType]
	pm.availableSpots[vehicleType] = currentSpots[1:]

	return spot, nil
}

func (pm *FirstAvailableParkingManager) UnparkVehicle(v vehicle.IVechile) error {
	// Check if the vehicle is parked
	spot, exists := pm.vehicleToSpotMap[v]
	if !exists {
		return parkingmanager.ErrVehicleNotParked // Vehicle is not parked
	}

	// Vacate the spot
	err := spot.Vacate()
	if err != nil {
		return err
	}

	// Remove the vehicle from the map
	delete(pm.vehicleToSpotMap, v)

	// Add the spot back to available spots
	vehicleType := v.GetType()
	pm.availableSpots[vehicleType] = append(pm.availableSpots[vehicleType], spot)

	return nil
}
