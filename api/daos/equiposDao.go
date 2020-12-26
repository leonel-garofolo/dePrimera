package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
)

// EquiposDaoImpl sarasa
type EquiposDaoImpl struct{}

// GetAll object
func (ed *EquiposDaoImpl) GetAll() []gorms.EquiposGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from equipos")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	var equipos []gorms.EquiposGorm
	for rows.Next() {
		equipo := gorms.EquiposGorm{}
		error := rows.Scan(&equipo.IDEquipo, &equipo.IDCampeonato, &equipo.Nombre, &equipo.Habilitado, &equipo.Foto)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}

		equipos = append(equipos, equipo)
	}
	return equipos
}

func (ed *EquiposDaoImpl) GetAllFromLiga(idLiga int64) []gorms.EquiposGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from equipos where id_campeonato = ?", idLiga)
	if err != nil {
		log.Fatalln("Failed to query")
	}

	var equipos []gorms.EquiposGorm
	for rows.Next() {
		equipo := gorms.EquiposGorm{}
		error := rows.Scan(&equipo.IDEquipo, &equipo.IDCampeonato, &equipo.Nombre, &equipo.Habilitado, &equipo.Foto)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}

		equipos = append(equipos, equipo)
	}
	return equipos
}

// Get equipo
func (ed *EquiposDaoImpl) Get(id int) gorms.EquiposGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select * from equipos where id_equipo = ?", id)
	equipo := gorms.EquiposGorm{}
	error := row.Scan(&equipo.IDEquipo, &equipo.IDCampeonato, &equipo.Nombre, &equipo.Habilitado, &equipo.Foto)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return equipo
}

// Save equipos
func (ed *EquiposDaoImpl) Save(e *gorms.EquiposGorm) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDEquipo > 0 {
		_, error := db.Exec("update equipos"+
			" set id_campeonato=?, nombre=?, habilitado=?, foto=? "+
			" where id_equipo = ?", e.IDCampeonato, e.Nombre, e.Habilitado, e.Foto, e.IDEquipo)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into equipos"+
			" (id_equipo, id_campeonato, nombre, habilitado, foto) "+
			" values(?,?,?,?,?)", e.IDEquipo, e.IDCampeonato, e.Nombre, e.Habilitado, e.Foto)
		if error != nil {
			log.Println(error)
			panic(error)
		}
		e.IDEquipo, _ = res.LastInsertId()
	}
	return e.IDEquipo
}

// Delete equipos
func (ed *EquiposDaoImpl) Delete(id int) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from equipos where id_equipo = ?", id)
	if error != nil {
		panic(error)
	}
	return true
}
