package domain

import (
	"errors"
	"time"
)

var (
	ErrValueZero         = errors.New("value can't be zero or negative")
	ErrDateBeforeToday   = errors.New("payment date must be today")
	ErrInsuficentBalance = errors.New("insuficient balance")
	ErrFraud             = errors.New("fraudulent transaction")
)

type Payment struct {
	Id          string
	Value       float64
	PaymentDate time.Time
	Status      string
}

func NewPayment() *Payment {
	return &Payment{}
}

func (p *Payment) WithId(id string) *Payment {
	p.Id = id
	return p
}

func (p *Payment) WithValue(value float64) *Payment {
	p.Value = value
	return p
}

func (p *Payment) WithPaymentDate(paymentDate time.Time) *Payment {
	p.PaymentDate = paymentDate
	return p
}

func (p *Payment) WithStatus(status string) *Payment {
	p.Status = status
	return p
}

func (p *Payment) Build() (*Payment, error) {
	err := p.validation()
	if err != nil {
		return &Payment{}, err
	}
	return p, nil
}

func (p *Payment) validation() error {
	if p.Value <= 0 {
		return ErrValueZero
	}
	today := time.Now().Truncate(24 * time.Hour)
	if p.PaymentDate.Before(today) {
		return ErrDateBeforeToday
	}
	if err:= PaymentStatus(p.Status).fromString(); err != nil{
		return ErrInvalidStatus
	}
	return nil
}

func (p *Payment) CheckFraud(isFraud bool) error {
	if isFraud {
		return ErrFraud
	}
	return nil
}

func (p *Payment) CheckBalance(balance float64) error {
	if p.Value > balance {
		return ErrInsuficentBalance
	}
	return nil
}
