package typedef

type CryptoCurrency string

func (c CryptoCurrency)String()string{
	return string(c)
}

var(
	BTC CryptoCurrency = "btc"
	ETH CryptoCurrency = "eth"
	ETC CryptoCurrency = "etc"
	DOGE CryptoCurrency = "doge"
	INVALID CryptoCurrency = "INVALID"
)
