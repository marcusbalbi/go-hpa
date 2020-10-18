package balbi

import (
	"math"
	"fmt"
)

func ComplexCalc(value float64, times int) float64 {
	v := value

	if value < 0 || times < 0{
		return 2
	}

	fmt.Println("Iniciando Calculo com ", v , "Iterando ", times, " vezes")
	for i := 0; i < times; i++ {
		v = math.Sqrt(v)
		// fmt.Println("O valor na iteração ", i , " é ", v)
	}
	fmt.Println("Valor Final é ", v)
	fmt.Println("--------------------------------")
	return v
}
