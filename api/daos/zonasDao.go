package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
)

type ZonasDaoImpl struct{}

func (ed *ZonasDaoImpl) GetAll() []gorms.ZonasGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from zonas")
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}

	var zonas []gorms.ZonasGorm
	for rows.Next() {
		zona := gorms.ZonasGorm{}
		error := rows.Scan(&zona.IDZona, &zona.IDCampeonato, &zona.Nombre)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		zonas = append(zonas, zona)
	}
	return zonas
}

func (ed *ZonasDaoImpl) Get(id int) gorms.ZonasGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select * from zonas where id_zona = $1", id)
	zona := gorms.ZonasGorm{}
	error := row.Scan(&zona.IDZona, &zona.IDCampeonato, &zona.Nombre)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return zona
}

func (ed *ZonasDaoImpl) Save(e *gorms.ZonasGorm) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDZona > 0 {
		_, error := db.Exec("update zonas"+
			" set id_campeonato=$1, nombre=$2"+
			" where id_zona = $3", e.IDCampeonato, e.Nombre, e.IDZona)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into zonas"+
			" (id_zona, id_campeonato, nombre) "+
			" values($1,$2,$3)", e.IDZona, e.IDCampeonato, e.Nombre)
		if error != nil {
			log.Println(error)
			panic(error)
		}
		e.IDZona, _ = res.LastInsertId()
	}
	return e.IDZona
}

func (ed *ZonasDaoImpl) Delete(id int) (bool, error) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from zonas where id_zona = $1", id)
	if error != nil {
		log.Println(error)
		return false, error
	}
	return true, nil
}
