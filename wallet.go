package valr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

type DepositAddress struct {
	Currency             string
	Address              string
	PaymentReference     string
	PaymentReferenceName string
}

func (v *Valr) GetDepositAddress(currencyCode string) (address *DepositAddress, err error) {
	path := fmt.Sprintf("/wallet/crypto/%s/deposit/address", currencyCode)
	resp, err := v.client.do("GET", path, []byte(""), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &address)
	return
}

type CurrencyInfo struct {
	Currency                string
	MinimumWithdrawAmount   decimal.Decimal
	IsActive                bool
	WithdrawCost            decimal.Decimal
	SupportPaymentReference bool
}

func (v *Valr) GetCurrencyWithdrawalInfo(currencyCode string) (info *CurrencyInfo, err error) {
	path := fmt.Sprintf("/wallet/crypto/%s/withdraw", currencyCode)
	resp, err := v.client.do("GET", path, []byte(""), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &info)
	return
}

type newWithdrawal struct {
	Amount           decimal.Decimal `json:"amount"`
	Address          string          `json:"address"`
	PaymentReference string          `json:"paymentReference"`
}

func structToBytes(val interface{}) ([]byte, error) {
	bytesBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(bytesBuffer).Encode(val); err != nil {
		return nil, err
	}

	return bytesBuffer.Bytes(), nil
}

type WithdrawalID struct {
	ID string
}

func (v *Valr) NewCryptoWithdrawal(currency, address string, amount decimal.Decimal, paymentReference string) (id *WithdrawalID, err error) {
	path := fmt.Sprintf("/wallet/crypto/%s/withdraw", currency)
	withdraw := newWithdrawal{amount, address, paymentReference}

	body, err := structToBytes(withdraw)
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

type WithdrawalStatus struct {
	Currency           string
	Address            string
	Amount             decimal.Decimal
	FeeAmount          decimal.Decimal
	TransactionHash    string
	Confirmations      uint8
	LastConfirmationAt string
	UniqueID           string
	CreatedAt          string
	Verified           bool
	Status             string
}

func (v *Valr) GetCryptoWithdrawalStatus(currency, WithdrawalID string) (status *WithdrawalStatus, err error) {
	path := fmt.Sprintf("/wallet/crypto/%s/withdraw/%s", currency, WithdrawalID)
	resp, err := v.client.do("GET", path, []byte(""), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &status)
	return
}

type Deposit struct {
	CurrencyCode    string
	ReceiveAddress  string
	TransactionHash string
	Amount          decimal.Decimal
	CreatedAt       time.Time
	Confirmations   uint8
	Confirmed       bool
	ConfirmedAt     string
}

func (v *Valr) GetCryptoDepositHistory(currency string, skip, limit uint32) (history []Deposit, err error) {
	path := fmt.Sprintf("/wallet/crypto/%s/deposit/history?skip=%d&limit=%d", currency, skip, limit)
	resp, err := v.client.do("GET", path, []byte(""), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &history)
	return
}

type Withdrawal struct {
	Currency           string
	Address            string
	Amount             decimal.Decimal
	FeeAmount          decimal.Decimal
	TransactionHash    string
	Confirmations      uint8
	LastConfirmationAt string
	UniqueID           string
	CreatedAt          string
	Verified           bool
	Status             string
}

func (v *Valr) GetCryptoWithdrawalHistory(currency string, skip, limit uint32) (history []Withdrawal, err error) {
	path := fmt.Sprintf("/wallet/crypto/%s/withdraw/history?skip=%d&limit=%d", currency, skip, limit)
	resp, err := v.client.do("GET", path, []byte(""), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &history)
	return
}

type BankAccount struct {
	ID            string
	Bank          string
	AccountHolder string
	AccountNumber string
	BranchCode    string
	AccountType   string
	CreatedAt     string
}

func (v *Valr) GetBankAccounts() (banks []BankAccount, err error) {
	path := "/wallet/fiat/ZAR/accounts"
	resp, err := v.client.do("GET", path, []byte(""), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &banks)
	return
}

type fiatWithdraw struct {
	LinkedBankAccountId string          `json:"linkedBankAccountId"`
	Amount              decimal.Decimal `json:"amount"`
	Fast                bool            `json:"fast"`
}

func (v *Valr) NewFiatWithdrawal(bankAccountId string, amount decimal.Decimal, fastWithdraw bool) (id *WithdrawalID, err error) {
	path := "/wallet/fiat/ZAR/withdraw"
	withdraw := fiatWithdraw{bankAccountId, amount, fastWithdraw}

	body, err := structToBytes(withdraw)
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
