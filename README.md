# Crypto-Wazir
## Introduction
`cw` is the command line tool that uses **wazirx** public API to track and monitor various crypto-currencies to help make smart decisions for investing. 

## Installation
1. Copy link to the binary of your current system from *Releases*
2. Download it `wget https://github.com/souvikhaldar/crypto-wazir/releases/download/1.0.0/crypto-wazir<version_os>tar.gz`.
3. Extract it using `tar -xf <filename>.tar.gz`
4. Make it executable and copy it to $PATH location `chmod +x crypto-wazir && sudo mv crypto-wazir /usr/local/bin/cw`

## Usage
### Fetch the current market information
You can fetch the current market statistics of your preferred crypto using the `fetch` command.  
Eg. for getting details of bitcoin you can do:  
`cw fetch -c btc`

```
Fetching details of: btc
Details of btc : {At:1620556416 BaseUnit:btc Buy:4505403.0 High:4660025.0 Last:4508000.0 Low:4427663.0 Name:BTC/INR Open:4.529e+06 QuoteUnit:inr Sell:4508000.0 Type:SPOT Volume:404.55108}
```

### Monitor the current price
You can monitor the price of the crypto every 10 secs interval using `monitor command`.
Eg:
Monitoring the price of **dogecoin**:
```
cw monitor -c doge
Monitoring:  doge
Price of dogeinr is:38.780000
Price of dogeinr is:38.653800
Price of dogeinr is:38.600000
```

### Knowing when to buy and sell
This an interesting use of this tool that would make you invest/withdraw smartly.  
The price of these crypto currencies and always changing, that too very rapidly! So the best way to be *happy* is to set limits for benefit and loss. Along with it follow a generic principle that, when the price goes down, buy more of that coin and when the price goes up, sell at a point which can make you content.  

The way I would do it is, let's say I've purchased `doge` at 35 INR per coin, so I would set the limit that if the price goes below 25 INR I would buy more and if it reaches a price of 60 INR per coin, I would sell it!  
The command for this would be:  
`cw monitor -c doge -l 25 -u 60` here -u flag stands for upper limit and -l stands for lower limit.  


### Telegram bot
Here is the [bot](t.me/go_wazirx_bot)
