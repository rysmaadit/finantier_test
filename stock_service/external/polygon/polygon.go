package polygon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rysmaadit/finantier_test/stock_service/common/errors"

	log "github.com/sirupsen/logrus"
)

type PolygonClientInterface interface {
	GetDailyTimeSeriesStock(code string) (*GetDailyTimeSeriesStockResponse, error)
}

type polygonClient struct {
	HTTPClient *http.Client
	BaseURL    string
	APIKey     string
}

func NewClient(httpClient *http.Client, baseURL string, apiKey string) *polygonClient {
	return &polygonClient{
		HTTPClient: httpClient,
		BaseURL:    baseURL,
		APIKey:     apiKey,
	}
}

func (pc *polygonClient) GetDailyTimeSeriesStock(code string) (*GetDailyTimeSeriesStockResponse, error) {
	yesterdayDate := time.Now().AddDate(0, 0, -2).Format("2006-01-02")
	fullURL := fmt.Sprintf("%s/open-close/%s/%s?adjusted=true&apiKey=%s", pc.BaseURL, code, yesterdayDate, pc.APIKey)
	req, err := http.NewRequest(http.MethodGet, fullURL, nil)

	if err != nil {
		log.Error("error initiate HTTP request call: ", err)
		return nil, errors.NewInternalError(errors.New("internal server error"), "")
	}

	res, err := pc.HTTPClient.Do(req)

	if err != nil {
		log.Error("error do HTTP request call: ", err)
		return nil, errors.NewInternalError(errors.New("error call upstream API"), "")
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Error("error read response body: ", err)
		return nil, errors.NewExternalError("error fetch stock data from upstream")
	}

	if res.StatusCode == http.StatusOK {
		response := new(GetDailyTimeSeriesStockResponse)
		err := json.Unmarshal(bodyBytes, response)
		return response, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, errors.NewNotFoundError(fmt.Errorf("stock symbol: %s not found", code))
	}

	log.Error(fmt.Sprintf("error external service, with URL: %s, status_code: %d, response: %v", "", res.StatusCode, string(bodyBytes)))
	return nil, fmt.Errorf("%s", string(bodyBytes))
}
