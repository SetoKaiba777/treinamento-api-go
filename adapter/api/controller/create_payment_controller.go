package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"payments-go/adapter/api/handler"
	"payments-go/adapter/api/response"
	"payments-go/adapter/facade"
	"payments-go/core/usecase/input"
	"payments-go/infrastructure/logger"
)

type CreatePaymentController struct {
	f facade.CreatePaymentFacade
}

func NewCreatePaymentController(f facade.CreatePaymentFacade) CreatePaymentController {
	return CreatePaymentController{
		f: f,
	}
}

func (c CreatePaymentController) Execute(w http.ResponseWriter, r *http.Request) {
	logger.Infof("create payment init")
	jsonBody, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handler.HandleErros(w, err)
	}

	var input input.CreatePaymentInput
	if err := json.Unmarshal(jsonBody, &input); err != nil {
		handler.HandleErros(w, err)
		return
	}
	logger.WithFields(logger.Fields{"id": input.Id, "value": input.Value,
		"status": input.Status, "paymentDate": input.PaymentDate}).Infof("builded payment from request")
	output, err := c.f.Execute(r.Context(), input)
	if err != nil {
		handler.HandleErros(w, err)
	}
	logger.WithFields(logger.Fields{"id": input.Id, "value": input.Value,
		"status": input.Status, "paymentDate": input.PaymentDate}).Infof("create payment finish")
	response.NewSuccess(output, http.StatusCreated).Send(w)
}
