package usecase

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"payments-go/core/gateway"
	"payments-go/core/usecase/input"
	"payments-go/infrastructure/logger"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestCheckDebitBalanceUseCase(t *testing.T) {
	tt := []DebitTableTest{
		{
			name: "Check balance with success",
			input: input.CreatePaymentInput{
				Id:          "1",
				Value:       100,
				PaymentDate: time.Now(),
				Status:      "Completed",
			},
			output: nil,
			firstMockedHttpResponse: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"balance": 100.0}`))),
				StatusCode: 200,
			},
			secondMockedHttpResponse: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"balance": 0.0}`))),
				StatusCode: 200,
			},
			expectedError: nil,
		},
		{
			name: "Check balance with connection error",
			input: input.CreatePaymentInput{
				Id:          "1",
				Value:       100.0,
				PaymentDate: time.Now(),
				Status:      "Completed",
			},
			output: nil,
			firstMockedHttpResponse: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"balance": 110.0}`))),
				StatusCode: 200,
			},
			secondMockedHttpResponse: &http.Response{
				StatusCode: 500,
			},
			expectedError: errors.New("Generic client error"),
		},
	}
	logger.NewZapLogger()
	for _, scenario := range tt {
		t.Run(scenario.name, func(t *testing.T) {
			fmt.Println("Scenario:", scenario.name)
			hcMock := &HtttpClientMock{}
			hcMock.On("Do", mock.Anything).Return(gateway.Response{Resp: *scenario.firstMockedHttpResponse, Err: nil}).Once()
			hcMock.On("Do", mock.Anything).Return(gateway.Response{Resp: *scenario.secondMockedHttpResponse, Err: scenario.expectedError}).Once()
			uc := NewCheckDebitBalanceUseCase(hcMock)
			err := uc.Execute(context.TODO(), scenario.input.(input.CreatePaymentInput))
			hcMock.AssertNumberOfCalls(t, "Do", 2)
			assert.Equal(t, scenario.expectedError, err)
		})
	}

}
