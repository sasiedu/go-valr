package valr

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
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

type OrderBookOrder struct {
	Side         string
	Quantity     decimal.Decimal
	Price        decimal.Decimal
	CurrencyPair string
	OrderCount   uint32
}

type OrderBook struct {
	Asks []OrderBookOrder
	Bids []OrderBookOrder
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
	MinBaseAmount  decimal.Decimal
	MaxBaseAmount  decimal.Decimal
	MinQuoteAmount decimal.Decimal
	MaxQuoteAmount decimal.Decimal
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
	AskPrice           decimal.Decimal
	BidPrice           decimal.Decimal
	LastTradedPrice    decimal.Decimal
	PreviousClosePrice decimal.Decimal
	BaseVolume         decimal.Decimal
	HighPrice          decimal.Decimal
	LowPrice           decimal.Decimal
	Created            string
	ChangeFromPrevious decimal.Decimal
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
