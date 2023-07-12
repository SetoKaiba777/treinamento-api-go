package domain

import (
	"encoding/json"
	"io"
	"net/http"
)

type Fraud struct {
	IsFraud bool `json:"isFraud"`
}

func NewFraud(httpResponse http.Response) (Fraud, error) {
	f, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return Fraud{}, nil
	}
	var fraud Fraud
	err = json.Unmarshal(f, &fraud)
	if err != nil {
		return Fraud{}, nil
	}
	return fraud, nil
}
