package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"log"
)

type ZonasEquiposDaoImpl struct{}

func (ed *ZonasEquiposDaoImpl) Save(e *models.ZonasEquipos) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	isDelete := ed.Delete(e.IDZona, e.IDEquipo)
	if isDelete == true {
		_, error := db.Exec("insert into zonas_equipos (id_zonas, id_equipo) values(?,?)", e.IDZona, e.IDEquipo)

		if error != nil {
			panic(error)
		}
	}
	return e.IDEquipo
}

func (ed *ZonasEquiposDaoImpl) Delete(IDZona int64, IDEquipo int64) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from zonas_equipos where id_zonas = ? and id_equipo = ?", IDZona, IDEquipo)
	if error != nil {
		panic(error)
	}
	return true
}
