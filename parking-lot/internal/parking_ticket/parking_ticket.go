package parkingticket

import (
	parkingspot "parking-lot/internal/parking_spot"
	"parking-lot/internal/vehicle"
	"time"
)

type ParkingTicket struct {
	ticketID    string
	vehicle     vehicle.IVechile
	parkingSpot parkingspot.IParkingSpot
	entryTime   time.Time
	exitTime    time.Time
	paid        bool
}

func NewParkingTicket(ticketID string, vehicle vehicle.IVechile, spot parkingspot.IParkingSpot, entryTime time.Time) *ParkingTicket {
	return &ParkingTicket{
		ticketID:    ticketID,
		vehicle:     vehicle,
		parkingSpot: spot,
		entryTime:   entryTime,
		exitTime:    time.Time{},
		paid:        false,
	}
}

func (pt *ParkingTicket) GetTicketID() string {
	return pt.ticketID
}

func (pt *ParkingTicket) GetVehicle() vehicle.IVechile {
	return pt.vehicle
}

func (pt *ParkingTicket) GetParkingSpot() parkingspot.IParkingSpot {
	return pt.parkingSpot
}

func (pt *ParkingTicket) GetEntryTime() time.Time {
	return pt.entryTime
}

func (pt *ParkingTicket) GetExitTime() time.Time {
	return pt.exitTime
}

func (pt *ParkingTicket) SetExitTime(exitTime time.Time) {
	pt.exitTime = exitTime
}

func (pt *ParkingTicket) IsPaid() bool {
	return pt.paid
}

func (pt *ParkingTicket) MarkPaid() {
	pt.paid = true
}

func (pt *ParkingTicket) MarkExit(exitTime time.Time) {
	pt.exitTime = exitTime
}

func (pt *ParkingTicket) CalculateParkingDuration() time.Duration {
	if pt.exitTime.IsZero() {
		return time.Since(pt.entryTime)
	}

	return pt.exitTime.Sub(pt.entryTime)
}
