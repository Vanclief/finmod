package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderRequestCalculateFields(t *testing.T) {
	type TestCase struct {
		Action           ActionType
		OrderType        OrderType
		Quantity         float64
		Price            float64
		Total            float64
		CurrentPrice     float64
		ExpectedError    bool
		ExpectedQuantity float64
		ExpectedTotal    float64
	}

	baseAsset, _ := NewCryptoAsset("ETH")
	quoteAsset, _ := NewForexAsset("USD")
	pair := NewPair(baseAsset, quoteAsset)

	or, err := NewOrderRequest(pair, BuyAction, LimitOrder, 5, 2000, 0)
	assert.Nil(t, err)

	// Case 1: Should not overwrite the original order request
	newOR, err := or.CalculateFields(0)
	assert.Nil(t, err)
	assert.NotEqual(t, or, newOR)

	// Case 2: Should work

	testCases := []TestCase{
		// Action   OrderType    Q  P  T   CP     EE    EQ  ET
		{BuyAction, MarketOrder, 10, 0, 0, 2000, false, 10, 20000},
		{BuyAction, MarketOrder, 0, 0, 10000, 2000, false, 5, 10000},
		{BuyAction, MarketOrder, 10, 10000, 0, 2000, false, 10, 20000},
		{BuyAction, LimitOrder, 10, 2000, 0, 0, false, 10, 20000},
		{BuyAction, LimitOrder, 0, 2000, 10000, 2000, false, 5, 10000},
		{SellAction, MarketOrder, 10, 0, 0, 2000, false, 10, 20000},
		{SellAction, MarketOrder, 0, 0, 10000, 2000, false, 5, 10000},
		{SellAction, MarketOrder, 10, 10000, 0, 2000, false, 10, 20000},
		{SellAction, LimitOrder, 10, 2000, 0, 0, false, 10, 20000},
		{SellAction, LimitOrder, 0, 2000, 10000, 2000, false, 5, 10000},
	}

	for _, testCase := range testCases {
		or, err = NewOrderRequest(pair, testCase.Action, testCase.OrderType, testCase.Quantity, testCase.Price, testCase.Total)
		assert.Nil(t, err)
		result, err := or.CalculateFields(testCase.CurrentPrice)
		if testCase.ExpectedError {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, testCase.ExpectedTotal, result.Total)
			assert.Equal(t, testCase.ExpectedQuantity, result.Quantity)
		}
	}
}
