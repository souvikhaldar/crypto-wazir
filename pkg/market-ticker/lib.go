package market-ticker
import "github.com/souvikhaldar/cw/pkg/typedef"

type Repo interface{
	GetAllMarketTicker()(map[string]interface{},err)
	GetLatestPrice(cryto typedef.CryptoCurrency)(float64,error)
}



