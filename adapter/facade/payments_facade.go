package facade

import (
	"context"
	"payments-go/core/usecase"
	"payments-go/core/usecase/input"

	"golang.org/x/sync/errgroup"
)

type (
	CreatePaymentFacade struct {
		checkBalanceUseCase usecase.CheckBalanceUseCase
		checkFraudUseCase   usecase.CheckFraudUseCase
		savePaymentsUseCase usecase.SavePaymentsUseCase
	}

	CreatePaymentOutput struct{
		Id string `json:"id"`
		Status string `json:"status"`
	}
)

func NewFacade(checkBalanceUseCase usecase.CheckBalanceUseCase,
	checkFraudUseCase usecase.CheckFraudUseCase,
	savePaymentsUseCase usecase.SavePaymentsUseCase) CreatePaymentFacade{
		return CreatePaymentFacade{
			checkBalanceUseCase: checkBalanceUseCase,
			checkFraudUseCase: checkFraudUseCase,
			savePaymentsUseCase: savePaymentsUseCase,
		}
	}

func (f CreatePaymentFacade) Execute(ctx context.Context, i input.CreatePaymentInput) (CreatePaymentOutput, error){
	eg := &errgroup.Group{}

	eg.Go(func () error {return f.checkBalanceUseCase.Execute(ctx,i)})
	eg.Go(func () error {return f.checkFraudUseCase.Execute(ctx,i)})
	
	if err:= eg.Wait(); err !=nil{
		return CreatePaymentOutput{},err
	}

	if err := f.savePaymentsUseCase.Execute(ctx,i); err != nil{
		return CreatePaymentOutput{}, err
	}
	
	return CreatePaymentOutput{
		Id: i.Id,
		Status: i.Status,
	}, nil
}