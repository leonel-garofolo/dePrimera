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
		error := rows.Scan(&equipo.IDEquipo, &equipo.Nombre, &equipo.Habilitado, &equipo.Foto)
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

func (ed *EquiposDaoImpl) GetAllFromCampeonato(IDLiga int, IDCampeonato int) []gorms.EquiposGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select id_equipo from campeonatos_equipos where id_liga=? and id_campeonato = ?", IDLiga, IDCampeonato)
	if err != nil {
		log.Fatalln("Failed to query")
	}

	var equipos []gorms.EquiposGorm
	for rows.Next() {
		equipo := gorms.EquiposGorm{}
		error := rows.Scan(&equipo.IDEquipo)
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

func (ed *EquiposDaoImpl) UpdateNro(IDLiga int, IDCampeonato int) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select id_equipo from campeonatos_equipos where id_liga=? and id_campeonato = ?", IDLiga, IDCampeonato)
	if err != nil {
		log.Fatalln("Failed to query")
	}

	nroEquipo := 1
	for rows.Next() {
		idEquipo := int64(0)
		error := rows.Scan(&idEquipo)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}

		_, errorInsert := db.Exec("update campeonatos_equipos set nro_equipo = ? where id_liga =? and id_campeonato=? and id_equipo=?",
			nroEquipo, IDLiga, IDCampeonato, idEquipo)
		if errorInsert != nil {
			log.Println(errorInsert)
			panic(errorInsert)
		}
		nroEquipo++
	}
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
	error := row.Scan(&equipo.IDEquipo, &equipo.Nombre, &equipo.Habilitado, &equipo.Foto)
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
			" set nombre=?, habilitado=?, foto=? "+
			" where id_equipo = ?", e.Nombre, e.Habilitado, e.Foto, e.IDEquipo)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into equipos"+
			" (id_equipo, nombre, foto) "+
			" values(?,?,?,?)", e.IDEquipo, e.Nombre, e.Foto)
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
