package daos

import (
	"database/sql"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	"log"
)

// ArbitrosDaoImpl struct
type ArbitrosDaoImpl struct{}

// GetAll arbritros
func (ed *ArbitrosDaoImpl) GetAll() []gorms.ArbitrosGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from arbitros")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	arbitros := []gorms.ArbitrosGorm{}
	for rows.Next() {
		arbitro := gorms.ArbitrosGorm{}
		error := rows.Scan(&arbitro.IDArbitro, &arbitro.IDPersona)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		arbitros = append(arbitros, arbitro)
	}
	return arbitros
}

// Save arbritros
func (ed *ArbitrosDaoImpl) Save(e *gorms.ArbitrosGorm) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	isDelete := ed.Delete(e.IDArbitro, e.IDPersona)
	if isDelete == true {
		_, error := db.Exec("insert into arbitros (id_arbitro, id_persona) values(?,?)", e.IDArbitro, e.IDPersona)
		if error != nil {
			log.Println(error)
			panic(error)
		}
	}
	return e.IDArbitro
}

// Delete arbitro
func (ed *ArbitrosDaoImpl) Delete(IDArbitro int64, IDPersona int64) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from arbitros where id_arbitro = ? and id_persona = ?", IDArbitro, IDPersona)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return true
}
