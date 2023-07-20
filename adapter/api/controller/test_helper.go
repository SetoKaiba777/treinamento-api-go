package controller

import (
	"context"
	"payments-go/adapter/facade"
	"payments-go/core/usecase/input"

	"github.com/stretchr/testify/mock"
)

type (
	TableTestController struct {
		name          string
		input         any
		facadeMock    facade.CreatePaymentOutput
		expectedStatus int
		expectedError error
	}

	FacadeMock struct{
		mock.Mock
	}
)

var _ facade.PaymentFacade = (*FacadeMock)(nil)

func (fm * FacadeMock) Execute(ctx context.Context, i input.CreatePaymentInput) (facade.CreatePaymentOutput, error){
	args := fm.Called()
	return args.Get(0).(facade.CreatePaymentOutput) , args.Error(1)
}



