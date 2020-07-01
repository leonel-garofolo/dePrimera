package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"log"
)

type EquiposJugadoresDaoImpl struct{}

func (ed *EquiposJugadoresDaoImpl) Save(e *models.EquiposJugadores) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	isDelete := ed.Delete(e.IDEquipos, e.IDJugadores)
	if isDelete == true {
		_, error := db.Exec("insert into equipos_jugadores (id_equipos, id_jugadores) values(?,?)", e.IDEquipos, e.IDJugadores)

		if error != nil {
			panic(error)
		}
	}
	return e.IDEquipos
}

func (ed *EquiposJugadoresDaoImpl) Delete(IDEquipos int64, IDJugadores int64) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from equipos_jugadores where id_equipos = ? and id_jugadores = ?", IDEquipos, IDJugadores)
	if error != nil {
		panic(error)
	}
	return true
}
