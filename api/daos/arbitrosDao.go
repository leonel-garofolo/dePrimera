package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
)

// ArbitrosDaoImpl struct
type ArbitrosDaoImpl struct{}

// GetAll arbritros
func (ed *ArbitrosDaoImpl) GetAll() []models.Arbitros {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from arbitros")
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}

	arbitros := []models.Arbitros{}
	for rows.Next() {
		arbitro := models.Arbitros{}
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

func (ed *ArbitrosDaoImpl) Get(id int) models.Arbitros {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	arbitro := models.Arbitros{}
	row := db.QueryRow("select * from arbitros where id_arbitro = $1", id)
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
func (ed *ArbitrosDaoImpl) Save(e *models.Arbitros) int64 {
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
		_, error := db.Exec("insert into arbitros (id_arbitro, id_persona, id_campeonato) values($1,$2,$3)",
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

	_, error := db.Exec("delete from arbitros where id_arbitro = $1 and id_persona = $2 and id_campeonato = $3", IDArbitro, IDPersona, IDCampeonato)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			return false, error
		}
	}
	return true, nil
}
