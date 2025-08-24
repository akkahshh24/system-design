package farecalculator

import parkingticket "parking-lot/internal/parking_ticket"

// There are multiple approaches to implement Strategy Pattern
// 1. Set the strategy at the time of object creation and make it immutable
// 2. Allow changing the strategy at runtime by providing a setter method
// 3. Take a list of strategies and choose one based on some criteria
// 4. Take a list of strategies and apply all strategies sequentially to calculate the final fare.
type FareCalculator struct {
	calculationStrategy IFareStrategy
}

func NewFareCalculator(strategy IFareStrategy) *FareCalculator {
	return &FareCalculator{
		calculationStrategy: strategy,
	}
}

func (fc *FareCalculator) CalculateFare(ticket *parkingticket.ParkingTicket) (float64, error) {
	return fc.calculationStrategy.CalculateFare(ticket)
}
