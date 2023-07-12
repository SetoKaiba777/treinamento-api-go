package domain

import (
	"encoding/json"
	"io"
	"net/http"
)

type Balance struct {
	Balance float64 `json:"balance"`
}

func NewBalance(httpResponse http.Response) (Balance, error) {
	b, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return Balance{}, err
	}
	var balance Balance
	err = json.Unmarshal(b, &balance)
	if err != nil {
		return Balance{}, err
	}
	return balance, nil
}

func (b* Balance) DebitBalance(payment Payment){
	b.Balance = b.Balance - payment.Value
}