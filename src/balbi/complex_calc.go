package balbi

import (
	"math"
	"fmt"
)

func ComplexCalc(value float64, times int) float64 {
	v := value
	fmt.Println("Iniciando Cálculo com ", v , "Iterando ", times, " vezes")
	for i := 0; i < times; i++ {
		v = math.Sqrt(v)
		// fmt.Println("O valor na iteração ", i , " é ", v)
	}
	fmt.Println("Valor Final é ", v)
	fmt.Println("--------------------------------")
	return v
}
