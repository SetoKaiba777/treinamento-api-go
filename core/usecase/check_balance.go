package usecase

import (
	"bytes"
	"context"
	"net/http"
	"payments-go/core/domain"
	"payments-go/core/gateway"
	"payments-go/core/usecase/input"
	"payments-go/infrastructure/logger"
)

const (
	targetBalance = "http://localhost:8882/v1/balance"
	methodBalance = http.MethodGet
)

type (
	CheckBalanceUseCase interface {
		Execute(ctx context.Context, i input.CreatePaymentInput) error
	}

	checkBalanceUseCase struct {
		httpClient gateway.HttpClient
	}
)

func NewCheckBalanceUseCase(httpClient gateway.HttpClient) CheckBalanceUseCase {
	return &checkBalanceUseCase{httpClient: httpClient}
}

func (c *checkBalanceUseCase) Execute(ctx context.Context, i input.CreatePaymentInput) error {
	logger.WithFields(logger.Fields{"id": i.Id,"value": i.Value,
	"status": i.Status,"paymentDate": i.PaymentDate}).
	Infof("check payment usecase init")
	
	resp := c.httpClient.Do(methodBalance, targetBalance, nil, new(bytes.Buffer))
	if resp.Err != nil {
		return resp.Err
	}

	payment, err := domain.NewPayment().
		WithId(i.Id).
		WithValue(i.Value).
		WithPaymentDate(i.PaymentDate).
		WithStatus(i.Status).
		Build()
	if err != nil {
		return err
	}

	balance, err := domain.NewBalance(resp.Resp)
	if err != nil {
		return err
	}

	err = payment.CheckBalance(balance.Balance)
	if err != nil {
		return err
	}
	logger.WithFields(logger.Fields{"id": i.Id,"value": i.Value,
	"status": i.Status,"paymentDate": i.PaymentDate}).
	Infof("check payment usecase finish")
	return nil
}
