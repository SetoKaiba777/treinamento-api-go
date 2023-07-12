package usecase

import (
	"bytes"
	"context"
	"net/http"
	"payments-go/core/domain"
	"payments-go/core/gateway"
	"payments-go/core/usecase/input"
)

const(
	targetFraud = "http://localhost:8882/v1/fraud"
	methodFraud = http.MethodPost
)

type (
	CheckFraudUseCase interface {
		Execute(ctx context.Context, i input.CreatePaymentInput) error
	}

	checkFraudUseCase struct{
		httpClient gateway.HttpClient
	}
)

func NewCheckFraudUseCase(httpClient gateway.HttpClient) CheckFraudUseCase{
	return &checkFraudUseCase{httpClient: httpClient}
}

func (c * checkFraudUseCase) Execute(ctx context.Context, i input.CreatePaymentInput) error{
	resp := c.httpClient.Do(methodFraud,targetFraud,nil, new(bytes.Buffer))
	if resp.Err != nil{
		return resp.Err
	}

	fraud, err := domain.NewFraud(resp.Resp)
	if err !=nil{
		return err
	}

	payment, err := domain.NewPayment().
						   WithId(i.Id).
						   WithPaymentDate(i.PaymentDate).
						   WithStatus(i.Status).
						   WithValue(i.Value).
						   Build()
	if err != nil{
		return err
	}

	err = payment.CheckFraud(fraud.IsFraud)
	if err!= nil {
		return err
	}

	return nil
}