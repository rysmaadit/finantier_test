package auth_service_wrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type authServiceWrapper struct {
	HTTPClient *http.Client
	BaseURL    string
}

type AuthServiceWrapperInterface interface {
	GetToken() (*GetTokenResponseContract, error)
	ValidateToken(contract *AuthValidateTokenRequestContract) (*AuthValidateTokenResponseContract, error)
}

func New(httpClient *http.Client, baseURL string) *authServiceWrapper {
	return &authServiceWrapper{
		HTTPClient: httpClient,
		BaseURL:    baseURL,
	}
}

func (asw *authServiceWrapper) GetToken() (*GetTokenResponseContract, error) {
	fullURL := fmt.Sprintf("%s/auth/token", asw.BaseURL)
	req, err := http.NewRequest(http.MethodGet, fullURL, nil)

	if err != nil {
		log.Error("error initiate HTTP request call: ", err)
		return nil, err
	}

	res, err := asw.HTTPClient.Do(req)

	if err != nil {
		log.Error("error do HTTP request call: ", err)
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Error("error read response body: ", err)
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		response := new(GetTokenResponseContract)
		err := json.Unmarshal(bodyBytes, response)
		return response, err
	}

	log.Error(fmt.Sprintf("error external service, with URL: %s, status_code: %d, response: %v", "", res.StatusCode, string(bodyBytes)))
	return nil, fmt.Errorf("%s", string(bodyBytes))
}

func (asw *authServiceWrapper) ValidateToken(contract *AuthValidateTokenRequestContract) (*AuthValidateTokenResponseContract, error) {
	fullURL := fmt.Sprintf("%s/auth/token/validate", asw.BaseURL)

	bodyReqBytes, err := json.Marshal(contract)

	if err != nil {
		log.Error("error marshal request body: ", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fullURL, bytes.NewBuffer(bodyReqBytes))

	if err != nil {
		log.Error("error initiate HTTP request call: ", err)
		return nil, err
	}

	res, err := asw.HTTPClient.Do(req)

	if err != nil {
		log.Error("error do HTTP request call: ", err)
		return nil, err
	}

	bodyResBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Error("error read response body: ", err)
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		response := new(AuthValidateTokenResponseContract)
		err := json.Unmarshal(bodyResBytes, response)
		return response, err
	}

	log.Error(fmt.Sprintf("error external service, with URL: %s, status_code: %d, response: %v", "", res.StatusCode, string(bodyResBytes)))
	return nil, fmt.Errorf("%s", string(bodyResBytes))
}
