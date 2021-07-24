package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/rysmaadit/finantier_test/auth_service/common/errors"
	"github.com/rysmaadit/finantier_test/auth_service/config"
	"github.com/rysmaadit/finantier_test/auth_service/contract"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type authService struct {
	appConfig *config.Config
}

type AuthServiceInterface interface {
	GetToken() (*contract.GetTokenResponseContract, error)
	VerifyToken(req *contract.ValidateTokenRequestContract) (*contract.JWTMapClaim, error)
}

func NewAuthService(appConfig *config.Config) *authService {
	return &authService{
		appConfig: appConfig,
	}
}

func (s *authService) GetToken() (*contract.GetTokenResponseContract, error) {
	atClaims := contract.JWTMapClaim{
		Authorized: true,
		RequestID:  uuid.New().String(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(s.appConfig.JWTSecret))

	if err != nil {
		errMsg := fmt.Sprintf("error signed JWT credentials: %v", err)
		log.Errorf(errMsg)
		return nil, errors.NewInternalError(err, errMsg)
	}

	return &contract.GetTokenResponseContract{Token: token}, err
}

func (s *authService) VerifyToken(req *contract.ValidateTokenRequestContract) (*contract.JWTMapClaim, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(req.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.appConfig.JWTSecret), nil
	})

	if err != nil {
		log.Error("error verified token", err)
		return nil, errors.NewUnauthorizedError("error claim token")
	}

	if !token.Valid {
		return nil, errors.NewUnauthorizedError("invalid token")
	}

	authorized := fmt.Sprintf("%v", claims["authorized"])
	requestID := fmt.Sprintf("%v", claims["requestID"])

	if authorized == "" || requestID == "" {
		return nil, errors.NewUnauthorizedError("invalid payload")
	}

	ok, err := strconv.ParseBool(authorized)

	if err != nil || !ok {
		return nil, errors.NewUnauthorizedError("invalid payload")
	}

	resp := &contract.JWTMapClaim{
		Authorized:     claims["authorized"].(bool),
		RequestID:      claims["requestID"].(string),
		StandardClaims: jwt.StandardClaims{},
	}

	return resp, nil
}
