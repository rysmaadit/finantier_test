package service

import (
	"reflect"
	"testing"

	"github.com/rysmaadit/finantier_test/stock_service/common/errors"
	"github.com/rysmaadit/finantier_test/stock_service/config"
	"github.com/rysmaadit/finantier_test/stock_service/contract"
	"github.com/rysmaadit/finantier_test/stock_service/external/encryption_service_wrapper"
	"github.com/rysmaadit/finantier_test/stock_service/external/mocks"
	"github.com/rysmaadit/finantier_test/stock_service/external/polygon"
)

func Test_stockService_GetEncryptedStockData(t *testing.T) {
	const (
		stockCodeSample        = "TSLA"
		encryptedMessageSample = "encrypted message response"
	)

	dailyStockResponse := &polygon.GetDailyTimeSeriesStockResponse{
		Status:     "",
		From:       "",
		Symbol:     stockCodeSample,
		Open:       0,
		High:       0,
		Low:        0,
		Close:      0,
		Volume:     0,
		AfterHours: 0,
		PreMarket:  0,
	}

	response := &contract.GetStockByCodeContractResponse{Data: []byte(encryptedMessageSample)}

	encWrapperResp := &encryption_service_wrapper.EncryptedResponseContract{
		Status: false,
		Error:  nil,
		Result: encryption_service_wrapper.EncryptedData{
			Data: []byte(encryptedMessageSample),
		},
	}

	type fields struct {
		appConfig         *config.Config
		polygonClient     polygon.PolygonClientInterface
		encryptionWrapper encryption_service_wrapper.EncryptionServiceWrapperInterface
	}
	type args struct {
		request *contract.GetStockByCodeContractRequest
	}
	tests := []struct {
		name                  string
		fields                fields
		args                  args
		mockPolygon           func(mock *mocks.PolygonClientInterface)
		mockEncryptionWrapper func(mock *mocks.EncryptionServiceWrapperInterface)
		want                  *contract.GetStockByCodeContractResponse
		wantErr               bool
	}{
		{
			name: "given correct contract should return no error",
			args: args{request: &contract.GetStockByCodeContractRequest{Code: stockCodeSample}},
			mockPolygon: func(mock *mocks.PolygonClientInterface) {
				mock.On("GetDailyTimeSeriesStock", stockCodeSample).Return(dailyStockResponse, nil)
			},
			mockEncryptionWrapper: func(mock *mocks.EncryptionServiceWrapperInterface) {
				mock.On("Encrypt", dailyStockResponse).Return(encWrapperResp, nil)
			},
			want:    response,
			wantErr: false,
		},
		{
			name: "given correct contract, error get stock API should return error",
			args: args{request: &contract.GetStockByCodeContractRequest{Code: stockCodeSample}},
			mockPolygon: func(mock *mocks.PolygonClientInterface) {
				mock.On("GetDailyTimeSeriesStock", stockCodeSample).
					Return(nil, errors.New("err"))
			},
			mockEncryptionWrapper: func(mock *mocks.EncryptionServiceWrapperInterface) {},
			want:                  nil,
			wantErr:               true,
		},
		{
			name: "given correct contract, error encrypt stock API should return error",
			args: args{request: &contract.GetStockByCodeContractRequest{Code: stockCodeSample}},
			mockPolygon: func(mock *mocks.PolygonClientInterface) {
				mock.On("GetDailyTimeSeriesStock", stockCodeSample).Return(dailyStockResponse, nil)
			},
			mockEncryptionWrapper: func(mock *mocks.EncryptionServiceWrapperInterface) {
				mock.On("Encrypt", dailyStockResponse).Return(nil, errors.New("err"))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			polygonClientMock := new(mocks.PolygonClientInterface)
			encryptionWrapperMock := new(mocks.EncryptionServiceWrapperInterface)
			s := &stockService{
				appConfig:         tt.fields.appConfig,
				polygonClient:     polygonClientMock,
				encryptionWrapper: encryptionWrapperMock,
			}

			tt.mockPolygon(polygonClientMock)
			tt.mockEncryptionWrapper(encryptionWrapperMock)

			got, err := s.GetEncryptedStockData(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEncryptedStockData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEncryptedStockData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
