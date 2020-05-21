package valr

import (
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
