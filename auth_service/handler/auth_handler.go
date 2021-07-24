package handler

import (
	"net/http"

	"github.com/rysmaadit/finantier_test/auth_service/common/responder"
	"github.com/rysmaadit/finantier_test/auth_service/contract"
	"github.com/rysmaadit/finantier_test/auth_service/service"
	log "github.com/sirupsen/logrus"
)

func GetToken(dependencies service.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := dependencies.AuthService.GetToken()

		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
			return
		}

		responder.NewHttpResponse(r, w, http.StatusOK, resp, nil)
	}
}

func ValidateToken(dependencies service.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := contract.NewValidateTokenRequest(r)

		if err != nil {
			log.Error(err)
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		resp, err := dependencies.AuthService.VerifyToken(req)

		if err != nil {
			log.Error(err)
			responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
			return
		}

		responder.NewHttpResponse(r, w, http.StatusOK, resp, nil)
		return
	}
}
