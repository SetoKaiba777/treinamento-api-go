package usecase

import (
	"context"
	"payments-go/core/domain"
	"payments-go/core/repository"
	"payments-go/core/usecase/input"
)


type (
	SavePaymentsUseCase interface {
		Execute(ctx context.Context, i input.CreatePaymentInput) error
	}
	savePaymentsUseCase struct{
		repository repository.PaymentsRepository
	}
)

func NewSavePaymentsUseCase(repository repository.PaymentsRepository) SavePaymentsUseCase{
	return &savePaymentsUseCase{repository: repository}
}

func (s * savePaymentsUseCase) Execute(ctx context.Context, i input.CreatePaymentInput) error{
	payment, err := domain.NewPayment().
						   WithId(i.Id).
						   WithPaymentDate(i.PaymentDate).
						   WithStatus(i.Status).
						   WithValue(i.Value).
						   Build()
	if err != nil{
		return err
	}

	_,err = s.repository.Save(ctx,*payment)
	if err != nil{
		return err
	}

	return nil
}
