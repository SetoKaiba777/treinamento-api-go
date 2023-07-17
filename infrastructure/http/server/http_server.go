package server

import (
	"payments-go/adapter/httpclient"
	"payments-go/adapter/repository"
	"payments-go/core/gateway"
	coreRepository "payments-go/core/repository"
	appConfig "payments-go/infrastructure/config"
	"payments-go/infrastructure/database/dynamodb"
	"payments-go/infrastructure/database/redis"
	"payments-go/infrastructure/http/router"
	"payments-go/infrastructure/logger"
	"strconv"
	"time"
)

type config struct {
	configApp  *appConfig.AppConfig
	webServer  router.Server
	db         *dynamodb.DynamoDBClient
	cache      *redis.RedisTemplate
	httpClient gateway.HttpClient
	repo       coreRepository.PaymentsRepository
}

func NewConfig() *config {
	return &config{}
}

func (c *config) WithAppConfig() *config {
	var err error
	c.configApp, err = appConfig.LoadConfig()
	if err != nil {
		logger.Fatal(err)
	}
	return c
}

func (c *config) InitLogger() *config {
	logger.NewZapLogger()
	logger.Infof("Log has been successfully configured")
	return c
}

func (c *config) WithDB() *config {
	c.db = dynamodb.NewDynamoDBClient(c.configApp.Aws.Region, c.configApp.Aws.Endpoint)
	logger.Infof("DB has been successfully configured")
	return c
}

func (c *config) WithCache() *config {
	c.cache = redis.New(c.configApp.Redis.Host)
	logger.Infof("Redis has been successfully configured")
	return c
}

func (c *config) WithHttpClient() *config {
	c.httpClient = httpclient.NewHttpClient()
	logger.Infof("Http Client has been successfully configured")
	return c
}

func (c *config) WithRepository() *config {
	c.repo = repository.NewPaymentRepository(c.cache, c.db)
	logger.Infof("Repository has been successfully configured")
	return c
}

func (c *config) WithWebServer() *config {
	intPort, err := strconv.ParseInt(c.configApp.Application.Server.Port, 10, 64)
	if err != nil {
		logger.Fatal(err)
	}

	intDuration, err := time.ParseDuration(c.configApp.Application.Server.Timeout + "s")
	if err != nil {
		logger.Fatal(err)
	}

	c.webServer = router.NewGinServer(intPort, intDuration, c.httpClient, c.repo)
	logger.Infof("Web server has been successfully configurated")
	return c
}

func (c *config) Start() {
	logger.Infof("App running on port %s", c.configApp.Application.Server.Port)
	c.webServer.Listen()
	
}
