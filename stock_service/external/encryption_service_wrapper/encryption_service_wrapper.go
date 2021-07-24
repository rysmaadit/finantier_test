package encryption_service_wrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rysmaadit/finantier_test/stock_service/external/polygon"

	log "github.com/sirupsen/logrus"
)

type encryptionServiceWrapper struct {
	HTTPClient *http.Client
	BaseURL    string
}

type EncryptionServiceWrapperInterface interface {
	Encrypt(payload *polygon.GetDailyTimeSeriesStockResponse) (*EncryptedResponseContract, error)
}

func New(httpClient *http.Client, baseURL string) *encryptionServiceWrapper {
	return &encryptionServiceWrapper{
		HTTPClient: httpClient,
		BaseURL:    baseURL,
	}
}

func (asw *encryptionServiceWrapper) Encrypt(payload *polygon.GetDailyTimeSeriesStockResponse) (*EncryptedResponseContract, error) {
	fullURL := fmt.Sprintf("%s/encrypt", asw.BaseURL)

	reqBodyBytes, err := json.Marshal(payload)

	if err != nil {
		log.Error("error marshal request body: ", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fullURL, bytes.NewBuffer(reqBodyBytes))

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
		response := new(EncryptedResponseContract)
		err := json.Unmarshal(bodyBytes, response)
		return response, err
	}

	log.Error(fmt.Sprintf("error external service, with URL: %s, status_code: %d, response: %v", "", res.StatusCode, string(bodyBytes)))
	return nil, fmt.Errorf("%s", string(bodyBytes))
}
