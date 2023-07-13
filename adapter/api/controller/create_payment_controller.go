package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"payments-go/adapter/api/handler"
	"payments-go/adapter/api/response"
	"payments-go/adapter/facade"
	"payments-go/core/usecase/input"
)

type CreatePaymentController struct {
	f facade.CreatePaymentFacade
}

func NewCreatePaymentController(f facade.CreatePaymentFacade) CreatePaymentController{
	return CreatePaymentController{
		f: f,
	}
}

func (c CreatePaymentController) Execute(w http.ResponseWriter, r *http.Request){
	jsonBody, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil{
		handler.HandleErros(w,err)
	}

	var input input.CreatePaymentInput
	if err := json.Unmarshal(jsonBody,&input); err!=nil{
		handler.HandleErros(w,err)
		return
	} 
	
	output, err := c.f.Execute(r.Context(),input)
	if err!= nil{
		handler.HandleErros(w,err)
	}

	response.NewSuccess(output,http.StatusCreated).Send(w)
}