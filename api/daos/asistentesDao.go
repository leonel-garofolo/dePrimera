package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"log"
)

type AsistentesDaoImpl struct{}

func (ed *AsistentesDaoImpl) Save(e *models.Asistentes) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	isDelete := ed.Delete(e.IDAsistente, e.IDPersona)
	if isDelete == true {
		_, error := db.Exec("insert into arbitros (id_arbitros, id_personas) values(?,?)", e.IDAsistente, e.IDPersona)

		if error != nil {
			panic(error)
		}
	}
	return e.IDAsistente
}

func (ed *AsistentesDaoImpl) Delete(IDAsistente int64, IDPersona int64) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from asistentes where id_asistente = ? and id_persona = ?", IDAsistente, IDPersona)
	if error != nil {
		panic(error)
	}
	return true
}
