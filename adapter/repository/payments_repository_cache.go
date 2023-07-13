package repository

import "context"

type PaymentsRepositoryCache interface {
	FindByKeyCache(ctx context.Context, key string) (string, error)
	SaveCache(ctx context.Context, key,value string) error
}