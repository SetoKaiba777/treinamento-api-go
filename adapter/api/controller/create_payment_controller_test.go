package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"payments-go/adapter/facade"
	"payments-go/core/usecase/input"
	"testing"
	"time"
)


func TestRunControllerTest(t *testing.T){
	tt := []TableTestController{
		{
			name: "Successful Transaction",
			input: input.CreatePaymentInput{
				Id:          "1",
				Value:       100,
				PaymentDate: time.Now(),
				Status:      "Completed",
			},
			facadeMock: facade.CreatePaymentOutput{
				Id: "1",
				Status: "Completed",
			},
			expectedError: nil,
		},
	}
	for _, scenario := range tt{
		t.Run(scenario.name, func (t *testing.T)  {
			jsonBody, _ := json.Marshal(scenario.input)
			request, _ := http.NewRequest("POST", "v1/payments", bytes.NewBuffer(jsonBody))
			fmt.Println(request)
			
		})
	}
}