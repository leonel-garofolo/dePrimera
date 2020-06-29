package daos

import (
	"deprimera/api/daos"
	"log"
	"testing"
)

func Test(t *testing.T) {
	myDao := &daos.DePrimeraDaos{}
	equipos := myDao.GetEquiposDao().GetAll()
	log.Println(len(equipos))
}
