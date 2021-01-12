package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
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
		error := rows.Scan(&arbitro.IDArbitro, &arbitro.IDPersona, &arbitro.IDCampeonato)
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

func (ed *ArbitrosDaoImpl) Get(id int) gorms.ArbitrosGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	arbitro := gorms.ArbitrosGorm{}
	row := db.QueryRow("select * from arbitros where id_arbitro = ?", id)
	error := row.Scan(&arbitro.IDArbitro, &arbitro.IDPersona, &arbitro.IDCampeonato)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return arbitro
}

// Save arbritros
func (ed *ArbitrosDaoImpl) Save(e *gorms.ArbitrosGorm) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	isDelete := true
	if e.IDArbitro > 0 {
		isDelete, _ = ed.Delete(e.IDArbitro, e.IDPersona, e.IDCampeonato)
	}

	if isDelete == true {
		_, error := db.Exec("insert into arbitros (id_arbitro, id_persona, id_campeonato) values(?,?,?)",
			e.IDArbitro,
			e.IDPersona,
			e.IDCampeonato,
		)
		if error != nil {
			log.Println(error)
			panic(error)
		}
	}
	return e.IDArbitro
}

// Delete arbitro
func (ed *ArbitrosDaoImpl) Delete(IDArbitro int64, IDPersona int64, IDCampeonato int64) (bool, error) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from arbitros where id_arbitro = ? and id_persona = ? and id_campeonato", IDArbitro, IDPersona, IDCampeonato)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			return false, error
		}
	}
	return true, nil
}
