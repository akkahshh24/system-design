package farecalculator

import parkingticket "parking-lot/internal/parking_ticket"

type IFareStrategy interface {
	CalculateFare(ticket *parkingticket.ParkingTicket) (float64, error)
}
