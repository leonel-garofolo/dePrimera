package help

import (
	"fmt"
	"strconv"
)

type Partido struct {
	Local     int
	Visitante int
}

type FixtureHelp struct{}

func (ed *FixtureHelp) CalcularLiga(numEquipos int) [][]Partido {
	if numEquipos%2 == 0 {
		return calcularLigaNumEquiposPar(numEquipos)
	} else {
		return calcularLigaNumEquiposImpar(numEquipos)
	}
}

func calcularLigaNumEquiposPar(numEquipos int) [][]Partido {
	numRondas := numEquipos - 1
	numPartidosPorRonda := numEquipos / 2

	rondas := make([][]Partido, numRondas)
	k := 0
	for i := 0; i < numRondas; i++ {
		rondas[i] = make([]Partido, numPartidosPorRonda)
		for j := 0; j < numPartidosPorRonda; j++ {
			rondas[i][j] = Partido{}

			rondas[i][j].Local = k

			k++

			if k == numRondas {
				k = 0
			}
		}
	}

	for i := 0; i < numRondas; i++ {
		if i%2 == 0 {
			rondas[i][0].Visitante = numEquipos - 1
		} else {
			rondas[i][0].Visitante = rondas[i][0].Local
			rondas[i][0].Local = numEquipos - 1
		}
	}

	equipoMasAlto := numEquipos - 1
	equipoImparMasAlto := equipoMasAlto - 1

	k = equipoImparMasAlto
	for i := 0; i < numRondas; i++ {
		for j := 1; j < numPartidosPorRonda; j++ {
			rondas[i][j].Visitante = k

			k--

			if k == -1 {
				k = equipoImparMasAlto
			}
		}
	}

	return rondas
}

func calcularLigaNumEquiposImpar(numEquipos int) [][]Partido {
	numRondas := numEquipos
	numPartidosPorRonda := numEquipos / 2

	rondas := make([][]Partido, numRondas)
	k := 0
	for i := 0; i < numRondas; i++ {
		rondas[i] = make([]Partido, numPartidosPorRonda)
		for j := -1; j < numPartidosPorRonda; j++ {
			if j >= 0 {
				rondas[i][j] = Partido{}

				rondas[i][j].Local = k
			}

			k++

			if k == numRondas {
				k = 0
			}
		}
	}

	equipoMasAlto := numEquipos - 1
	k = equipoMasAlto
	for i := 0; i < numRondas; i++ {
		for j := 0; j < numPartidosPorRonda; j++ {
			rondas[i][j].Visitante = k

			k--

			if k == -1 {
				k = equipoMasAlto
			}
		}
	}

	return rondas
}

func (ed *FixtureHelp) MostrarPartidos(rondas [][]Partido) {
	fmt.Println("IDA")
	for i := 0; i < len(rondas); i++ {
		fmt.Print("Ronda " + strconv.Itoa((i + 1)) + ": ")
		for j := 0; j < len(rondas[i]); j++ {
			fmt.Print("   " + strconv.Itoa((1 + rondas[i][j].Local)) + "-" + strconv.Itoa((1 + rondas[i][j].Visitante)))
		}

		fmt.Println()
	}

	fmt.Println("VUELTA")

	for i := 0; i < len(rondas); i++ {
		fmt.Print("Ronda " + strconv.Itoa((i + 1)) + ": ")

		for j := 0; j < len(rondas[i]); j++ {
			fmt.Print("   " + strconv.Itoa((1 + rondas[i][j].Visitante)) + "-" + strconv.Itoa((1 + rondas[i][j].Local)))
		}

		fmt.Println()
	}

}
