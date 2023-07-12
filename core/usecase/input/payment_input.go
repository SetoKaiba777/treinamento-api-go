package input

import "time"

type CreatePaymentInput struct {
	Id          string    `json:"id"`
	Value       float64   `json:"value"`
	PaymentDate time.Time `json:"paymentDate"`
	Status      string    `json:"status"`
}
