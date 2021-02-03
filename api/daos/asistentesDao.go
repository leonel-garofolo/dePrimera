package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
)

// AsistentesDaoImpl struct
type AsistentesDaoImpl struct{}

// GetAll asistentes
func (ed *AsistentesDaoImpl) GetAll() []models.Asistentes {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from asistentes")
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}

	asistentes := []models.Asistentes{}
	for rows.Next() {
		asistente := models.Asistentes{}
		error := rows.Scan(&asistente.IDAsistente, &asistente.IDPersona, &asistente.IDCampeonato)
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

// Save asistentes
func (ed *AsistentesDaoImpl) Save(e *models.Asistentes) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	isDelete := true
	if e.IDAsistente > 0 {
		isDelete, _ = ed.Delete(e.IDAsistente, e.IDPersona, e.IDCampeonato)
	}

	if isDelete == true {
		_, error := db.Exec("insert into asistentes (id_asistente, id_persona, id_campeonato) values($1,$2, $3)",
			e.IDAsistente, e.IDPersona, e.IDCampeonato)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	}
	return e.IDAsistente
}

// Delete asistentes
func (ed *AsistentesDaoImpl) Delete(IDAsistente int64, IDPersona int64, IDCampeonato int64) (bool, error) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from asistentes where id_asistente = $1 and id_persona = $2 and id_campeonato = $3", IDAsistente, IDPersona, IDCampeonato)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			return false, error
		}
	}
	return true, nil
}
