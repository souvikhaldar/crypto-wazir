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
	ADA     CryptoCurrency = "ada"
	PUSH    CryptoCurrency = "push"
	BNB     CryptoCurrency = "bnb"
	XLM     CryptoCurrency = "xlm"
	INVALID CryptoCurrency = "INVALID"
)

func GetMtnFromCC(arg CryptoCurrency) string {
	//switch arg {
	//case BTC:
	//	return "btcinr"
	//case ETH:
	//	return "ethinr"
	//case ETC:
	//	return "etcinr"
	//case DOGE:
	//	return "dogeinr"
	//case BNB:
	//	return "bnbinr"
	//case ADA:
	//	return ADA
	//default:
	//	return ""
	//}
	return arg.String() + "inr"
}

func GetCCFromMtn(s string) CryptoCurrency {
	switch s {
	case "btcinr":
		return BTC
	case "ethinr":
		return ETH
	case "etcinr":
		return ETC
	case "dogeinr":
		return DOGE
	case "bnbinr":
		return BNB
	case "adainr":
		return ADA
	case "pushinr":
		return PUSH
	case "xlminr":
		return XLM
	default:
		return INVALID
	}
}

func ConvertLingo(arg string) string {
	return arg + "inr"
}
