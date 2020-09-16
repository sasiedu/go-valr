package valr

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	Side            OrderSide
	Quantity        float64 `json:",string"`
	Price           float64 `json:",string"`
	CurrencyPair    string
	ID              string
	PositionAtPrice uint8
	OrderCount      uint8
}

type OrderBook struct {
	Asks []Order
	Bids []Order
}

func (v *Valr) GetOrderBook(currencyPair string) (orderBook *OrderBook, err error) {
	path := fmt.Sprintf("/marketdata/%s/orderbook", currencyPair)
	resp, err := v.client.do("GET", path, []byte(""), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &orderBook)
	return
}

func (v *Valr) GetNonAggregatedOrderBook(currencyPair string) (orderBook *OrderBook, err error) {
	path := fmt.Sprintf("/marketdata/%s/orderbook/full", currencyPair)
	resp, err := v.client.do("GET", path, []byte(""), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &orderBook)
	return
}

type Trade struct {
	Price        float64 `json:",string"`
	Quantity     float64 `json:",string"`
	CurrencyPair string
	TradeAt      string
	TakerSide    OrderSide
	SequenceID   uint32
	ID           string
}

func (v *Valr) GetCurrencyPairTradeHistory(currencyPair string, limit uint8) (history []Trade, err error) {
	path := fmt.Sprintf("/marketdata/%s/tradehistory?limit=%d", currencyPair, limit)
	resp, err := v.client.do("GET", path, []byte(""), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &history)
	return
}
