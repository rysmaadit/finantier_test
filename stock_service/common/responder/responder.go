package responder

import (
	"encoding/json"
	pkgErrors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/rysmaadit/finantier_test/stock_service/common/errors"
	"net/http"
)

type Template struct {
	Status bool        `json:"status"`
	Error  interface{} `json:"error"`
	Result interface{} `json:"result"`
}

func NewHttpResponse(r *http.Request, w http.ResponseWriter, httpCode int, result interface{}, err error) {
	if err != nil {
		Error(r, w, err, httpCode)
	} else {
		Success(w, result, httpCode)
	}
}

func Error(r *http.Request, w http.ResponseWriter, err error, httpCode int) {
	switch err := pkgErrors.Cause(err).(type) {
	case *errors.NotFoundError:
		notFoundError(r, w, err)
	case *errors.DBError, *errors.ExternalError:
		serviceUnavailableError(r, w, err)
	case *errors.BadRequestError:
		badRequestError(r, w, err)
	case *errors.UnauthorizedError:
		unauthorizedError(r, w, err)
	default:
		GenericError(r, w, err, err.Error(), httpCode)
	}
}

func Success(w http.ResponseWriter, successResponse interface{}, responseCode ...int) {
	w.Header().Set("Content-Type", "application/json")

	if len(responseCode) == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(responseCode[0])
	}

	t := Template{
		Status: true,
		Result: successResponse,
		Error:  nil,
	}

	if successResponse != nil {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(t)
	}
}

func GenericError(r *http.Request, w http.ResponseWriter, err error, errorResponse interface{}, responseCode int) {
	log := logrus.WithFields(logrus.Fields{
		"Method": r.Method,
		"Host":   r.Host,
		"Path":   r.URL.Path,
	}).WithField("ResponseCode", responseCode)

	if responseCode < 500 {
		log.Warn(err.Error())
	} else {
		log.Error(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)

	t := Template{
		Status: false,
		Result: nil,
		Error:  errorResponse,
	}

	if errorResponse != nil {
		_ = json.NewEncoder(w).Encode(t)
	}
}

func notFoundError(r *http.Request, w http.ResponseWriter, err error) {
	GenericError(r, w, err, err.Error(), http.StatusNotFound)
}

func serviceUnavailableError(r *http.Request, w http.ResponseWriter, err error) {
	GenericError(r, w, err, err.Error(), http.StatusServiceUnavailable)
}

func badRequestError(r *http.Request, w http.ResponseWriter, err *errors.BadRequestError) {
	GenericError(r, w, err, err.Error(), http.StatusBadRequest)
}

func unauthorizedError(r *http.Request, w http.ResponseWriter, err *errors.UnauthorizedError) {
	GenericError(r, w, err, err.Error(), http.StatusUnauthorized)
}
