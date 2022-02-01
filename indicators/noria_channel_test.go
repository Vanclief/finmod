package indicators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoriaChannel(t *testing.T) {
	length := 100

	candles, _, _, _, _, err := loadCandlesFromFile("./test_dataset/BINANCE_ETHUSD_60.csv")
	assert.Nil(t, err)
	ans := NoriaChannel(candles[:length], length)
	for i := 0; i < len(ans); i++ {
		ans[i].Print()
	}

	//zeros := FindProperty(ans, 10)
	//for _, v := range zeros {
	//	fmt.Println(len(v))
	//	for _, vv := range v {
	//		vv.Print()
	//	}
	//	fmt.Println("------------------")
	//}
}
