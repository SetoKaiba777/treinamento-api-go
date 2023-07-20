package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"payments-go/adapter/facade"
	"payments-go/core/domain"
	"payments-go/core/usecase/input"
	"payments-go/infrastructure/logger"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
			expectedStatus: http.StatusCreated,
			expectedError: nil,
		},
		{
			name: "Invalid Transaction",
			input: input.CreatePaymentInput{
				Id:          "1",
				Value:       100,
				PaymentDate: time.Now(),
				Status:      "Batata",
			},
			facadeMock: facade.CreatePaymentOutput{
				Id: "",
				Status: "",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError: domain.ErrInvalidStatus,
		},
	}
	logger.NewZapLogger()
	for _, scenario := range tt{
		t.Run(scenario.name, func (t *testing.T)  {
			jsonBody, _ := json.Marshal(scenario.input)

			request, _ := http.NewRequest("POST", "v1/payments", bytes.NewBuffer(jsonBody))
			fmt.Println(request)
			responseRecorder := httptest.NewRecorder()
			
			mockFacade := &FacadeMock{}
			mockFacade.On("Execute", mock.Anything).Return(scenario.facadeMock,scenario.expectedError)
			
			controller := NewCreatePaymentController(mockFacade)
			controller.Execute(responseRecorder,request)
			
			assert.Equal(t, scenario.expectedStatus, responseRecorder.Code)

		})
	}
}