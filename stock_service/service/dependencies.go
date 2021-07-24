package service

import (
	"net/http"

	"github.com/rysmaadit/finantier_test/stock_service/app"
	"github.com/rysmaadit/finantier_test/stock_service/external/auth_service_wrapper"
	"github.com/rysmaadit/finantier_test/stock_service/external/encryption_service_wrapper"
	"github.com/rysmaadit/finantier_test/stock_service/external/polygon"
)

type Dependencies struct {
	StockService      StockServiceInterface
	AuthWrapper       auth_service_wrapper.AuthServiceWrapperInterface
	EncryptionWrapper encryption_service_wrapper.EncryptionServiceWrapperInterface
}

func InstantiateDependencies(application *app.Application) Dependencies {
	httpClient := &http.Client{}
	polygonClient := polygon.NewClient(httpClient, application.Config.PolygonBaseURL, application.Config.PolygonAPIKey)
	encryptionWrapper := encryption_service_wrapper.New(httpClient, application.Config.EncryptionServiceBaseURL)
	stockService := NewStockService(application.Config, polygonClient, encryptionWrapper)
	authWrapper := auth_service_wrapper.New(httpClient, application.Config.AuthServiceBaseURL)

	return Dependencies{
		StockService: stockService,
		AuthWrapper:  authWrapper,
	}
}
