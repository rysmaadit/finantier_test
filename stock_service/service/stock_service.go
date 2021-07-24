package service

import (
	"fmt"

	"github.com/rysmaadit/finantier_test/stock_service/common/errors"
	"github.com/rysmaadit/finantier_test/stock_service/config"
	"github.com/rysmaadit/finantier_test/stock_service/contract"
	"github.com/rysmaadit/finantier_test/stock_service/external/encryption_service_wrapper"
	"github.com/rysmaadit/finantier_test/stock_service/external/polygon"
	log "github.com/sirupsen/logrus"
)

type stockService struct {
	appConfig         *config.Config
	polygonClient     polygon.PolygonClientInterface
	encryptionWrapper encryption_service_wrapper.EncryptionServiceWrapperInterface
}

type StockServiceInterface interface {
	GetEncryptedStockData(request *contract.GetStockByCodeContractRequest) (*contract.GetStockByCodeContractResponse, error)
}

func NewStockService(appConfig *config.Config,
	polygon polygon.PolygonClientInterface,
	encryptionWrapper encryption_service_wrapper.EncryptionServiceWrapperInterface) *stockService {
	return &stockService{
		appConfig:         appConfig,
		polygonClient:     polygon,
		encryptionWrapper: encryptionWrapper,
	}
}

func (s *stockService) GetEncryptedStockData(request *contract.GetStockByCodeContractRequest) (*contract.GetStockByCodeContractResponse, error) {
	resp, err := s.polygonClient.GetDailyTimeSeriesStock(request.Code)

	if err != nil {
		errMsg := fmt.Sprintf("error fetch data from external source, message: %s", err.Error())
		log.Errorf(errMsg)
		return nil, errors.NewExternalError(errMsg)
	}

	if resp == nil {
		errMsg := fmt.Sprintf("empty response data from external source")
		log.Errorf(errMsg)
		return nil, errors.NewExternalError(errMsg)
	}

	displayStockInLog(resp)

	encryptedData, err := s.encryptionWrapper.Encrypt(resp)

	if err != nil {
		errMsg := fmt.Sprintf("error encrypt stock data, err: %v", err)
		log.Errorf(errMsg)
		return nil, errors.NewExternalError(errMsg)
	}

	response := &contract.GetStockByCodeContractResponse{Data: encryptedData.Result.Data}
	return response, nil
}

func displayStockInLog(resp *polygon.GetDailyTimeSeriesStockResponse) {
	log.Infoln(fmt.Sprintf(
		"\nDate: %v\n"+
			"Symbol: %v\n"+
			"Open: %v\n"+
			"High: %v\n"+
			"Low: %v\n"+
			"Close: %v\n"+
			"Volume: %v\n"+
			"AfterHours: %v\n"+
			"PreMarket: %v\n",
		resp.From,
		resp.Symbol,
		resp.Open,
		resp.High,
		resp.Low,
		resp.Close,
		resp.Volume,
		resp.AfterHours,
		resp.PreMarket))
}
