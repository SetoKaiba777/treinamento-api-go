package server

import (
	"payments-go/core/gateway"
	coreRepository 	"payments-go/core/repository"
	appConfig "payments-go/infrastructure/config"
	"payments-go/infrastructure/database/dynamodb"
	"payments-go/infrastructure/database/redis"
	"payments-go/infrastructure/http/router"
)

type config struct {
	configApp *appConfig.AppConfig
	webServer router.Server
	db *dynamodb.DynamoDBClient
	cache *redis.RedisTemplate
	httpClient gateway.HttpClient
	repo coreRepository.PaymentsRepository
}