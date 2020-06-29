package daos

import (
	"deprimera/api/daos"
	"log"
	"testing"
)

func Test(t *testing.T) {
	myDao := daos.NewDePrimeraDaos()
	equiposDao := myDao.GetEquiposDao()
	equipos := equiposDao.GetAll()
	log.Println(len(equipos))
}
