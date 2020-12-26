package fixture

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println("Liga con 8 equipos:")
	fixtureService := FixtureService{}
	mostrarPartidos(fixtureService.calcularLiga(8))

	fmt.Println()

	fmt.Println("Liga con 7 equipos:")

	mostrarPartidos(fixtureService.calcularLiga(7))
}
