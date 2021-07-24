package middleware

import (
	"context"
	"encoding/json"
	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/gorilla/mux"
	"github.com/rysmaadit/finantier_test/stock_service/common/errors"
	"github.com/rysmaadit/finantier_test/stock_service/common/responder"
	"github.com/rysmaadit/finantier_test/stock_service/constant"
	"github.com/rysmaadit/finantier_test/stock_service/external/auth_service_wrapper"
	"github.com/rysmaadit/finantier_test/stock_service/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func AuthMiddleware(dependencies service.Dependencies) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString, err := jwtMiddleware.FromAuthHeader(r)

			if err != nil {
				log.Error("error get token from request header", err)
				responder.NewHttpResponse(r, w, http.StatusUnauthorized, nil, errors.NewUnauthorizedError("invalid token"))
				return
			}

			req := &auth_service_wrapper.AuthValidateTokenRequestContract{Token: tokenString}
			ctxValue, err := dependencies.AuthWrapper.ValidateToken(req)

			if err != nil {
				responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
				return
			}

			authCtxValue, err := json.Marshal(ctxValue)

			if err != nil {
				responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, constant.AuthContextRequestKey, authCtxValue)

			handler.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
