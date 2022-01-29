package indicators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoriaChannel(t *testing.T) {
	length := 100

	candles, _, _, _, _, err := loadCandlesFromFile("./test_dataset/BINANCE_ETHUSD_60.csv")
	assert.Nil(t, err)
	ans := NoriaChannel(candles, length)
	for i := 0; i < 10; i++ {
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

// TODO
// 1. Calcular el punto medio entre el low y el high
// 2. Sacar la regresion lineal de los puntos medios de las candles
// 3. Rotar todos los puntos por el angulo de la regresion lineal
// 4. Obtener la regresion lineal del low y el high con los puntos rotados
// 5. Hacer el offset de la regresion lineal para cubrir todos los puntos
// 6. Rotar las lineas y los puntos en el otro sentido
