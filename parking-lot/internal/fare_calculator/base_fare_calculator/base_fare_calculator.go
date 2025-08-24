package basefarecalculator

import (
	farecalculator "parking-lot/internal/fare_calculator"
	parkingticket "parking-lot/internal/parking_ticket"
	"parking-lot/internal/vehicle"
)

const (
	SMALL_VECHICLE_BASE_FARE_MULTIPLIER  = 1.0
	MEDIUM_VECHICLE_BASE_FARE_MULTIPLIER = 2.0
	LARGE_VECHICLE_BASE_FARE_MULTIPLIER  = 3.0
)

type BaseFareCalculator struct {
	baseRate float64
}

func NewBaseFareCalculator(baseRate float64) *BaseFareCalculator {
	return &BaseFareCalculator{baseRate: baseRate}
}

func (b BaseFareCalculator) CalculateFare(ticket *parkingticket.ParkingTicket) (float64, error) {
	// Get the multipler according to vehicle type
	var multiplier float64
	switch ticket.GetVehicle().GetType() {
	case vehicle.Small:
		multiplier = SMALL_VECHICLE_BASE_FARE_MULTIPLIER
	case vehicle.Medium:
		multiplier = MEDIUM_VECHICLE_BASE_FARE_MULTIPLIER
	case vehicle.Large:
		multiplier = LARGE_VECHICLE_BASE_FARE_MULTIPLIER
	default:
		return 0, farecalculator.ErrInvalidVehicleType
	}

	duration := ticket.CalculateParkingDuration()
	return b.baseRate + multiplier*duration.Hours(), nil
}
