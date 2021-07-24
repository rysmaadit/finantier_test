package polygon

type GetDailyTimeSeriesStockResponse struct {
	Status     string  `json:"status"`
	From       string  `json:"from"`
	Symbol     string  `json:"symbol"`
	Open       float32 `json:"open"`
	High       float32 `json:"high"`
	Low        float32 `json:"low"`
	Close      float32 `json:"close"`
	Volume     int64   `json:"volume"`
	AfterHours float32 `json:"afterHours"`
	PreMarket  float32 `json:"preMarket"`
}
