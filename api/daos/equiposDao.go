package daos

import (
	"database/sql"
	"deprimera/api/application"
	"deprimera/api/models"
	"log"
)

type EquiposDaoImpl struct{}

func (ed *EquiposDaoImpl) GetAll() []models.Equipos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from equipos")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	var equipos []models.Equipos
	for rows.Next() {
		equipo := models.Equipos{}
		error := rows.Scan(&equipo.IDEquipo, &equipo.IDLiga, &equipo.Nombre, &equipo.Habilitado, &equipo.Foto)
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

func (ed *EquiposDaoImpl) Get(id int) models.Equipos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select * from equipos where id_equipo = ?", id)
	equipo := models.Equipos{}
	error := row.Scan(&equipo.IDEquipo, &equipo.IDLiga, &equipo.Nombre, &equipo.Habilitado, &equipo.Foto)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return equipo
}

func (ed *EquiposDaoImpl) Save(e *models.Equipos) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDEquipo > 0 {
		_, error := db.Exec("update equipos"+
			" set id_liga=?, nombre=?, habilitado=?, foto=? "+
			" where id_equipo = ?", e.IDLiga, e.Nombre, e.Habilitado, e.Foto, e.IDEquipo)

		if error != nil {
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into equipos"+
			" (id_equipo, id_liga, nombre, habilitado, foto) "+
			" values(?,?,?,?,?)", e.IDEquipo, e.IDLiga, e.Nombre, e.Habilitado, e.Foto)
		if error != nil {
			panic(error)
		}
		e.IDEquipo, _ = res.LastInsertId()
	}
	return e.IDEquipo
}

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
