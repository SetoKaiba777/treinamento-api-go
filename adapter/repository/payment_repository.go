package repository

import (
	"context"
	"encoding/json"
	"payments-go/core/domain"
	"payments-go/core/repository"

	"golang.org/x/sync/errgroup"
)

type PaymentsRepository struct {
	cache PaymentsRepositoryCache
	db    PaymentsRepositoryDb
}

var _ repository.PaymentsRepository = (*PaymentsRepository)(nil)

func NewPaymentRepository(cache PaymentsRepositoryCache, db PaymentsRepositoryDb) PaymentsRepository {
	return PaymentsRepository{
		cache: cache,
		db:    db,
	}
}

func (r PaymentsRepository) Save(ctx context.Context, p domain.Payment) (domain.Payment, error){
	jsonPayment, err := json.Marshal(p)
	if err != nil{
		return domain.Payment{}, nil
	}

	eg := &errgroup.Group{}

	eg.Go(func() error{
		_,err := r.db.PutItem(ctx, p)
		return err
	})

	eg.Go(func() error{
		return r.cache.SaveCache(ctx, p.Id,string(jsonPayment))
	})

	if err=eg.Wait(); err !=nil{
		return domain.Payment{},err
	}
	return p, nil
}

func (r PaymentsRepository) FindById(ctx context.Context, id string) (domain.Payment, error){
	p,err := r.cache.FindByKeyCache(ctx,id)
	if err != nil || p==""{
		p, err := r.db.FindById(ctx,id)
		if err != nil{
			return domain.Payment{}, err
		}
		return p, nil
	}

	var out domain.Payment
	err = json.Unmarshal([]byte(p), &out)
	if err != nil{
		return domain.Payment{}, err
	}
	return out, err
}