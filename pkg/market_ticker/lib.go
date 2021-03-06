package market_ticker

import "github.com/souvikhaldar/cw/pkg/typedef"

type Repo interface {
	GetAllMarketTicker() (map[string]typedef.Ticker, error)
	GetLatestPrice(string) (float64, error)
}

func GetPrice(r Repo, mtn string) (float64, error) {
	return r.GetLatestPrice(mtn)
}
func GetAllDetails(r Repo, mtn string) (*typedef.Ticker, error) {
	m, err := r.GetAllMarketTicker()
	if err != nil {
		return nil, err
	}
	return GetMarketTicker(
		typedef.GetCCFromMtn(mtn),
		m,
	)
}
