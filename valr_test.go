package valr

import (
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestValrHttpPublicApi(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	httpBase := os.Getenv("HTTP_BASE")

	valr := New("", "")
	if httpBase != "" {
		valr.SetHttpBase(httpBase)
	}

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

func TestValrHttpAccountApi(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	apiKey := os.Getenv("VALR_API_KEY")
	apiSecret := os.Getenv("VALR_API_SECRET")
	httpBase := os.Getenv("HTTP_BASE")

	valr := New(apiKey, apiSecret)
	if httpBase != "" {
		valr.SetHttpBase(httpBase)
	}

	balances, err := valr.GetBalance()
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(balances), 1)

	transactionHistory, err := valr.GetTransactionHistory()
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(transactionHistory), 1)

	transactionHistorySkipAndLimit, err := valr.GetTransactionHistorySkipAndLimit(1, 5)
	assert.Nil(t, err)
	assert.Equal(t, len(transactionHistorySkipAndLimit), 5)
	assert.Equal(t, transactionHistory[1], transactionHistorySkipAndLimit[0])

	transactionHistory2, err := valr.GetTransactionHistoryFiltered(
		&TransactionFilter{"", "", "", "BTC", "", ""},
	)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(transactionHistory2), 1)

	transactionHistory3, err := valr.GetTransactionHistoryFiltered(
		&TransactionFilter{"", "", "LIMIT_BUY", "BTC", "", ""},
	)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(transactionHistory3), 1)

	transactionHistoryLimitById, err := valr.GetTransactionHistoryLimitById(
		2,
		transactionHistory[0].ID,
	)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(transactionHistoryLimitById))

	transactionHistoryCurrencyPair, err := valr.GetTransactionHistoryForCurrencyPair("BTCZAR", 2)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(transactionHistoryCurrencyPair))
}

func TestValrHttpWalletApi(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	apiKey := os.Getenv("VALR_API_KEY")
	apiSecret := os.Getenv("VALR_API_SECRET")
	httpBase := os.Getenv("HTTP_BASE")

	valr := New(apiKey, apiSecret)
	if httpBase != "" {
		valr.SetHttpBase(httpBase)
	}

	depositAddress, err := valr.GetDepositAddress("XRP")
	assert.Nil(t, err)
	assert.NotNil(t, depositAddress)
	assert.Equal(t, "XRP", depositAddress.Currency)

	withdrawalInfo, err := valr.GetCurrencyWithdrawalInfo("XRP")
	assert.Nil(t, err)
	assert.NotNil(t, withdrawalInfo)
	assert.Equal(t, "XRP", withdrawalInfo.Currency)
	assert.Equal(t, true, withdrawalInfo.IsActive)
	assert.Equal(t, false, withdrawalInfo.SupportPaymentReference)

	withdrawID, err := valr.NewCryptoWithdrawal("XRP", depositAddress.Address, withdrawalInfo.MinimumWithdrawAmount, depositAddress.PaymentReference)
	assert.Nil(t, err)
	assert.NotNil(t, withdrawID)
	assert.NotEqual(t, "", withdrawID.ID)

	status, err := valr.GetCryptoWithdrawalStatus("XRP", withdrawID.ID)
	assert.Nil(t, err)
	assert.NotNil(t, status)
	assert.Equal(t, "XRP", status.Currency)
	assert.Equal(t, depositAddress.Address, status.Address)
	assert.Equal(t, withdrawalInfo.MinimumWithdrawAmount, status.Amount)

	depositHistory, err := valr.GetCryptoDepositHistory("BTC", 0, 2)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(depositHistory), 1)
	assert.Equal(t, "BTC", depositHistory[0].CurrencyCode)

	withdrawHistory, err := valr.GetCryptoWithdrawalHistory("XRP", 0, 2)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(withdrawHistory), 1)
	assert.Equal(t, "XRP", withdrawHistory[0].Currency)
	assert.Equal(t, depositAddress.Address, withdrawHistory[0].Address)
	assert.Equal(t, withdrawalInfo.MinimumWithdrawAmount, withdrawHistory[0].Amount)

	accounts, err := valr.GetBankAccounts()
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(accounts), 1)

	fiatWithdraw, err := valr.NewFiatWithdrawal(accounts[0].ID, decimal.NewFromInt(1), false)
	assert.Nil(t, err)
	assert.NotNil(t, fiatWithdraw)
	assert.NotEqual(t, "", fiatWithdraw.ID)
}

func TestValrHttpMarketApi(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	apiKey := os.Getenv("VALR_API_KEY")
	apiSecret := os.Getenv("VALR_API_SECRET")
	httpBase := os.Getenv("HTTP_BASE")

	valr := New(apiKey, apiSecret)
	if httpBase != "" {
		valr.SetHttpBase(httpBase)
	}

	orderBook, err := valr.GetOrderBook("BTCZAR")
	assert.Nil(t, err)
	assert.NotNil(t, orderBook)
	assert.GreaterOrEqual(t, len(orderBook.Asks), 1)
	assert.GreaterOrEqual(t, len(orderBook.Bids), 1)
	assert.Equal(t, "BTCZAR", orderBook.Bids[0].CurrencyPair)
	assert.Equal(t, "", orderBook.Bids[0].ID)
	assert.Equal(t, "BTCZAR", orderBook.Asks[0].CurrencyPair)

	orderBookNonAggregated, err := valr.GetNonAggregatedOrderBook("BTCZAR")
	assert.Nil(t, err)
	assert.NotNil(t, orderBookNonAggregated)
	assert.GreaterOrEqual(t, len(orderBookNonAggregated.Asks), 1)
	assert.GreaterOrEqual(t, len(orderBookNonAggregated.Bids), 1)
	assert.Equal(t, "BTCZAR", orderBookNonAggregated.Bids[0].CurrencyPair)
	assert.NotEqual(t, "", orderBookNonAggregated.Bids[0].ID)
	assert.Equal(t, "BTCZAR", orderBookNonAggregated.Asks[0].CurrencyPair)

	tradeHistory, err := valr.GetCurrencyPairTradeHistory("BTCZAR", 10)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(tradeHistory), 1)
	assert.Equal(t, "BTCZAR", tradeHistory[0].CurrencyPair)
	assert.Equal(t, false, tradeHistory[0].Price.IsZero())
}

func TestValrHttpSimpleApi(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	apiKey := os.Getenv("VALR_API_KEY")
	apiSecret := os.Getenv("VALR_API_SECRET")
	httpBase := os.Getenv("HTTP_BASE")

	valr := New(apiKey, apiSecret)
	if httpBase != "" {
		valr.SetHttpBase(httpBase)
	}

	buyQuote, err := valr.SimpleBuyQuote("BTCZAR", "ZAR", decimal.NewFromInt32(10))
	assert.Nil(t, err)
	assert.NotNil(t, buyQuote)
	assert.Equal(t, "BTCZAR", buyQuote.CurrencyPair)

	sellQuote, err := valr.SimpleSellQuote("BTCZAR", "BTC", decimal.NewFromFloat(0.0001))
	assert.Nil(t, err)
	assert.NotNil(t, sellQuote)
	assert.Equal(t, "BTCZAR", sellQuote.CurrencyPair)

	buyOrder, err := valr.SimpleBuyOrder("XRPZAR", "ZAR", decimal.NewFromInt32(10))
	assert.Nil(t, err)
	assert.NotNil(t, buyOrder)
	assert.NotEqual(t, "", buyOrder.ID)

	sellOrder, err := valr.SimpleSellOrder("XRPZAR", "XRP", decimal.NewFromFloat(3))
	assert.Nil(t, err)
	assert.NotNil(t, sellOrder)
	assert.NotEmpty(t, sellOrder.ID)
}
