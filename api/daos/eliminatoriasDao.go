package daos

import (
	"database/sql"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	"log"
)

// EliminatoriasDaoImpl struct
type EliminatoriasDaoImpl struct{}

// GetAll eliminatorias
func (ed *EliminatoriasDaoImpl) GetAll() []gorms.EliminatoriasGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from eliminatorias")
	if err != nil {
		log.Fatalln("Failed to query")
	}
	var eliminatorias []gorms.EliminatoriasGorm
	for rows.Next() {
		eliminatoria := gorms.EliminatoriasGorm{}
		error := rows.Scan(&eliminatoria.IDEliminatoria, &eliminatoria.IDCampeonato, &eliminatoria.IDPartido, &eliminatoria.NroLlave)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		eliminatorias = append(eliminatorias, eliminatoria)
	}
	return eliminatorias
}

// Get eliminatoria
func (ed *EliminatoriasDaoImpl) Get(id int) gorms.EliminatoriasGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select * from eliminatorias where id_eliminatoria = ?", id)
	eliminatoria := gorms.EliminatoriasGorm{}
	error := row.Scan(&eliminatoria.IDEliminatoria, &eliminatoria.IDCampeonato, &eliminatoria.IDPartido, &eliminatoria.NroLlave)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return eliminatoria
}

// Save eliminatoria
func (ed *EliminatoriasDaoImpl) Save(e *gorms.EliminatoriasGorm) int64 {
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
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into eliminatorias"+
			" (id_eliminatoria, id_campeonato, id_partido, nro_llave) "+
			" values(?,?,?,?)", e.IDEliminatoria, e.IDCampeonato, e.IDPartido, e.NroLlave)
		if error != nil {
			log.Println(error)
			panic(error)
		}
		e.IDEliminatoria, _ = res.LastInsertId()
	}
	return e.IDEliminatoria
}

// Delete eliminatoria
func (ed *EliminatoriasDaoImpl) Delete(id int) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}
	_, error := db.Exec("delete from eliminatorias where id_eliminatoria = ?", id)
	if error != nil {
		log.Println(error)
		panic(error)
	}
	return true
}
