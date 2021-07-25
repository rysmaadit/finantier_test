package service

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/rysmaadit/finantier_test/encryption_service/config"
	"github.com/rysmaadit/finantier_test/encryption_service/contract"
)

func Test_encryptionService_Encrypt(t *testing.T) {
	const (
		sampleKey = "abcdefghijklmnopqrstuvwxyz012390"
	)

	req := &contract.GetDailyTimeSeriesStockResponse{
		Status:     "",
		From:       "",
		Symbol:     "",
		Open:       0,
		High:       0,
		Low:        0,
		Close:      0,
		Volume:     0,
		AfterHours: 0,
		PreMarket:  0,
	}

	reqBytes, _ := json.Marshal(req)

	type fields struct {
		appConfig *config.Config
	}
	type args struct {
		req *contract.GetDailyTimeSeriesStockResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *contract.EncryptedDataResponse
		wantErr bool
	}{
		{
			name:    "given correct contract should return no error",
			args:    args{req: req},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appConfig := &config.Config{
				EncryptionKey: sampleKey,
			}
			s := &encryptionService{
				appConfig: appConfig,
			}
			got, err := s.Encrypt(tt.args.req)

			if err != nil {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
			}

			decrypted, err := aes256CBCDecryptor([]byte(sampleKey), got.Data)

			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(decrypted, reqBytes) {
				t.Errorf("Encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
