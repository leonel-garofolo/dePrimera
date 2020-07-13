package daos

import (
	"database/sql"
	"deprimera/api/application"
	"deprimera/api/models"
	"log"
)

type AsistentesDaoImpl struct{}

func (ed *AsistentesDaoImpl) GetAll() []models.Asistentes {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from asistentes")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	asistentes := []models.Asistentes{}
	for rows.Next() {
		asistente := models.Asistentes{}
		error := rows.Scan(&asistente.IDAsistente, &asistente.IDPersona)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		asistentes = append(asistentes, asistente)
	}
	return asistentes
}

func (ed *AsistentesDaoImpl) Save(e *models.Asistentes) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	isDelete := ed.Delete(e.IDAsistente, e.IDPersona)
	if isDelete == true {
		_, error := db.Exec("insert into asistentes (id_asistente, id_persona) values(?,?)", e.IDAsistente, e.IDPersona)

		if error != nil {
			log.Println(error)
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
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return true
}
