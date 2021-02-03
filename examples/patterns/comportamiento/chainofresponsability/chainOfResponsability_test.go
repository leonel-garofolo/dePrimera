package main

import (
	"fmt"
	"testing"
)

func TestMensajeAltaPrioridad(t *testing.T) {
	manejadores := ReceptorBajaPrioridad{
		siguiente: ReceptorAltaPrioridad{},
	}

	fmt.Println(manejadores.ProcesarMensaje(4, "Mensaje 1 - Prioridad 4"))
	fmt.Println(manejadores.ProcesarMensaje(5, "Mensaje 2 - Prioridad 5"))
	fmt.Println(manejadores.ProcesarMensaje(10, "Mensaje 3 - Prioridad 10"))
}
