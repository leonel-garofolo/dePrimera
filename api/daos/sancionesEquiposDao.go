package daos

import (
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
)

type SancionesEquiposDaoImpl struct{}

func (ed *SancionesEquiposDaoImpl) Save(e *gorms.SancionesEquiposGorm) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	isDelete := ed.Delete(e.IDEquipo, e.IDSanciones, e.IDCampeonato)
	if isDelete == true {
		_, error := db.Exec("insert into sanciones_equipos (id_equipos, id_sanciones, id_campeonato) values($1,$2, $3)",
			e.IDEquipo, e.IDSanciones, e.IDCampeonato)

		if error != nil {
			panic(error)
		}
	}
	return e.IDEquipo
}

func (ed *SancionesEquiposDaoImpl) Delete(IDEquipo int64, IDSanciones int64, IDCampeonato int64) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from sanciones_equipos where id_equipos = $1 and id_sanciones = $2",
		IDEquipo, IDSanciones, IDCampeonato)
	if error != nil {
		panic(error)
	}
	return true
}
