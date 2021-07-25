package router

import (
	"net/http"
	"os"

	"github.com/rysmaadit/finantier_test/stock_service/handler"
	"github.com/rysmaadit/finantier_test/stock_service/middleware"
	"github.com/rysmaadit/finantier_test/stock_service/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(dependencies service.Dependencies) http.Handler {
	r := mux.NewRouter()

	privateRouter := r.NewRoute().Subrouter()
	privateRouter.Use(middleware.AuthMiddleware(dependencies))

	setStockRouter(privateRouter, dependencies)
	setAuthRouter(r, dependencies)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	return loggedRouter
}

func setStockRouter(router *mux.Router, dependencies service.Dependencies) {
	router.Methods(http.MethodGet).Path("/stock/{stock_code}").Handler(handler.GetStockByCodeHandler(dependencies))
}

func setAuthRouter(router *mux.Router, dependencies service.Dependencies) {
	router.Methods(http.MethodGet).Path("/auth/token").Handler(handler.GetTokenHandler(dependencies))
}
