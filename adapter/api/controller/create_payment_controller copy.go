package controller

import (
	"net/http"
	"payments-go/adapter/api/handler"
	"payments-go/adapter/api/response"
	"payments-go/core/usecase"
	"payments-go/core/usecase/input"
)

type( 
		GetPaymentController interface {
			Execute(w http.ResponseWriter, r *http.Request)
	}
		getPaymentController struct {
			uc usecase.GetPaymentUseCase
	}
)

func NewGetPaymentController(uc usecase.GetPaymentUseCase) GetPaymentController{
	return &getPaymentController{
		uc: uc,
	}
}

func (gp getPaymentController) Execute(w http.ResponseWriter, r *http.Request){
	var paymentId = r.URL.Query().Get("paymentId")
	var i = input.GetPaymentInput{Id: paymentId}

	output, err := gp.uc.Execute(r.Context(),i)
	if err!= nil{
		handler.HandleErros(w,err)
		return 
	}
	response.NewSuccess(output,http.StatusOK).Send(w)
}