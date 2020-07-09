package daos

import (
	"database/sql"
	"deprimera/api/application"
	"deprimera/api/models"
	"log"
)

type SancionesDaoImpl struct{}

func (ed *SancionesDaoImpl) GetAll() []models.Sanciones {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from sanciones")
	if err != nil {
		log.Fatalln("Failed to query")
	}
	var sanciones []models.Sanciones
	for rows.Next() {
		sancion := models.Sanciones{}
		error := rows.Scan(&sancion.IDSanciones, &sancion.IDLigas, &sancion.Descripcion, &sancion.Observaciones)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		sanciones = append(sanciones, sancion)
	}
	return sanciones
}

func (ed *SancionesDaoImpl) Get(id int) models.Sanciones {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select * from sanciones where id_sancion = ?", id)
	if err != nil {
		log.Fatalln("Failed to query")
	}
	sancion := models.Sanciones{}
	error := row.Scan(&sancion.IDSanciones, &sancion.IDLigas, &sancion.Descripcion, &sancion.Observaciones)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return sancion
}

func (ed *SancionesDaoImpl) Save(e *models.Sanciones) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDSanciones > 0 {
		_, error := db.Exec("update sanciones"+
			" set  id_ligas=?, descripcion=?, observaciones=? "+
			" where id_sanciones=?", e.IDLigas, e.Descripcion, e.Observaciones, e.IDSanciones)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into sanciones"+
			" (id_sanciones, id_ligas, descripcion, observaciones) "+
			" values(?,?,?,?)", e.IDSanciones, e.IDLigas, e.Descripcion, e.Observaciones)
		if error != nil {
			log.Println(error)
			panic(error)
		}
		e.IDSanciones, _ = res.LastInsertId()
	}
	return e.IDSanciones
}

func (ed *SancionesDaoImpl) Delete(id int) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from sanciones where id_sancion = ?", id)
	if error != nil {
		log.Println(error)
		panic(error)
	}
	return true
}
