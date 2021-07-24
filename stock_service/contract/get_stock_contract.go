package contract

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rysmaadit/finantier_test/stock_service/common/errors"
)

type GetStockByCodeContractRequest struct {
	Code string
}

type GetStockByCodeContractResponse struct {
	Data []byte `json:"data"`
}

func NewGetStockByCodeRequest(r *http.Request) (*GetStockByCodeContractRequest, error) {
	params := mux.Vars(r)
	stockCode := params["stock_code"]

	if stockCode == "" {
		return nil, errors.NewBadRequestError(errors.New("empty stock code value"))
	}

	contract := &GetStockByCodeContractRequest{Code: stockCode}
	return contract, nil
}
