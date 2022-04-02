package market

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairSymbol(t *testing.T) {

	base := Asset{Symbol: "ETH"}
	quote := Asset{Symbol: "USD"}

	pair := Pair{Base: base, Quote: quote}

	baseAlt := Asset{Symbol: "#US30"}
	pairAlt := Pair{Base: baseAlt}

	assert.Equal(t, "ETHUSD", pair.Symbol(""))
	assert.Equal(t, "ETH/USD", pair.Symbol("/"))
	assert.Equal(t, "#US30", pairAlt.Symbol(""))
	assert.Equal(t, "#US30", pairAlt.Symbol("/"))
}

func TestCreateMapping(t *testing.T) {

	resp, _ := http.Get("https://api.binance.com/api/v3/exchangeInfo")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	type Symbol struct {
		Symbol     string `json:"symbol"`
		BaseAsset  string `json:"baseAsset"`
		QuoteAsset string `json:"quoteAsset"`
	}

	type Response struct {
		Symbols []Symbol `json:"symbols"`
	}

	result := Response{}

	err = json.Unmarshal(body, &result)
	assert.Nil(t, err)

	// for _, symbol := range result.Symbols {
	// mapping := fmt.Sprintf(`"%s": {Base: Asset{Symbol: "%s"}, Quote: Asset{Symbol: "%s"}},`,
	// symbol.Symbol, symbol.BaseAsset, symbol.QuoteAsset)
	// fmt.Println(mapping)
	// }

}
