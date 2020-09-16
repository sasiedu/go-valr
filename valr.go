package valr

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const (
	HttpBase   = "https://api.valr.com"
	ApiVersion = "v1"
	WsBase     = ""
)

type Valr struct {
	client *client
}

type OrderSide string

const (
	SELL OrderSide = "sell"
	BUY            = "buy"
)

// New returns an instantiated Valr struct
func New(apiKey, apiSecret string) *Valr {
	client := NewClient(apiKey, apiSecret)
	return &Valr{client}
}

// NewWithCustomHttpClient returns an instantiated Valr struct with custom http client
func NewWithCustomHttpClient(apiKey, apiSecret string, httpClient *http.Client) *Valr {
	client := NewClientWithCustomHttpConfig(apiKey, apiSecret, httpClient)
	return &Valr{client}
}

// NewWithCustomTimeout returns an instantiated Valr struct with custom timeout
func NewWithCustomTimeout(apiKey, apiSecret string, timeout time.Duration) *Valr {
	client := NewClientWithCustomTimeout(apiKey, apiSecret, timeout)
	return &Valr{client}
}

// set enable/disable http request/response dump
func (v *Valr) SetDebug(enable bool) {
	v.client.debug = enable
}

func (v *Valr) SetHttpBase(base string) {
	v.client.setHttpBase(base)
}

func (v *Valr) SetWsBase(base string) {
	v.client.setWsBase(base)
}
func (v *Valr) SetApiVersion(version string) {
	v.client.setApiVersion(version)
}

func structToBytes(val interface{}) ([]byte, error) {
	bytesBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(bytesBuffer).Encode(val); err != nil {
		return nil, err
	}

	return bytesBuffer.Bytes(), nil
}
