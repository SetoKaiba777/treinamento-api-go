package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"payments-go/core/domain"
	"payments-go/core/gateway"
	"payments-go/core/usecase/input"
)

const (
	targetPush = "http://localhost:8882/v1/push/notification"
	methodPush = http.MethodPost
)

type(
	SendNotificationUseCase interface{
		Execute(ctx context.Context, i input.CreatePaymentInput) error
	}
	sendNotificationUseCase struct{
		httpClient gateway.HttpClient
	}
)

func NewSendNotificationUseCase(httpClient gateway.HttpClient) SendNotificationUseCase{
	return &sendNotificationUseCase{httpClient: httpClient}
}

func (sn * sendNotificationUseCase) Execute(ctx context.Context, i input.CreatePaymentInput) error{

	payment, err := domain.NewPayment().
						   WithId(i.Id).
						   WithPaymentDate(i.PaymentDate).
						   WithStatus(i.Status).
						   WithValue(i.Value).
						   Build()
	if err != nil{
		return err
	}

	jsonBody, err := json.Marshal(payment)
    if err != nil{
		return err
	}

	resp := sn.httpClient.Do(methodPush,targetPush, nil, bytes.NewBuffer(jsonBody))
	if resp.Err != nil{
		return resp.Err
	}

	return nil
}