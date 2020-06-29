// Interface
package strategy

type Estrategia interface {
	RealizarOperacion(int, int) int
}

type EstrategiaSuma struct{}

func (e EstrategiaSuma) RealizarOperacion(num1 int, num2 int) int {
	return num1 + num2
}

type EstrategiaResta struct{}

func (e EstrategiaResta) RealizarOperacion(num1 int, num2 int) int {
	return num1 - num2
}

type EstrategiaMultiplica struct{}

func (e EstrategiaMultiplica) RealizarOperacion(num1 int, num2 int) int {
	return num1 * num2
}

type EstrategiaDivicion struct{}

func (e EstrategiaDivicion) RealizarOperacion(num1 int, num2 int) int {
	return num1 / num2
}

type Contexto struct {
	estrategia Estrategia
}

func (c *Contexto) EjecutarOperacion(num1 int, num2 int) int {
	return c.estrategia.RealizarOperacion(num1, num2)
}
