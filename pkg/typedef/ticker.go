package typedef

type Ticker struct {
	At        int64       `json:"at"`
	BaseUnit  string      `json:"base_unit"`
	Buy       string      `json:"buy"`
	High      string      `json:"high"`
	Last      string      `json:"last"`
	Low       string      `json:"low"`
	Name      string      `json:"name"`
	Open      interface{} `json:"open"`
	QuoteUnit string      `json:"quote_unit"`
	Sell      string      `json:"sell"`
	Type      string      `json:"type"`
	Volume    string      `json:"volume"`
}
