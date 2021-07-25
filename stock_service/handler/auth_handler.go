package handler

import (
	"net/http"

	"github.com/rysmaadit/finantier_test/stock_service/common/responder"
	"github.com/rysmaadit/finantier_test/stock_service/service"
	log "github.com/sirupsen/logrus"
)

func GetTokenHandler(dependencies service.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := dependencies.AuthWrapper.GetToken()

		if err != nil {
			log.Error(err)
			responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
			return
		}

		responder.NewHttpResponse(r, w, http.StatusOK, result.Result, nil)
	}
}
