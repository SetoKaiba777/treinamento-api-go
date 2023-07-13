package usecase

import (
	"context"
	"errors"
	"payments-go/core/domain"
	"payments-go/core/usecase/input"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPayment(t *testing.T) {
	tt := []GetTableTest{
		{
			name: "Get successful",
			input: input.GetPaymentInput{
				Id: "1",
			},
			output: GetPaymentOutput{
				Id:          "1",
				Value:       10,
				PaymentDate: time.Now(),
				Status:      "Completed",
			},
			mockedResponse:domain.Payment{
				Id:          "1",
				Value:       10,
				PaymentDate: time.Now(),
				Status:      "Completed",
			},
			mockedError: nil,
			expectedError: nil,
		},
		{
			name: "Not found payment",
			input: input.GetPaymentInput{
				Id: "1",
			},
			output: GetPaymentOutput{},
			mockedResponse:domain.Payment{},
			mockedError: nil,
			expectedError: domain.ErrPaymentNotFound,
		},
		{
			name: "Repository error",
			input: input.GetPaymentInput{
				Id: "1",
			},
			output: GetPaymentOutput{},
			mockedError: errors.New("repository error"),
			mockedResponse:domain.Payment{},
			expectedError: errors.New("repository error"),
		},
	}
	for _, scenario := range tt{
		t.Run(scenario.name,func(t *testing.T) {
			rMock := PaymentRepositoryMock{}
			rMock.On("FindById",mock.Anything).Return(scenario.mockedResponse.(domain.Payment),scenario.mockedError)
			uc := NewGetPaymentUseCase(&rMock)
			res, err := uc.Execute(context.TODO(),scenario.input.(input.GetPaymentInput))
			assert.Equal(t, scenario.expectedError, err)
			assert.Equal(t,scenario.output,res)
		})
	}
}