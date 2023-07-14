package facade

import (
	"context"
	"payments-go/core/usecase"
	"payments-go/core/usecase/input"
	"payments-go/infrastructure/logger"

	"golang.org/x/sync/errgroup"
)

type (
	CreatePaymentFacade struct {
		checkBalanceUseCase usecase.CheckBalanceUseCase
		checkFraudUseCase   usecase.CheckFraudUseCase
		debitBalanceUseCase usecase.CheckDebitBalanceUseCase
		savePaymentsUseCase usecase.SavePaymentsUseCase
		sendNotificationUseCase usecase.SendNotificationUseCase
	}

	CreatePaymentOutput struct{
		Id string `json:"id"`
		Status string `json:"status"`
	}
)

func NewFacade(checkBalanceUseCase usecase.CheckBalanceUseCase,
	checkFraudUseCase   usecase.CheckFraudUseCase,
	debitBalanceUseCase usecase.CheckDebitBalanceUseCase,
	savePaymentsUseCase usecase.SavePaymentsUseCase,
	sendNotificationUseCase usecase.SendNotificationUseCase) CreatePaymentFacade{
		return CreatePaymentFacade{
			checkBalanceUseCase: checkBalanceUseCase,
			checkFraudUseCase: checkFraudUseCase,
			debitBalanceUseCase : debitBalanceUseCase,
			savePaymentsUseCase: savePaymentsUseCase,
			sendNotificationUseCase :sendNotificationUseCase,
		}
	}

func (f CreatePaymentFacade) Execute(ctx context.Context, i input.CreatePaymentInput) (CreatePaymentOutput, error){
	logger.Infof("facade init")
	eg := &errgroup.Group{}

	eg.Go(func () error {return f.checkBalanceUseCase.Execute(ctx,i)})
	eg.Go(func () error {return f.checkFraudUseCase.Execute(ctx,i)})
	
	if err:= eg.Wait(); err !=nil{
		return CreatePaymentOutput{},err
	}

	if err := f.debitBalanceUseCase.Execute(ctx,i); err != nil{
		return CreatePaymentOutput{}, err
	}

	if err := f.savePaymentsUseCase.Execute(ctx,i); err != nil{
		return CreatePaymentOutput{}, err
	}

	if err := f.sendNotificationUseCase.Execute(ctx,i); err != nil{
		return CreatePaymentOutput{}, err
	}
	logger.Infof("facade finish")
	return CreatePaymentOutput{
		Id: i.Id,
		Status: i.Status,
	}, nil
}