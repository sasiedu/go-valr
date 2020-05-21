package valr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValrHttpPublicApi(t *testing.T) {
	valr := New("", "")

	currencies, err := valr.GetCurrencies()
	assert.Nil(t, err)
	assert.NotNil(t, currencies)
	assert.GreaterOrEqual(t, len(currencies), 1)

	orderBook, err := valr.GetPublicOrderBook("BTCZAR")
	assert.Nil(t, err)
	assert.NotNil(t, orderBook)
	assert.GreaterOrEqual(t, len(orderBook.Asks), 1)
	assert.GreaterOrEqual(t, len(orderBook.Bids), 1)
	assert.Equal(t, orderBook.Bids[0].CurrencyPair, "BTCZAR")
	assert.Equal(t, orderBook.Asks[0].CurrencyPair, "BTCZAR")

	currencyPairs, err := valr.GetCurrencyPairs()
	assert.Nil(t, err)
	assert.NotNil(t, currencyPairs)
	assert.GreaterOrEqual(t, len(currencyPairs), 1)

	currencyPairsOrderTypes, err := valr.GetAllCurrencyPairOrderTypes()
	assert.Nil(t, err)
	assert.NotNil(t, currencyPairsOrderTypes)
	assert.GreaterOrEqual(t, len(currencyPairsOrderTypes), 1)
	assert.GreaterOrEqual(t, len(currencyPairsOrderTypes[0].OrderTypes), 1)

	orderTypes, err := valr.GetOrderTypesForCurrencyPair("BTCZAR")
	assert.Nil(t, err)
	assert.NotNil(t, orderTypes)
	assert.GreaterOrEqual(t, len(orderTypes), 1)

	marketSummaries, err := valr.GetAllCurrencyPairMarketSummary()
	assert.Nil(t, err)
	assert.NotNil(t, marketSummaries)
	assert.GreaterOrEqual(t, len(marketSummaries), 1)

	marketSummary, err := valr.GetMarketSummaryForCurrencyPair("BTCZAR")
	assert.Nil(t, err)
	assert.NotNil(t, marketSummary)
	assert.Equal(t, marketSummary.CurrencyPair, "BTCZAR")

	marketSummaryInvalid, err := valr.GetMarketSummaryForCurrencyPair("BTCZA")
	assert.NotNil(t, err)
	assert.Nil(t, marketSummaryInvalid)

	serverTime, err := valr.GetServerTime()
	assert.Nil(t, err)
	assert.NotNil(t, serverTime)
}
