package entity

import (
	"payments-go/core/domain"
	"time"
)

type PaymentEntity struct {
	Id          string `dynamodbav:"id"`
	Value       float64 `dynamodbav:"value"`
	PaymentDate time.Time `dynamodbav:"paymentDate"`
	Status string `dynamodbav:"status"`
}

func (p PaymentEntity) PaymentEntityToPayment() domain.Payment{
	return domain.Payment{
		Id: p.Id,
		Value: p.Value,
		PaymentDate: p.PaymentDate,
		Status: p.Status,
	}
}

func PaymentToPaymentEntity(p domain.Payment) PaymentEntity{
	return PaymentEntity{
		Id: p.Id,
		Value: p.Value,
		PaymentDate: p.PaymentDate,
		Status: p.Status,
	}
}