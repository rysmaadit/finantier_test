package contract

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/rysmaadit/finantier_test/encryption_service/common/errors"
	"github.com/rysmaadit/finantier_test/encryption_service/common/util"
	log "github.com/sirupsen/logrus"
)

type GetDailyTimeSeriesStockResponse struct {
	Status     string  `json:"status" validate:"required"`
	From       string  `json:"from" validate:"required"`
	Symbol     string  `json:"symbol" validate:"required"`
	Open       float32 `json:"open" validate:"required"`
	High       float32 `json:"high" validate:"required"`
	Low        float32 `json:"low" validate:"required"`
	Close      float32 `json:"close" validate:"required"`
	Volume     int64   `json:"volume" validate:"required"`
	AfterHours float32 `json:"afterHours" validate:"required"`
	PreMarket  float32 `json:"preMarket" validate:"required"`
}

type EncryptedDataResponse struct {
	Data []byte `json:"data"`
}

func NewGetEncryptionPayloadRequest(r *http.Request) (*GetDailyTimeSeriesStockResponse, error) {
	hashPayloadRequest := new(GetDailyTimeSeriesStockResponse)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(hashPayloadRequest); err != nil {
		log.Error(err)
		return nil, err
	}

	validate := validator.New()
	util.UseJsonFieldValidation(validate)

	if err := validate.Struct(validate); err != nil {
		log.Error(err)
		return nil, errors.NewValidationError(errors.ValidateErrToMapString(err.(validator.ValidationErrors)))
	}

	return hashPayloadRequest, nil
}
