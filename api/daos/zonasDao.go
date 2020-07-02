package daos

import (
	"database/sql"
	"deprimera/api/application"
	"deprimera/api/models"
	"log"
)

type ZonasDaoImpl struct{}

func (ed *ZonasDaoImpl) GetAll() []models.Zonas {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from zonas")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	var zonas []models.Zonas
	for rows.Next() {
		zona := models.Zonas{}
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

func (ed *ZonasDaoImpl) Get(id int) models.Zonas {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select * from zonas where id_zona = ?", id)
	zona := models.Zonas{}
	error := row.Scan(&zona.IDZona, &zona.IDCampeonato, &zona.Nombre)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return zona
}

func (ed *ZonasDaoImpl) Save(e *models.Zonas) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDZona > 0 {
		_, error := db.Exec("update zonas"+
			" set id_campeonato=?, nombre=?"+
			" where id_zona = ?", e.IDCampeonato, e.Nombre, e.IDZona)

		if error != nil {
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into zonas"+
			" (id_zona, id_campeonato, nombre) "+
			" values(?,?,?)", e.IDZona, e.IDCampeonato, e.Nombre)
		if error != nil {
			panic(error)
		}
		e.IDZona, _ = res.LastInsertId()
	}
	return e.IDZona
}

func (ed *ZonasDaoImpl) Delete(id int) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from zonas where id_zona = ?", id)
	if error != nil {
		panic(error)
	}
	return true
}
