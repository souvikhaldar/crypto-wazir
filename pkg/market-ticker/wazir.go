package market-ticker

type Wazir struct{
}

func NewWazir()*Wazir{
	return &Wazir{}
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
		if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
			fmt.Println("Error in decoding the JSON: ", err)
			return 
		}
}
