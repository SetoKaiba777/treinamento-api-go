package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"payments-go/core/domain"
	"payments-go/core/gateway"
	"payments-go/core/usecase/input"
	"payments-go/infrastructure/logger"
)

const (
	targetDebitBalance = "http://localhost:8882/v1/balance/debit"
	methodDebitBalance = http.MethodPost
)

type (
	CheckDebitBalanceUseCase interface {
		Execute(ctx context.Context, i input.CreatePaymentInput) error
	}
	checkDebitBalanceUseCase struct {
		httpClient gateway.HttpClient
	}
)

func NewCheckDebitBalanceUseCase(htttpClient gateway.HttpClient) CheckDebitBalanceUseCase {
	return &checkDebitBalanceUseCase{httpClient: htttpClient}
}

func (c *checkDebitBalanceUseCase) Execute(ctx context.Context, i input.CreatePaymentInput) error {
	logger.WithFields(logger.Fields{"id": i.Id,"value": i.Value,
	"status": i.Status,"paymentDate": i.PaymentDate}).
	Infof("check debit balance usecase init")
	resp := c.httpClient.Do(methodBalance, targetBalance, nil, new(bytes.Reader))
	if resp.Err != nil {
		return resp.Err
	}

	balance, err := domain.NewBalance(resp.Resp)
	if err != nil {
		return err
	}

	payment, err := domain.NewPayment().
		WithId(i.Id).
		WithPaymentDate(i.PaymentDate).
		WithStatus(i.Status).
		WithValue(i.Value).
		Build()
	if err != nil {
		return err
	}

	balance.DebitBalance(*payment)

	jsonBody, err := json.Marshal(balance)
	if err != nil {
		return err
	}

	resp = c.httpClient.Do(methodDebitBalance, targetDebitBalance, nil, bytes.NewBuffer(jsonBody))
	if resp.Err != nil {
		return resp.Err
	}
	logger.WithFields(logger.Fields{"id": i.Id,"value": i.Value,
	"status": i.Status,"paymentDate": i.PaymentDate}).
	Infof("check debit balance usecase finish")
	return nil
}
