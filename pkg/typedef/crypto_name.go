package typedef

type CryptoCurrency string

func (c CryptoCurrency) String() string {
	return string(c)
}

var (
	BTC     CryptoCurrency = "btc"
	ETH     CryptoCurrency = "eth"
	ETC     CryptoCurrency = "etc"
	DOGE    CryptoCurrency = "doge"
	INVALID CryptoCurrency = "INVALID"
)

func GetMtnFromCC(arg CryptoCurrency) string {
	switch arg {
	case BTC:
		return "btcinr"
	case ETH:
		return "ethinr"
	case ETC:
		return "etcinr"
	case DOGE:
		return "dogeinr"
	default:
		return ""
	}
}

func GetCCFromMtn(s string) CryptoCurrency {
	switch s {
	case "btcnr":
		return BTC
	case "ethinr":
		return ETH
	case "etcinr":
		return ETC
	case "dogeinr":
		return DOGE
	default:
		return INVALID
	}
}

func ConvertLingo(arg string) string {
	return arg + "inr"
}
