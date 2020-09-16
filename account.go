package valr

import (
	"encoding/json"
	"fmt"
	"time"
)

type Balance struct {
	Currency  string
	Available float64 `json:",string"`
	Reserved  float64 `json:",string"`
	Total     float64 `json:",string"`
}

func (v *Valr) GetBalance() (balances []Balance, err error) {
	resp, err := v.client.do("GET", "/account/balances", []byte(""), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &balances)
	return
}

type TransactionType struct {
	Type        string
	Description string
}

type Transaction struct {
	TransactionType TransactionType
	DebitCurrency   string
	DebitValue      float64 `json:",string"`
	CreditCurrency  string
	CreditValue     float64 `json:",string"`
	FeeCurrency     string
	FeeValue        float64 `json:",string"`
	EventAt         time.Time
	AdditionalInfo  struct {
		CostPerCoin        float64
		CostPerCoinSymbol  string
		CurrencyPairSymbol string
		OrderId            string
	}
	ID string
}

type TransactionFilter struct {
	Skip            string
	Limit           string
	TransactionType string
	Currency        string
	StartTime       string
	EndTime         string
}

func NewTransactionFilter() TransactionFilter {
	return TransactionFilter{
		Skip:            "",
		Limit:           "",
		TransactionType: "",
		Currency:        "",
		StartTime:       "",
		EndTime:         "",
	}
}

func (v *Valr) GetTransactionHistory() (history []Transaction, err error) {
	path := "/account/transactionhistory?skip=0&limit=100"
	resp, err := v.client.do("GET", path, []byte(""), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &history)
	return
}

func (v *Valr) GetTransactionHistorySkipAndLimit(skip uint32, limit uint32) (history []Transaction, err error) {
	path := fmt.Sprintf("/account/transactionhistory?skip=%d&limit=%d", skip, limit)
	resp, err := v.client.do("GET", path, []byte(""), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &history)
	return
}

func addParams(params, name, value string) string {
	if value == "" {
		return params
	}
	if params == "" {
		params = fmt.Sprintf("?%s=%s", name, value)
	} else {
		params = fmt.Sprintf("%s&%s=%s", params, name, value)
	}
	return params
}

func (v *Valr) GetTransactionHistoryFiltered(filter *TransactionFilter) (history []Transaction, err error) {
	params := ""
	if filter != nil {
		params = addParams("", "skip", filter.Skip)
		params = addParams(params, "limit", filter.Limit)
		params = addParams(params, "transactionTypes", filter.TransactionType)
		params = addParams(params, "currency", filter.Currency)
		params = addParams(params, "startTime", filter.StartTime)
		params = addParams(params, "endTime", filter.EndTime)
	}

	path := "/account/transactionhistory" + params

	resp, err := v.client.do("GET", path, []byte(""), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &history)
	return
}

func (v *Valr) GetTransactionHistoryLimitById(limit uint32, id string) (history []Transaction, err error) {
	path := fmt.Sprintf("/account/transactionhistory?limit=%d&beforeId=%s", limit, id)
	resp, err := v.client.do("GET", path, []byte(""), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &history)
	return
}

func (v *Valr) GetTransactionHistoryForCurrencyPair(pair string, limit uint32) (history []Transaction, err error) {
	path := fmt.Sprintf("/account/%s/tradehistory?limit=%d", pair, limit)
	resp, err := v.client.do("GET", path, []byte(""), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &history)
	return
}
