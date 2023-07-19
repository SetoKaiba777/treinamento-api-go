package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"payments-go/adapter/api/controller"
	"payments-go/adapter/facade"
	"payments-go/core/gateway"
	corepository "payments-go/core/repository"
	"payments-go/core/usecase"
	"payments-go/infrastructure/logger"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	Port int64

	Server interface {
		Listen()
	}

	ginEngine struct {
		router     *gin.Engine
		port       int64
		ctx        context.Context
		ctxTimeout time.Duration
		httpClient gateway.HttpClient
		repo       corepository.PaymentsRepository
	}
)

func NewGinServer(port int64,
	timeout time.Duration,
	httpClient gateway.HttpClient,
	repo corepository.PaymentsRepository) *ginEngine {
	return &ginEngine{
		router:     gin.New(),
		port:       port,
		ctxTimeout: timeout,
		httpClient: httpClient,
		repo:       repo,
	}
}

func (engine ginEngine) Listen() {
	gin.SetMode(gin.ReleaseMode)
	gin.Recovery()

	engine.setAppHandlers(engine.router)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", engine.port),
		Handler:      engine.router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Fatal("Error to starting HTTP Server")
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(engine.ctx, 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Sever Shutdown Failed")
	}
}

func (engine ginEngine) setAppHandlers(router *gin.Engine) {
	router.POST("/v1/payments", engine.HandlePayments())
	router.GET("/v1/payments/:paymentId", engine.HandleGetPayment())
}

func (e ginEngine) HandlePayments() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		checkBalanceUseCase := usecase.NewCheckBalanceUseCase(e.httpClient)
		checkFraudUseCase := usecase.NewCheckFraudUseCase(e.httpClient)
		debitBalanceUseCase := usecase.NewCheckDebitBalanceUseCase(e.httpClient)
		savePaymentsUseCase := usecase.NewSavePaymentsUseCase(e.repo)
		sendNotificationUseCase := usecase.NewSendNotificationUseCase(e.httpClient)

		f := facade.NewFacade(
			checkBalanceUseCase,
			checkFraudUseCase,
			debitBalanceUseCase,
			savePaymentsUseCase,
			sendNotificationUseCase)

		c := controller.NewCreatePaymentController(f)
		c.Execute(ctx.Writer, ctx.Request)

	}
}

func (e ginEngine) HandleGetPayment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		query.Add("paymentId", ctx.Param("paymentId"))
		ctx.Request.URL.RawQuery = query.Encode()

		uc := usecase.NewGetPaymentUseCase(e.repo)
		c := controller.NewGetPaymentController(uc)
		c.Execute(ctx.Writer, ctx.Request)
	}
}
