package usecase

import (
	"context"
	"io"
	"net/http"
	"payments-go/core/domain"
	"payments-go/core/gateway"
	"payments-go/core/usecase/input"
	"time"

	"github.com/stretchr/testify/mock"
)

type (
	TableTest struct {
		name               string
		input              any
		output             any
		mockedHttpResponse *http.Response
		expectedError      error
	}

	DebitTableTest struct {
		name                     string
		input                    any
		output                   any
		firstMockedHttpResponse  *http.Response
		secondMockedHttpResponse *http.Response
		expectedError            error
	}

	GetTableTest struct {
		name           string
		input          any
		output         any
		mockedResponse any
		mockedError error
		expectedError  error
	}

	HtttpClientMock struct {
		mock.Mock
	}

	PaymentRepositoryMock struct {
		mock.Mock
	}
)

func (hm *HtttpClientMock) Do(method, target string, httpHeaders http.Header, httpBody io.Reader) gateway.Response {
	args := hm.Called()
	return args.Get(0).(gateway.Response)
}

func (r *PaymentRepositoryMock) Save(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	args := r.Called()
	return args.Get(0).(domain.Payment), args.Error(1)
}

func (r *PaymentRepositoryMock) FindById(ctx context.Context, id string) (domain.Payment, error) {
	args := r.Called()
	return args.Get(0).(domain.Payment), args.Error(1)
}

func DummyPayment() input.CreatePaymentInput {
	return input.CreatePaymentInput{
		Id:          "1",
		Value:       10,
		PaymentDate: time.Now(),
		Status:      "Completed",
	}
}
