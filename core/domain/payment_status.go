package domain

import "errors"

var ErrInvalidStatus = errors.New("invalid Payments")

type PaymentStatus string

const (
	STARTED   PaymentStatus = "Started"
	COMPLETED PaymentStatus = "Completed"
)

func (p PaymentStatus) fromString() error {
	switch p {
	case STARTED, COMPLETED:
		return nil
	default:
		return ErrInvalidStatus
	}
}
