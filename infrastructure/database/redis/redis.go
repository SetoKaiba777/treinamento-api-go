package redis

import (
	"context"
	"payments-go/adapter/repository"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisTemplate struct {
	RedisClient *redis.Client
}

var _ repository.PaymentsRepositoryCache = (*RedisTemplate)(nil)

func New(redisHost string) *RedisTemplate{
	rdb := redis.NewClient(&redis.Options{
		Addr: redisHost,
		Password: "",
		DB: 0,
	})

	return &RedisTemplate{
		RedisClient: rdb,
	}
}

func (r RedisTemplate) FindByKeyCache(ctx context.Context, key string) (string, error){
	val, err := r.RedisClient.Get(ctx,key).Result()
	if err!= nil{
		return "", nil
	}

	return val,nil
}

func(r RedisTemplate) SaveCache(ctx context.Context, key,value string) error{
	err := r.RedisClient.Set(ctx, key, value, time.Hour*72).Err()
	if err !=nil{
		return err
	}
	return nil
}
