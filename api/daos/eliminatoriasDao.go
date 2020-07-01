package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"log"
)

type EliminatoriasDaoImpl struct{}

func (ed *EliminatoriasDaoImpl) GetAll() []models.Eliminatorias {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from eliminatorias")
	if err != nil {
		log.Fatalln("Failed to query")
	}
	var eliminatorias []models.Eliminatorias
	for rows.Next() {
		eliminatoria := models.Eliminatorias{}
		error := rows.Scan(&eliminatoria.IDEliminatoria, &eliminatoria.IDCampeonato, &eliminatoria.IDPartido, &eliminatoria.NroLlave)
		if error != nil {
			log.Println(error)
			panic(error)
		}
		eliminatorias = append(eliminatorias, eliminatoria)
	}
	return eliminatorias
}

func (ed *EliminatoriasDaoImpl) Get(id int) models.Eliminatorias {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select * from eliminatorias where id_eliminatoria = ?", id)
	eliminatoria := models.Eliminatorias{}
	error := row.Scan(&eliminatoria.IDEliminatoria, &eliminatoria.IDCampeonato, &eliminatoria.IDPartido, &eliminatoria.NroLlave)
	if error != nil {
		log.Println(error)
		panic(error)
	}
	return eliminatoria
}

func (ed *EliminatoriasDaoImpl) Save(e *models.Eliminatorias) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDEliminatoria > 0 {
		_, error := db.Exec("update eliminatorias"+
			" set id_campeonato=?, id_partido=?, nro_llave=? "+
			" where id_eliminatoria=?", e.IDCampeonato, e.IDPartido, e.NroLlave, e.IDEliminatoria)

		if error != nil {
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into eliminatorias"+
			" (id_eliminatoria, id_campeonato, id_partido, nro_llave) "+
			" values(?,?,?,?)", e.IDEliminatoria, e.IDCampeonato, e.IDPartido, e.NroLlave)
		if error != nil {
			panic(error)
		}
		e.IDEliminatoria, _ = res.LastInsertId()
	}
	return e.IDEliminatoria
}

func (ed *EliminatoriasDaoImpl) Delete(id int) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}
	_, error := db.Exec("delete from eliminatorias where id_eliminatoria = ?", id)
	if error != nil {
		panic(error)
	}
	return true
}
