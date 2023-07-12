package domain

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testTable struct {
	name          string
	input         Payment
	balance       float64
	isFraud       bool
	output        Payment
	expectedError error
}

func TestPayment(t *testing.T) {
	paymentSuccess := Payment{
		Id:          "1",
		Value:       100,
		PaymentDate: time.Now(),
		Status:      "Completed",
	}
	tt := []testTable{
		{
			name:          "Success Build",
			input:         paymentSuccess,
			output:        paymentSuccess,
			expectedError: nil,
		},
		{
			name:          "Error value less than zero",
			input:         Payment{
				Id:          "1",
				Value:       0,
				PaymentDate: time.Now(),
				Status:      "Completed",
			},
			expectedError: ErrValueZero,
		},
		{
			name:          "Error past date",
			input:         Payment{
				Id:          "1",
				Value:       10,
				PaymentDate: time.Now().Truncate(24*time.Hour).Add(-1*time.Minute),
				Status:      "Completed",
			},
			expectedError: ErrDateBeforeToday,
		},
		{
			name:          "Error invalid status",
			input:         Payment{
				Id:          "1",
				Value:       10,
				PaymentDate: time.Now(),
				Status:      "Batata",
			},
			expectedError: ErrInvalidStatus,
		},
	}
	for _, scenario := range tt {
		fmt.Println("Scenario:", scenario.name)
		p, err := NewPayment().
			WithId(scenario.input.Id).
			WithValue(scenario.input.Value).
			WithStatus(scenario.input.Status).
			WithPaymentDate(scenario.input.PaymentDate).
			Build()

		assert.Equal(t, scenario.expectedError, err)

		if err == nil {
			assert.Equal(t, scenario.output.Id, p.Id)
			assert.Equal(t, scenario.output.PaymentDate, p.PaymentDate)
			assert.Equal(t, scenario.output.Value, p.Value)
			assert.Equal(t, scenario.output.Status, p.Status)
		}
	}
}

func TestBalance(t *testing.T){
	tt := []testTable{
		{
			name: "Suficient balance",
			input: Payment{
				Id: "1",
				Value: 10,
				PaymentDate: time.Now(),
				Status: "Completed",
			},
			balance: 10,
			expectedError: nil,
		},
		{
			name: "Insuficient balance",
			input: Payment{
				Id: "1",
				Value: 100,
				PaymentDate: time.Now(),
				Status: "Completed",
			},
			balance: 10,
			expectedError: ErrInsuficentBalance,
		},
	}
	for _, scenario := range tt{
		fmt.Println("Scenario:",scenario.name)
		p, _ := NewPayment().
			WithId(scenario.input.Id).
			WithValue(scenario.input.Value).
			WithStatus(scenario.input.Status).
			WithPaymentDate(scenario.input.PaymentDate).
			Build()

		err := p.CheckBalance(scenario.balance)
		assert.Equal(t,scenario.expectedError,err)
	}
}

func TestFraud(t *testing.T){
	tt := []testTable{
		{
			name: "Regular Transaction",
			input: Payment{
				Id: "1",
				Value: 10,
				PaymentDate: time.Now(),
				Status: "Completed",
			},
			isFraud: false,
			expectedError: nil,
		},
		{
			name: "Fraudlent transaction",
			input: Payment{
				Id: "1",
				Value: 100,
				PaymentDate: time.Now(),
				Status: "Completed",
			},
			isFraud: true,
			expectedError: ErrFraud,
		},
	}
	for _, scenario := range tt{
		fmt.Println("Scenario:",scenario.name)
		p, _ := NewPayment().
			WithId(scenario.input.Id).
			WithValue(scenario.input.Value).
			WithStatus(scenario.input.Status).
			WithPaymentDate(scenario.input.PaymentDate).
			Build()

		err := p.CheckFraud(scenario.isFraud)
		assert.Equal(t,scenario.expectedError,err)
	}
}