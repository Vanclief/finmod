package indicators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoriaChannel(t *testing.T) {
	candles, _, _, _, _, err := loadCandlesFromFile("./test_dataset/BINANCE_ETHUSD_60.csv")
	assert.Nil(t, err)
	ans := Iterator(candles)
	for _, v := range ans {
		v.Print()
	}
}
