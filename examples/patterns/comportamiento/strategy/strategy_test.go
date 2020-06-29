// Interface
package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	var contexto Contexto
	num1 := 10
	num2 := 5

	contexto = Contexto{EstrategiaSuma{}}
	fmt.Printf("%d + %d = %d\n", num1, num2, contexto.EjecutarOperacion(num1, num2))

	contexto = Contexto{EstrategiaResta{}}
	fmt.Printf("%d - %d = %d\n", num1, num2, contexto.EjecutarOperacion(num1, num2))

	contexto = Contexto{EstrategiaMultiplica{}}
	fmt.Printf("%d * %d = %d\n", num1, num2, contexto.EjecutarOperacion(num1, num2))

	contexto = Contexto{EstrategiaDivicion{}}
	fmt.Printf("%d / %d = %d\n", num1, num2, contexto.EjecutarOperacion(num1, num2))
}
