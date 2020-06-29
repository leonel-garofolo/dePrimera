package factory

import (
	"fmt"
	"testing"
)

func TestMensajeAltaPrioridad(t *testing.T) {
	fabricaPuertaMadera := &FabricaPuertaMadera{}
	puertaMadera := fabricaPuertaMadera.ConstruirPuerta()
	fmt.Printf("Se construyo un puerta de: %s\n", puertaMadera.VerMaterial())

	fabricaPuertaMetal := &FabricaPuertaMetal{}
	puertaMetal := fabricaPuertaMetal.ConstruirPuerta()
	fmt.Printf("Se construyo un puerta de: %s\n", puertaMetal.VerMaterial())
}
