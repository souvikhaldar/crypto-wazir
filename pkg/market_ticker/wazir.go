package market_ticker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/souvikhaldar/cw/pkg/typedef"
)

type Wazir struct {
}

func NewWazir() *Wazir {
	return &Wazir{}
}

func GetMarketTicker(
	cc typedef.CryptoCurrency,
	resp map[string]typedef.Ticker,
) (*typedef.Ticker, error) {

	tn := typedef.GetMtnFromCC(cc)
	t, ok := resp[tn]
	if !ok {
		return nil, fmt.Errorf("No entry for %s", cc)
	}
	return &t, nil
}

func (w *Wazir) GetAllMarketTicker() (
	m map[string]typedef.Ticker,
	err error,
) {

	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		"https://api.wazirx.com/api/v2/tickers",
		nil,
	)
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&m)
	return
}

func (w *Wazir) GetLatestPrice(mtn string) (price float64, err error) {

	m, err := w.GetAllMarketTicker()
	if err != nil || m == nil {
		return
	}

	cc := typedef.GetCCFromMtn(mtn)
	t, err := GetMarketTicker(cc, m)
	if err != nil || t == nil {
		return

	}
	price, err = strconv.ParseFloat(t.Last, 64)
	return

}
