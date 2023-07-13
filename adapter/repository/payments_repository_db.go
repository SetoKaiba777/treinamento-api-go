package repository

import (
	"context"
	"payments-go/core/domain"
)

type PaymentsRepositoryDb interface {
	FindById(ctx context.Context, id string) (domain.Payment, error)
	PutItem(ctx context.Context, item domain.Payment) (domain.Payment, error)
}