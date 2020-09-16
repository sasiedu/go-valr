package valr

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Currency struct {
	Symbol    string
	IsActive  bool
	ShortName string
	LongName  string
}

func (v *Valr) GetCurrencies() (currencies []Currency, err error) {
	resp, err := v.client.do("GET", "/public/currencies", []byte(""), false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &currencies)
	return
}

func (v *Valr) GetPublicOrderBook(currencyPair string) (orderBook *OrderBook, err error) {
	path := fmt.Sprintf("/public/%s/orderbook", strings.ToUpper(currencyPair))
	resp, err := v.client.do("GET", path, []byte(""), false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &orderBook)
	return
}

type CurrencyPair struct {
	Symbol         string
	BaseCurrency   string
	QuoteCurrency  string
	ShortName      string
	Active         bool
	MinBaseAmount  float64 `json:",string"`
	MaxBaseAmount  float64 `json:",string"`
	MinQuoteAmount float64 `json:",string"`
	MaxQuoteAmount float64 `json:",string"`
}

func (v *Valr) GetCurrencyPairs() (currencyPairs []CurrencyPair, err error) {
	resp, err := v.client.do("GET", "/public/pairs", []byte(""), false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &currencyPairs)
	return
}

type CurrencyOrderTypes struct {
	CurrencyPair string
	OrderTypes   []string
}

func (v *Valr) GetAllCurrencyPairOrderTypes() (currencyPairsOrderTypes []CurrencyOrderTypes, err error) {
	resp, err := v.client.do("GET", "/public/ordertypes", []byte(""), false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &currencyPairsOrderTypes)
	return
}

func (v *Valr) GetOrderTypesForCurrencyPair(currencyPair string) (orderTypes []string, err error) {
	path := fmt.Sprintf("/public/%s/ordertypes", strings.ToUpper(currencyPair))
	resp, err := v.client.do("GET", path, []byte(""), false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &orderTypes)
	return
}

type MarketSummary struct {
	CurrencyPair       string
	AskPrice           float64 `json:",string"`
	BidPrice           float64 `json:",string"`
	LastTradedPrice    float64 `json:",string"`
	PreviousClosePrice float64 `json:",string"`
	BaseVolume         float64 `json:",string"`
	HighPrice          float64 `json:",string"`
	LowPrice           float64 `json:",string"`
	Created            string
	ChangeFromPrevious float64 `json:",string"`
}

func (v *Valr) GetAllCurrencyPairMarketSummary() (marketSummaries []MarketSummary, err error) {
	resp, err := v.client.do("GET", "/public/marketsummary", []byte(""), false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &marketSummaries)
	return
}

func (v *Valr) GetMarketSummaryForCurrencyPair(currencyPair string) (marketSummary *MarketSummary, err error) {
	path := fmt.Sprintf("/public/%s/marketsummary", strings.ToUpper(currencyPair))
	resp, err := v.client.do("GET", path, []byte(""), false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &marketSummary)
	return
}

type ServerTime struct {
	EpochTime uint64
	Time      string
}

func (v *Valr) GetServerTime() (serverTime *ServerTime, err error) {
	resp, err := v.client.do("GET", "/public/time", []byte(""), false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &serverTime)
	return
}
