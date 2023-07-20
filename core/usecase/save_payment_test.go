package usecase

import (
	"context"
	"errors"
	"fmt"
	"payments-go/core/domain"
	"payments-go/core/usecase/input"
	"payments-go/infrastructure/logger"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSavePaymentsUseCase(t *testing.T){
	tt := []TableTest{
		{
			name : "save payments with success",
			input: DummyPayment(),
			expectedError : nil,
		},
		{
			name : "save payments with error",
			input: DummyPayment(),
			expectedError : errors.New("Error during action of save in repository"),
		},
	}
	logger.NewZapLogger()
	for _, scenario := range tt{
		t.Run(scenario.name, func (t* testing.T)  {
			fmt.Println("Scenario:",scenario.name)
			rMock := &PaymentRepositoryMock{}
			rMock.On("Save",mock.Anything).Return(domain.Payment{}, scenario.expectedError)
			uc := NewSavePaymentsUseCase(rMock) 
			err := uc.Execute(context.TODO(), scenario.input.(input.CreatePaymentInput))
			assert.Equal(t,scenario.expectedError,err)
		})
	}
}