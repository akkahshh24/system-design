package parkinglot

import "errors"

var (
	ErrInvalidTicket        = errors.New("invaid ticket")
	ErrVehicleAlreadyExited = errors.New("vehicle already exited")
)
