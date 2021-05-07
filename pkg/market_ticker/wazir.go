package market_ticker

import (
	"github.com/souvikhaldar/cw/pkg/typedef"
)

type Wazir struct{
	Ticker map[string]typedef.Ticker
}

func NewWazir()*Wazir{
	return &Wazir{
		Ticker: nil,
	}
}


func (w *Wazir)GetAllMarketTicker()(m map[string]interface{}, err error){

		client := &http.Client{}
		var m map[string]interface{}

		req, err := http.NewRequest(
			"GET",
			"https://api.wazirx.com/api/v2/tickers",
			nil,
		)
		if err != nil {
			fmt.Println("Unable to create the request: ", err)
			return
		}
		req.Header.Add("Accept", "application/json")

		resp, err := client.Do(req)
		defer resp.Body.Close()

		m,err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error in reading resp:", err)
		}
		mt := make(map[string]Ticker)
		t, err := json.Unmarshal(m,&mt)
		if err != nil{
			return
		}
		w.Ticker = mt
		return
}

func GetMarketTicker(
	c typedef.CryptoCurrency,
	resp map[string]typedef.Ticker,
)(*typedef.Ticker,err error){
	tn := GetTickerName(c)
	t,ok := resp[tn]
	if !ok{
		return nil,fmt.Errorf("No entry for %s",c)
	}
	return &t,nil
}

func (w *Wazir)GetLatestPrice(c typedef.CryptoCurrency)(price float64,err error){

	if w.Ticker != nil{
		t, err := GetMarketTicker(c,w.Ticker)
		if t !=nil{
			price, _ = strconv.Atoi(t.Last)
			return

		}
	}
	m, err := w.GetAllMarketTicker()
	if err != nil{
		return 
	}

}
