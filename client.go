package valr

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
	"time"
)

func signRequest(apiSecret, method, path, body string, timestamp time.Time) (string, string) {
	// Create a new Keyed-Hash Message Authentication Code (HMAC) using SHA512 and API Secret
	mac := hmac.New(sha512.New, []byte(apiSecret))
	// Convert timestamp to nanoseconds then divide by 1000000 to get the milliseconds
	timestampString := strconv.FormatInt(timestamp.UnixNano()/1000000, 10)

	mac.Write([]byte(timestampString))
	mac.Write([]byte(strings.ToUpper(method)))
	mac.Write([]byte(path))
	mac.Write([]byte(body))
	// Gets the byte hash from HMAC and converts it into a hex string
	return hex.EncodeToString(mac.Sum(nil)), timestampString
}

type client struct {
	apiKey      string
	apiSecret   string
	httpClient  *http.Client
	httpTimeout time.Duration
	debug       bool
	httpBase    string
	apiVersion  string
	wsBase      string
}

func NewClient(apiKey, apiSecret string) (c *client) {
	return &client{apiKey, apiSecret, &http.Client{}, 30 * time.Second, false, HttpBase, ApiVersion, WsBase}
}

func NewClientWithCustomHttpConfig(apiKey, apiSecret string, httpClient *http.Client) (c *client) {
	timeout := httpClient.Timeout
	if timeout <= 0 {
		timeout = 30 * time.Second
	}
	return &client{apiKey, apiSecret, httpClient, timeout, false, HttpBase, ApiVersion, WsBase}
}

func NewClientWithCustomTimeout(apiKey, apiSecret string, timeout time.Duration) (c *client) {
	return &client{apiKey, apiSecret, &http.Client{}, timeout, false, HttpBase, ApiVersion, WsBase}
}

func (c client) dumpRequest(r *http.Request) {
	if r == nil {
		log.Print("dumpReq ok: <nil>")
		return
	}
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Print("dumpReq err:", err)
	} else {
		log.Print("dumpReq ok:", string(dump))
	}
}

func (c client) dumpResponse(r *http.Response) {
	if r == nil {
		log.Print("dumpResponse ok: <nil>")
		return
	}
	dump, err := httputil.DumpResponse(r, true)
	if err != nil {
		log.Print("dumpResponse err:", err)
	} else {
		log.Print("dumpResponse ok:", string(dump))
	}
}

func (c *client) setHttpBase(base string) {
	c.httpBase = base
}

func (c *client) setApiVersion(version string) {
	c.apiVersion = version
}

func (c *client) setWsBase(base string) {
	c.wsBase = base
}

// doTimeoutRequest do a HTTP request with timeout
func (c *client) doTimeoutRequest(timer *time.Timer, req *http.Request) (*http.Response, error) {
	// Do the request in the background so we can check the timeout
	type result struct {
		resp *http.Response
		err  error
	}
	done := make(chan result, 1)
	go func() {
		if c.debug {
			c.dumpRequest(req)
		}
		resp, err := c.httpClient.Do(req)
		if c.debug {
			c.dumpResponse(resp)
		}
		done <- result{resp, err}
	}()
	// Wait for the read or the timeout
	select {
	case r := <-done:
		return r.resp, r.err
	case <-timer.C:
		return nil, errors.New("timeout on reading data from Valr API")
	}
}

func (c *client) do(method, path string, data []byte, authNeeded bool) (response []byte, err error) {
	connectTimer := time.NewTimer(c.httpTimeout)

	url := fmt.Sprintf("%s/%s/%s", c.httpBase, c.apiVersion, strings.Trim(path, "/"))

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	if authNeeded {
		signaturePath := fmt.Sprintf("/%s/%s", c.apiVersion, strings.Trim(path, "/"))
		currentTime := time.Now()
		signature, timestamp := signRequest(c.apiSecret,
			method,
			signaturePath,
			string(data),
			currentTime)

		req.Header.Add("X-VALR-API-KEY", c.apiKey)
		req.Header.Add("X-VALR-SIGNATURE", signature)
		req.Header.Add("X-VALR-TIMESTAMP", timestamp)
	}

	resp, err := c.doTimeoutRequest(connectTimer, req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 202 && resp.StatusCode != 203 {
		err = errors.New(resp.Status + ": " + string(response))
	}
	return response, err
}
