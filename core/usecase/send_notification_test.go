package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"payments-go/core/gateway"
	"payments-go/core/usecase/input"
	"payments-go/infrastructure/logger"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSendNotifyUseCase(t *testing.T) {
	tt := []TableTest{
		{
			name: "Success to send msg",
			input: DummyPayment(),
			mockedHttpResponse: &http.Response{
				StatusCode: 200,
			},
			expectedError: nil,
		},
		{
			name: "Error to send msg",
			input: DummyPayment(),
			mockedHttpResponse: &http.Response{
				StatusCode: 500,
			},
			expectedError: errors.New("generic error"),
		},
	}
	logger.NewZapLogger()
	for _, scenario := range tt {
		t.Run(scenario.name, func(t *testing.T) {
			fmt.Println("Scenario:", scenario.name)
			hcMock := &HtttpClientMock{}
			hcMock.On("Do", mock.Anything).Return(gateway.Response{Resp: *scenario.mockedHttpResponse, Err: scenario.expectedError})
			uc := NewSendNotificationUseCase(hcMock)
			err := uc.Execute(context.TODO(), scenario.input.(input.CreatePaymentInput))
			hcMock.AssertNumberOfCalls(t, "Do", 1)
			assert.Equal(t, scenario.expectedError, err)
		})
	}

}
