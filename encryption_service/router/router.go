package router

import (
	"net/http"
	"os"

	"github.com/rysmaadit/finantier_test/encryption_service/handler"
	"github.com/rysmaadit/finantier_test/encryption_service/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(dependencies service.Dependencies) http.Handler {
	r := mux.NewRouter()

	setEncryptionRouter(r, dependencies)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	return loggedRouter
}

func setEncryptionRouter(router *mux.Router, dependencies service.Dependencies) {
	router.Methods(http.MethodPost).Path("/encrypt").Handler(handler.EncryptionHandler(dependencies))
}
