package handler

import (
	"net/http"
	"payments-go/adapter/api/response"
	"payments-go/core/domain"
)

func HandleErros(w http.ResponseWriter, err error){
	var status int 

	switch err{
	case domain.ErrPaymentNotFound:
		status = http.StatusNotFound
	case domain.ErrInvalidStatus:
		status = http.StatusBadRequest
	case domain.ErrDateBeforeToday, domain.ErrInsuficentBalance, domain.ErrValueZero, domain.ErrFraud:
		status = http.StatusUnprocessableEntity
	default:
		status = http.StatusInternalServerError
	}

	response.NewError(err,status).Send(w)
}