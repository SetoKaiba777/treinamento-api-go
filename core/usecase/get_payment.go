package usecase

import (
	"context"
	"payments-go/core/repository"
	"payments-go/core/domain"
	"payments-go/core/usecase/input"
	"time"
)

type (
	GetPaymentUseCase interface {
		Execute(ctx context.Context, i input.GetPaymentInput) (GetPaymentOutput,error)
	}
	
	getPaymentUseCase struct {
		repository repository.PaymentsRepository
	}

	GetPaymentOutput struct{
		Id          string    `json:"id"`
		Value       float64   `json:"value"`
		PaymentDate time.Time `json:"paymentDate"`
		Status      string    `json:"status"`
	}
)

var _ GetPaymentUseCase = (*getPaymentUseCase)(nil)

func NewGetPaymentUseCase(repository repository.PaymentsRepository) GetPaymentUseCase{
	return &getPaymentUseCase{
		repository: repository,
	}
}

func (g  *getPaymentUseCase) Execute(ctx context.Context, i input.GetPaymentInput) (GetPaymentOutput,error){
	p, err := g.repository.FindById(ctx,i.Id)
	if err != nil{
		return GetPaymentOutput{}, err
	}

	if p.Id==""{
		return GetPaymentOutput{}, domain.ErrPaymentNotFound
	}
	
	return GetPaymentOutput{
		Id: p.Id,
		Value: p.Value,
		PaymentDate: p.PaymentDate,
		Status: p.Status,
	}, nil
}