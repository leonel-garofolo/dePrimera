package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"log"
)

type ArbitrosDaoImpl struct{}

func (ed *ArbitrosDaoImpl) GetAll() []models.Arbitros {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from arbitros")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	var arbitros []models.Arbitros
	for rows.Next() {
		arbitro := models.Arbitros{}
		rows.Scan(&arbitro.IDArbitro)
		rows.Scan(&arbitro.IDPersona)
		arbitros = append(arbitros, arbitro)
	}
	return arbitros
}

func (ed *ArbitrosDaoImpl) Save(e *models.Arbitros) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	isDelete := ed.Delete(e.IDArbitro, e.IDPersona)
	if isDelete == true {
		_, error := db.Exec("insert into arbitros (id_arbitros, id_personas) values(?,?)", e.IDArbitro, e.IDPersona)

		if error != nil {
			panic(error)
		}
	}
	return e.IDArbitro
}

func (ed *ArbitrosDaoImpl) Delete(IDArbitro int64, IDPersona int64) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from arbitros where id_arbitro = ? and id_equipo = ?", IDArbitro, IDPersona)
	if error != nil {
		panic(error)
	}
	return true
}
