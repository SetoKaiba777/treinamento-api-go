package repository

import (
	"context"
	"payments-go/core/domain"
)

type PaymentsRepository interface {
	Save(context.Context, domain.Payment) (domain.Payment, error)
	FindById(context.Context, string) (domain.Payment, error)
}