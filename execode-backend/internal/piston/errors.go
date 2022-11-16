package piston

import "fmt"

// TODO: Add a piston error type implementing the error interface

type PistonError struct {
	PistonResponseCode int
	PistonErrorMessage string
	Message            string
}

func (err PistonError) Error() string {
	return fmt.Sprintf("%s, piston returned %s with code %d",
		err.Message, err.PistonErrorMessage, err.PistonResponseCode)
}
