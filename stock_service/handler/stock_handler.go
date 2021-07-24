package handler

import (
	"net/http"

	"github.com/rysmaadit/finantier_test/stock_service/common/responder"
	"github.com/rysmaadit/finantier_test/stock_service/contract"
	"github.com/rysmaadit/finantier_test/stock_service/service"
	log "github.com/sirupsen/logrus"
)

func GetStockByCodeHandler(dependencies service.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stockCodeContract, err := contract.NewGetStockByCodeRequest(r)

		if err != nil {
			log.Warning(err)
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		result, err := dependencies.StockService.GetEncryptedStockData(stockCodeContract)

		if err != nil {
			log.Error(err)
			responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
			return
		}

		responder.NewHttpResponse(r, w, http.StatusOK, result, nil)
	}
}
