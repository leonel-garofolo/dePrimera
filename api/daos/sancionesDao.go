package daos

import (
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
		rows.Scan(&sancion.IDSanciones)
		rows.Scan(&sancion.IDLigas)
		rows.Scan(&sancion.Descripcion)
		rows.Scan(&sancion.Observaciones)
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
	row.Scan(&sancion.IDSanciones)
	row.Scan(&sancion.IDLigas)
	row.Scan(&sancion.Descripcion)
	row.Scan(&sancion.Observaciones)
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
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into sanciones"+
			" (id_sanciones, id_ligas, descripcion, observaciones) "+
			" values(?,?,?,?)", e.IDSanciones, e.IDLigas, e.Descripcion, e.Observaciones)

		IDSanciones, error := res.LastInsertId()

		if error != nil {
			panic(error)
		}
		e.IDSanciones = IDSanciones
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
		panic(error)
	}
	return true
}
