package valr

import (
	"encoding/json"
	"fmt"
)

type LimitOrder struct {
	Side            OrderSide `json:"side"`
	Quantity        float64   `json:"quantity"`
	Price           float64   `json:"price"`
	PostOnly        bool      `json:"postOnly"`
	CustomerOrderID string    `json:"customerOrderId"`
}

type MarketOrder struct {
	Side            OrderSide `json:"side"`
	BaseAmount      float64   `json:"baseAmount"`
	QuoteAmount     float64   `json:"quoteAmount"`
	Pair            string    `json:"pair"`
	CustomerOrderID string    `json:"customerOrderId"`
}

type OrderStatus struct {
	OrderID           string
	OrderStatusType   string
	CurrencyPair      string
	OriginalPrice     float64
	RemainingQuantity float64
	OriginalQuantity  float64
	FilledPercentage  float64
	OrderSide         OrderSide
	OrderType         string
	FailedReason      string
	CustomerOrderID   string
	OrderUpdatedAt    string
	OrderCreatedAt    string
}

func (v *Valr) LimitOrder(order LimitOrder) (id *OrderID, err error) {
	path := "/orders/limit"

	body, err := structToBytes(order)
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

func (v *Valr) MarketOrder(order MarketOrder) (id *OrderID, err error) {
	path := "/orders/market"

	body, err := structToBytes(order)
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

func (v *Valr) GetOrderStatus(currencyPair, id string) (status *OrderStatus, err error) {
	path := fmt.Sprintf("/orders/%s/orderid/%s", currencyPair, id)
	resp, err := v.client.do("GET", path, []byte(""), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &status)
	return
}
