package utils

import (
	"api-openlive/common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"time"
)

const convertUrl = "https://api.coingecko.com/api/v3/coins/%v/market_chart?vs_currency=usd&days=%v&interval=daily"
const opv = "openlive-nft"

type HttpOption struct {
	Timeout time.Duration
}

type MarketChart []float64

type MarketChartResponse struct {
	Prices       []MarketChart `json:"prices"`
	MarketCaps   []MarketChart `json:"market_caps"`
	TotalVolumes []MarketChart `json:"total_volumes"`
}

// SendReqGet send http get
func SendReqGet(url string, headers map[string]string, opts ...HttpOption) (int, []byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, nil, err
	}
	if len(headers) > 0 {
		for k, val := range headers {
			req.Header.Set(k, val)
		}
	}

	client := &http.Client{}
	if len(opts) > 0 && opts[0].Timeout != 0 {
		client.Timeout = opts[0].Timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer func() {
		req.Close = true
		resp.Body.Close()
	}()
	body, _ := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, body, nil
}

func GetExchangeRates(target string, days int64) (map[string]float64, error) {
	opvRate := MarketChartResponse{}
	targetRate := MarketChartResponse{}
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	opvUrl := fmt.Sprintf(convertUrl, opv, days)
	targetUrl := fmt.Sprintf(convertUrl, target, days)

	// get convert rate
	statusCode, opvBody, err := SendReqGet(opvUrl, headers)

	if statusCode != 200 || err != nil {
		return nil, err
	}

	statusCode, targetBody, err := SendReqGet(targetUrl, headers)

	if statusCode != 200 || err != nil {
		return nil, err
	}

	// unmarshal response
	err = json.Unmarshal(opvBody, &opvRate)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(targetBody, &targetRate)
	if err != nil {
		return nil, err
	}

	targetRateByDay := make(map[string]float64)
	for _, v := range targetRate.Prices {
		timeFormat := ConvertTimeFromMiliSec(int64(v[0])).Format(common.ISO8601)
		targetRateByDay[timeFormat] = v[1]
	}

	result := make(map[string]float64)
	for _, v := range opvRate.Prices {
		timeFormat := ConvertTimeFromMiliSec(int64(v[0])).Format(common.ISO8601)
		result[timeFormat] = math.Round((v[1]/targetRateByDay[timeFormat])*100) / 100
	}

	return result, nil
}
