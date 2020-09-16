package valr

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
)

type simpleBuySell struct {
	PayInCurrency string          `json:"payInCurrency"`
	PayAmount     decimal.Decimal `json:"payAmount"`
	Side          OrderSide       `json:"side"`
}

type Quote struct {
	CurrencyPair  string
	PayAmount     decimal.Decimal
	ReceiveAmount decimal.Decimal
	Fee           decimal.Decimal
	FeeCurrency   string
	Created       string
	ID            string
}

func (v *Valr) SimpleBuyQuote(currencyPair, payInCurrency string, amount decimal.Decimal) (quote *Quote, err error) {
	path := fmt.Sprintf("/simple/%s/quote", currencyPair)
	buy := simpleBuySell{payInCurrency, amount, BUY}

	body, err := structToBytes(buy)
	if err != nil {
		return
	}

	resp, err := v.client.do("POST", path, body, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &quote)
	return
}

func (v *Valr) SimpleSellQuote(currencyPair, payInCurrency string, amount decimal.Decimal) (quote *Quote, err error) {
	path := fmt.Sprintf("/simple/%s/quote", currencyPair)
	buy := simpleBuySell{payInCurrency, amount, SELL}

	body, err := structToBytes(buy)
	if err != nil {
		return
	}

	resp, err := v.client.do("POST", path, body, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &quote)
	return
}

type OrderID struct {
	ID string
}

func (v *Valr) SimpleBuyOrder(currencyPair, payInCurrency string, amount decimal.Decimal) (id *OrderID, err error) {
	path := fmt.Sprintf("/simple/%s/order", currencyPair)
	buy := simpleBuySell{payInCurrency, amount, BUY}

	body, err := structToBytes(buy)
	if err != nil {
		return
	}

	resp, err := v.client.do("POST", path, body, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &id)
	return
}

func (v *Valr) SimpleSellOrder(currencyPair, payInCurrency string, amount decimal.Decimal) (id *OrderID, err error) {
	path := fmt.Sprintf("/simple/%s/order", currencyPair)
	buy := simpleBuySell{payInCurrency, amount, SELL}

	body, err := structToBytes(buy)
	if err != nil {
		return
	}

	resp, err := v.client.do("POST", path, body, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &id)
	return
}
