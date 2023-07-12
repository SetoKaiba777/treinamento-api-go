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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCheckFraudUseCase(t *testing.T) {
	tt := []TableTest{
		{
			name : "Check fraud with success",
			input: input.CreatePaymentInput{
				Id: "1",
				Value: 100,
				PaymentDate: time.Now(),
				Status: "Completed",
			},
			output: nil,
			mockedHttpResponse: &http.Response{
				Body: io.NopCloser(bytes.NewReader([]byte(`{"isFraud": false}`))),
				StatusCode: 200,
			},
			expectedError: nil,
		},
		{
			name : "Check fraudlent payment",
			input: input.CreatePaymentInput{
				Id: "1",
				Value: 100.0,
				PaymentDate: time.Now(),
				Status: "Completed",
			},
			output: nil,
			mockedHttpResponse: &http.Response{
				Body: io.NopCloser(bytes.NewReader([]byte(`{"isFraud": true}`))),
				StatusCode: 200,
			},
			expectedError: domain.ErrFraud,
		},
	}
	for _, scenario := range tt{
		t.Run(scenario.name,func(t *testing.T) {
			fmt.Println("Scenario:",scenario.name)
			hcMock := &HtttpClientMock{}
			hcMock.On("Do",mock.Anything).Return(gateway.Response{Resp: *scenario.mockedHttpResponse, Err: nil})
			uc := NewCheckFraudUseCase(hcMock)
			err := uc.Execute(context.TODO(),scenario.input.(input.CreatePaymentInput))
			hcMock.AssertNumberOfCalls(t,"Do",1)
			assert.Equal(t,scenario.expectedError,err)
		})
	}

}
