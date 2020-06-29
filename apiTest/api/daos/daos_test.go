package daos

import (
	"deprimera/api/daos"
	"log"
	"testing"
)

func Test(t *testing.T) {
	equipos := daos.GetEquiposDao().GetAll()
	log.Println(len(equipos))
}
