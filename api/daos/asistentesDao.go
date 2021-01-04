package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
)

// AsistentesDaoImpl struct
type AsistentesDaoImpl struct{}

// GetAll asistentes
func (ed *AsistentesDaoImpl) GetAll() []gorms.AsistentesGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from asistentes")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	asistentes := []gorms.AsistentesGorm{}
	for rows.Next() {
		asistente := gorms.AsistentesGorm{}
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
func (ed *AsistentesDaoImpl) Save(e *gorms.AsistentesGorm) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	isDelete := true
	if e.IDAsistente > 0 {
		isDelete = ed.Delete(e.IDAsistente, e.IDPersona, e.IDCampeonato)
	}

	if isDelete == true {
		_, error := db.Exec("insert into asistentes (id_asistente, id_persona, id_campeonato) values(?,?, ?)",
			e.IDAsistente, e.IDPersona, e.IDCampeonato)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	}
	return e.IDAsistente
}

// Delete asistentes
func (ed *AsistentesDaoImpl) Delete(IDAsistente int64, IDPersona int64, IDCampeonato int64) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from asistentes where id_asistente = ? and id_persona = ? and id_campeonato", IDAsistente, IDPersona, IDCampeonato)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return true
}
