package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"log"
)

type PartidosDaoImpl struct{}

func (ed *PartidosDaoImpl) GetAll() []models.Partidos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from partidos")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	var partidos []models.Partidos
	for rows.Next() {
		partido := models.Partidos{}
		rows.Scan(&partido.IDPartidos)
		rows.Scan(&partido.IDArbitro)
		rows.Scan(&partido.IDAsistente)
		rows.Scan(&partido.IDCampeonato)
		rows.Scan(&partido.IDEquipoLocal)
		rows.Scan(&partido.IDEquipoVisitante)
		rows.Scan(&partido.IDLiga)
		rows.Scan(&partido.MotivoSuspencion)
		rows.Scan(&partido.Observacion)
		rows.Scan(&partido.ResultadoLocal)
		rows.Scan(&partido.ResultadoVisitante)
		rows.Scan(&partido.Suspendido)
		partidos = append(partidos, partido)
	}
	return partidos
}

func (ed *PartidosDaoImpl) Get(id int) models.Partidos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select * from partidos where id_partido = ?", id)
	partido := models.Partidos{}
	row.Scan(&partido.IDPartidos)
	row.Scan(&partido.IDArbitro)
	row.Scan(&partido.IDAsistente)
	row.Scan(&partido.IDCampeonato)
	row.Scan(&partido.IDEquipoLocal)
	row.Scan(&partido.IDEquipoVisitante)
	row.Scan(&partido.IDLiga)
	row.Scan(&partido.MotivoSuspencion)
	row.Scan(&partido.Observacion)
	row.Scan(&partido.ResultadoLocal)
	row.Scan(&partido.ResultadoVisitante)
	row.Scan(&partido.Suspendido)
	return partido
}

func (ed *PartidosDaoImpl) Save(e *models.Partidos) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDPartidos > 0 {
		_, error := db.Exec("update partidos"+
			" set id_arbitro=?, id_asistente=?, id_campeonato=?, id_equipo_local=?, id_equipo_visitante=?, id_liga=?, motivo_suspencion=?, observacion=?, resultado_local=?, resultado_visitante=?, suspendido, fecha?encuentro= ? "+
			" where id_partido = ?", e.IDArbitro, e.IDAsistente, e.IDCampeonato, e.IDEquipoLocal, e.IDEquipoVisitante, e.IDLiga, e.MotivoSuspencion, e.Observacion, e.ResultadoLocal, e.ResultadoVisitante, e.FechaEncuentro, e.IDPartidos)

		if error != nil {
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into partidos"+
			" (id_partido, id_arbitro, id_asistente, id_campeonato, id_equipo_local, id_equipo_visitante, id_liga, motivo_suspencion, observacion, resultado_local, resultado_visitante, suspendido, encuentro) "+
			" values(?,?,?,?,?,?,?,?,?,?,?,?,?)", e.IDPartidos, e.IDArbitro, e.IDAsistente, e.IDCampeonato, e.IDEquipoLocal, e.IDEquipoVisitante, e.IDLiga, e.MotivoSuspencion, e.Observacion, e.ResultadoLocal, e.ResultadoVisitante, e.Suspendido, e.FechaEncuentro)

		IDPartidos, error := res.LastInsertId()

		if error != nil {
			panic(error)
		}
		e.IDPartidos = IDPartidos
	}
	return e.IDPartidos
}

func (ed *PartidosDaoImpl) Delete(id int) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from partidos where id_partido = ?", id)
	if error != nil {
		panic(error)
	}
	return true
}
