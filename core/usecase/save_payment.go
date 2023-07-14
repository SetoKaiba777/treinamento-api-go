package usecase

import (
	"context"
	"payments-go/core/domain"
	"payments-go/core/repository"
	"payments-go/core/usecase/input"
	"payments-go/infrastructure/logger"
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
	logger.WithFields(logger.Fields{"id": i.Id,"value": i.Value,
	"status": i.Status,"paymentDate": i.PaymentDate}).
	Infof("save payment usecase init")
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
	logger.WithFields(logger.Fields{"id": i.Id,"value": i.Value,
	"status": i.Status,"paymentDate": i.PaymentDate}).
	Infof("save payment usecase finish")
	return nil
}
