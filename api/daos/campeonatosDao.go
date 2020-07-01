package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"log"
)

type CampeonatosDaoImpl struct{}

func (ed *CampeonatosDaoImpl) GetAll() []models.Campeonatos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from campeonatos")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	var campeonatos []models.Campeonatos
	for rows.Next() {
		campeonato := models.Campeonatos{}
		rows.Scan(&campeonato.Descripcion)
		rows.Scan(&campeonato.FechaFin)
		rows.Scan(&campeonato.FechaInicio)
		rows.Scan(&campeonato.IDCampeonato)
		rows.Scan(&campeonato.IDLiga)
		rows.Scan(&campeonato.IDModelo)
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

	row := db.QueryRow("select * from campeonatos where id_campeonato = ?", id)
	campeonato := models.Campeonatos{}
	row.Scan(&campeonato.Descripcion)
	row.Scan(&campeonato.FechaFin)
	row.Scan(&campeonato.FechaInicio)
	row.Scan(&campeonato.IDCampeonato)
	row.Scan(&campeonato.IDLiga)
	row.Scan(&campeonato.IDModelo)
	return campeonato
}

func (ed *CampeonatosDaoImpl) Save(e *models.Campeonatos) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDLiga > 0 {
		_, error := db.Exec("update campeonatos set descripcion=?, fecha_fin=?, fecha_inicio=?, id_liga=?, id_modelo=? where id_campeonato = ?",
			e.Descripcion, e.FechaFin, e.FechaInicio, e.IDLiga, e.IDModelo, e.IDCampeonato)

		if error != nil {
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into campeonatos"+
			" (descripcion, fecha_fin, fecha_inicio, id_campeonato, id_liga, id_modelo) "+
			" values(?,?,?,?,?,?)", e.Descripcion, e.FechaFin, e.FechaInicio, e.IDCampeonato, e.IDLiga, e.IDModelo)
		IDCampeonato, error := res.LastInsertId()

		if error != nil {
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
		panic(error)
	}
	return true
}
