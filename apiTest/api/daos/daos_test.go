package daos

import (
	"log"
	"testing"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos"
)

func Test(t *testing.T) {
	myDao := daos.NewDePrimeraDaos()
	equiposDao := myDao.GetEquiposDao()
	equipos := equiposDao.GetAll()
	log.Println(len(equipos))
}
