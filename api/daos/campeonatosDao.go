package daos

import (
	"database/sql"
	"deprimera/api/application"
	"deprimera/api/models"
	"fmt"
	"log"
)

type CampeonatosDaoImpl struct{}

func (ed *CampeonatosDaoImpl) GetAll() []models.Campeonatos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select id_campeonato, id_liga, id_modelo, descripcion, fecha_inicio, fecha_fin from campeonatos")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	var campeonatos []models.Campeonatos
	for rows.Next() {
		campeonato := models.Campeonatos{}
		error := rows.Scan(&campeonato.IDCampeonato, &campeonato.IDLiga, &campeonato.IDModelo, &campeonato.Descripcion, &campeonato.FechaInicio, &campeonato.FechaFin)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		campeonatos = append(campeonatos, campeonato)
	}
	return campeonatos
}

func (ed *CampeonatosDaoImpl) Get(id int) models.Campeonatos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select id_campeonato, id_liga, id_modelo, descripcion, fecha_inicio, fecha_fin from campeonatos where id_campeonato = ?", id)
	campeonato := models.Campeonatos{}
	error := row.Scan(&campeonato.IDCampeonato, &campeonato.IDLiga, &campeonato.IDModelo, &campeonato.Descripcion, &campeonato.FechaInicio, &campeonato.FechaFin)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return campeonato
}

func (ed *CampeonatosDaoImpl) Save(e *models.Campeonatos) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	fechaInicio := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		e.FechaInicio.Year(), e.FechaInicio.Month(), e.FechaInicio.Day(),
		e.FechaInicio.Hour(), e.FechaInicio.Minute(), e.FechaInicio.Second())
	fechaFin := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		e.FechaFin.Year(), e.FechaFin.Month(), e.FechaFin.Day(),
		e.FechaFin.Hour(), e.FechaFin.Minute(), e.FechaFin.Second())
	log.Println(fechaInicio)
	log.Println(fechaFin)
	if e.IDCampeonato > 0 {

		_, error := db.Exec("update campeonatos set descripcion=?, fecha_fin=?, fecha_inicio=?, id_liga=?, id_modelo=? where id_campeonato = ?",
			e.Descripcion, fechaFin, fechaInicio, e.IDLiga, e.IDModelo, e.IDCampeonato)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into campeonatos"+
			" (descripcion, fecha_fin, fecha_inicio, id_campeonato, id_liga, id_modelo) "+
			" values(?,?,?,?,?,?)", e.Descripcion, fechaFin, fechaInicio, e.IDCampeonato, e.IDLiga, e.IDModelo)
		IDCampeonato, error := res.LastInsertId()

		if error != nil {
			log.Println(error)
			panic(error)
		}
		e.IDCampeonato = IDCampeonato
	}
	return e.IDCampeonato
}

func (ed *CampeonatosDaoImpl) Delete(id int64) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from campeonatos where id_campeonato = ?", id)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return true
}
