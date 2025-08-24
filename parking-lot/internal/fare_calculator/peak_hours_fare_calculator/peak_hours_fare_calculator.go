package peakhoursfarecalculator

import (
	farecalculator "parking-lot/internal/fare_calculator"
	parkingticket "parking-lot/internal/parking_ticket"
	"parking-lot/internal/vehicle"
)

// 50% more fare during peak hours
const (
	SMALL_VECHICLE_PEAK_FARE_MULTIPLIER  = 1.5
	MEDIUM_VECHICLE_PEAK_FARE_MULTIPLIER = 3.0
	LARGE_VECHICLE_PEAK_FARE_MULTIPLIER  = 4.5
)

type PeakHoursFareCalculator struct {
	baseRate float64
}

func NewPeakHoursFareCalculator(baseRate float64) *PeakHoursFareCalculator {
	return &PeakHoursFareCalculator{baseRate: baseRate}
}

func (p PeakHoursFareCalculator) CalculateFare(ticket parkingticket.ParkingTicket) (float64, error) {
	// Get the multipler according to vehicle type
	var multiplier float64
	switch ticket.GetVehicle().GetType() {
	case vehicle.Small:
		multiplier = SMALL_VECHICLE_PEAK_FARE_MULTIPLIER
	case vehicle.Medium:
		multiplier = MEDIUM_VECHICLE_PEAK_FARE_MULTIPLIER
	case vehicle.Large:
		multiplier = LARGE_VECHICLE_PEAK_FARE_MULTIPLIER
	default:
		return 0, farecalculator.ErrInvalidVehicleType
	}

	duration := ticket.CalculateParkingDuration()
	return p.baseRate + multiplier*duration.Hours(), nil
}
