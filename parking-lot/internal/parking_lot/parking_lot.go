package parkinglot

import (
	"log"
	farecalculator "parking-lot/internal/fare_calculator"
	parkingmanager "parking-lot/internal/parking_manager"
	parkingticket "parking-lot/internal/parking_ticket"
	"parking-lot/internal/vehicle"
	"time"
)

// Acts as a facade and provides a simple interface for clients.
// Delegates the spot allocation logic to ParkingManager and pricing logic to FareCalculator.
type ParkingLot struct {
	parkingManager parkingmanager.IParkingManager
	fareCalculator farecalculator.IFareStrategy
}

func NewParkingLot(parkingManager parkingmanager.IParkingManager, fareCalculator farecalculator.IFareStrategy) *ParkingLot {
	return &ParkingLot{
		parkingManager: parkingManager,
		fareCalculator: fareCalculator,
	}
}

func (pl *ParkingLot) ParkVehicle(vehicle vehicle.IVechile) (*parkingticket.ParkingTicket, error) {
	var ticket *parkingticket.ParkingTicket

	// Park the vehicle using ParkingManager
	spot, err := pl.parkingManager.ParkVehicle(vehicle)
	if err != nil {
		return ticket, err
	}

	// Generate a new ticket and return to the client
	ticket = parkingticket.NewParkingTicket("ticket-id", vehicle, spot, time.Now())
	return ticket, nil
}

func (pl *ParkingLot) UnparkVehicle(ticket *parkingticket.ParkingTicket) error {
	// Initial sanity checks
	if ticket == nil {
		return ErrInvalidTicket
	}

	if !ticket.GetExitTime().IsZero() {
		return ErrVehicleAlreadyExited
	}

	// Set the exit time
	ticket.SetExitTime(time.Now())

	// Unpark the vehicle using ParkingManager
	err := pl.parkingManager.UnparkVehicle(ticket.GetVehicle())
	if err != nil {
		return err
	}

	// Calculate the fare
	fare, err := pl.fareCalculator.CalculateFare(ticket)
	if err != nil {
		return err
	}

	log.Printf("Vehicle: %s has left. Total fare: %.2f", ticket.GetVehicle().GetLicensePlate(), fare)
	return nil
}
