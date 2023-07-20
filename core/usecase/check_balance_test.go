package usecase

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"payments-go/core/domain"
	"payments-go/core/gateway"
	"payments-go/core/usecase/input"
	"payments-go/infrastructure/logger"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCheckBalanceUseCase(t *testing.T) {
	tt := []TableTest{
		{
			name: "Check balance with success",
			input: input.CreatePaymentInput{
				Id:          "1",
				Value:       100,
				PaymentDate: time.Now(),
				Status:      "Completed",
			},
			output: nil,
			mockedHttpResponse: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"balance": 100.0}`))),
				StatusCode: 200,
			},
			expectedError: nil,
		},
		{
			name: "Check balance with insuficient value",
			input: input.CreatePaymentInput{
				Id:          "1",
				Value:       100.0,
				PaymentDate: time.Now(),
				Status:      "Completed",
			},
			output: nil,
			mockedHttpResponse: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"balance": 50.0}`))),
				StatusCode: 200,
			},
			expectedError: domain.ErrInsuficentBalance,
		},
	}
	logger.NewZapLogger()
	for _, scenario := range tt {
		t.Run(scenario.name, func(t *testing.T) {
			fmt.Println("Scenario:", scenario.name)
			hcMock := &HtttpClientMock{}
			hcMock.On("Do", mock.Anything).Return(gateway.Response{Resp: *scenario.mockedHttpResponse, Err: nil})
			uc := NewCheckBalanceUseCase(hcMock)
			err := uc.Execute(context.TODO(), scenario.input.(input.CreatePaymentInput))
			hcMock.AssertNumberOfCalls(t, "Do", 1)
			assert.Equal(t, scenario.expectedError, err)
		})
	}

}
