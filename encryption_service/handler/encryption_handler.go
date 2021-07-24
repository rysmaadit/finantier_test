package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/rysmaadit/finantier_test/encryption_service/contract"

	"github.com/rysmaadit/finantier_test/encryption_service/common/responder"
	"github.com/rysmaadit/finantier_test/encryption_service/service"
)

func EncryptionHandler(dependencies service.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := contract.NewGetEncryptionPayloadRequest(r)

		if err != nil {
			log.Errorln("error parse request payload: ", err)
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		resp, err := dependencies.EncryptionService.Encrypt(req)

		if err != nil {
			log.Errorln("error encrypt request payload: ", err)
			responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
			return
		}

		responder.NewHttpResponse(r, w, http.StatusOK, resp, nil)
	}
}
