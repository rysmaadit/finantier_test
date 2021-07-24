package contract

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/rysmaadit/finantier_test/auth_service/common/errors"
	"github.com/rysmaadit/finantier_test/auth_service/common/util"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type JWTMapClaim struct {
	Authorized bool   `json:"authorized"`
	RequestID  string `json:"requestID"`
	jwt.StandardClaims
}

type GetTokenResponseContract struct {
	Token string `json:"token"`
}

type ValidateTokenRequestContract struct {
	Token string `json:"token"`
}

func NewValidateTokenRequest(r *http.Request) (*ValidateTokenRequestContract, error) {
	validateTokenContract := new(ValidateTokenRequestContract)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(validateTokenContract); err != nil {
		log.Error(err)
		return nil, err
	}

	validate := validator.New()
	util.UseJsonFieldValidation(validate)

	if err := validate.Struct(validate); err != nil {
		log.Error(err)
		return nil, errors.NewValidationError(errors.ValidateErrToMapString(err.(validator.ValidationErrors)))
	}

	return validateTokenContract, nil
}
