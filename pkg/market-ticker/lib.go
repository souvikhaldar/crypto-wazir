package market-ticker
import "github.com/souvikhaldar/cw/pkg/typedef"

type Repo interface{
	GetAllMarketTicker()(map[string]interface{},err)
}

GetMarketTicker(typedef.CryptoCurrency)(m map[string]interface{},err err){

}


