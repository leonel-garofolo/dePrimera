package daos

import (
	"database/sql"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	"log"
)

type PartidosDaoImpl struct{}

func (ed *PartidosDaoImpl) GetAll() []gorms.PartidosGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from partidos")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	var partidos []gorms.PartidosGorm
	for rows.Next() {
		partido := gorms.PartidosGorm{}
		error := rows.Scan(&partido.IDPartidos, &partido.IDArbitro, &partido.IDAsistente, &partido.IDCampeonato, &partido.IDEquipoLocal, &partido.IDEquipoVisitante, &partido.IDLiga, &partido.MotivoSuspencion, &partido.Observacion, &partido.ResultadoLocal, &partido.ResultadoVisitante, &partido.Suspendido)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		partidos = append(partidos, partido)
	}
	return partidos
}

func (ed *PartidosDaoImpl) Get(id int) gorms.PartidosGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select * from partidos where id_partidos = ?", id)
	partido := gorms.PartidosGorm{}
	error := row.Scan(&partido.IDPartidos, &partido.IDArbitro, &partido.IDAsistente, &partido.IDCampeonato, &partido.IDEquipoLocal, &partido.IDEquipoVisitante, &partido.IDLiga, &partido.MotivoSuspencion, &partido.Observacion, &partido.ResultadoLocal, &partido.ResultadoVisitante, &partido.Suspendido)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return partido
}

func (ed *PartidosDaoImpl) Save(e *gorms.PartidosGorm) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDPartidos > 0 {
		_, error := db.Exec("update partidos"+
			" set id_arbitro=?, id_asistente=?, id_campeonato=?, id_equipo_local=?, id_equipo_visitante=?, id_liga=?, motivo_suspencion=?, observacion=?, resultado_local=?, resultado_visitante=?, suspendido, fecha_encuentro= ? "+
			" where id_partidos = ?", e.IDArbitro, e.IDAsistente, e.IDCampeonato, e.IDEquipoLocal, e.IDEquipoVisitante, e.IDLiga, e.MotivoSuspencion, e.Observacion, e.ResultadoLocal, e.ResultadoVisitante, e.FechaEncuentro, e.IDPartidos)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into partidos"+
			" (id_partidos, id_arbitro, id_asistente, id_campeonato, id_equipo_local, id_equipo_visitante, id_liga, motivo_suspencion, observacion, resultado_local, resultado_visitante, suspendido, fecha_encuentro) "+
			" values(?,?,?,?,?,?,?,?,?,?,?,?,?)", e.IDPartidos, e.IDArbitro, e.IDAsistente, e.IDCampeonato, e.IDEquipoLocal, e.IDEquipoVisitante, e.IDLiga, e.MotivoSuspencion, e.Observacion, e.ResultadoLocal, e.ResultadoVisitante, e.Suspendido, e.FechaEncuentro)
		if error != nil {
			log.Println(error)
			panic(error)
		}
		e.IDPartidos, _ = res.LastInsertId()
	}
	return e.IDPartidos
}

func (ed *PartidosDaoImpl) Delete(id int) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from partidos where id_partidos = ?", id)
	if error != nil {
		panic(error)
	}
	return true
}
