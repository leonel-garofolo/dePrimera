package daos

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/help"
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
		error := rows.Scan(&partido.IDPartidos, &partido.IDLiga, &partido.IDCampeonato, &partido.IDEquipoLocal, &partido.IDEquipoVisitante, &partido.IDArbitro, &partido.IDAsistente, &partido.FechaEncuentro, &partido.ResultadoLocal, &partido.ResultadoVisitante, &partido.Suspendido, &partido.MotivoSuspencion, &partido.Observacion)
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

func (ed *PartidosDaoImpl) GetAllFromDate(datePartidos string) []gorms.PartidosFromDateGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select p.id_partidos, p.fecha_encuentro, "+
		" l.nombre as liga_name, c.descripcion as campeonato_name, "+
		" e_local.nombre as e_local_name, e_visit.nombre as e_visit_name, "+
		" p.resultado_local, p.resultado_visitante, "+
		" p.suspendido "+
		" from partidos p "+
		" inner join ligas l on l.id_liga = p.id_liga "+
		" inner join campeonatos c on c.id_campeonato = p.id_campeonato "+
		" inner join equipos e_local on e_local.id_equipo = p.id_equipo_local "+
		" inner join equipos e_visit on e_visit.id_equipo = p.id_equipo_visitante "+
		" inner join arbitros a on a.id_arbitro = p.id_arbitro "+
		" inner join asistentes asis on asis.id_asistente = p.id_asistente "+
		" where fecha_encuentro like ?", datePartidos+"%")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	var partidos []gorms.PartidosFromDateGorm
	for rows.Next() {
		partido := gorms.PartidosFromDateGorm{}
		error := rows.Scan(
			&partido.IDPartidos,
			&partido.FechaEncuentro,
			&partido.LigaName,
			&partido.CampeonatoName,
			&partido.ELocalName,
			&partido.EVisitName,
			&partido.ResultadoLocal,
			&partido.ResultadoVisitante,
			&partido.Suspendido,
		)
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

// Get equipo
func (ed *PartidosDaoImpl) HistoryPlays(id int) []gorms.PartidosFromDateGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select p.id_partidos, p.fecha_encuentro, "+
		" l.nombre as liga_name, c.descripcion as campeonato_name, "+
		" e_local.nombre as e_local_name, e_visit.nombre as e_visit_name, "+
		" p.resultado_local, p.resultado_visitante, "+
		" p.suspendido "+
		" from partidos p "+
		" inner join ligas l on l.id_liga = p.id_liga "+
		" inner join campeonatos c on c.id_campeonato = p.id_campeonato "+
		" inner join equipos e_local on e_local.id_equipo = p.id_equipo_local "+
		" inner join equipos e_visit on e_visit.id_equipo = p.id_equipo_visitante "+
		" inner join arbitros a on a.id_arbitro = p.id_arbitro "+
		" inner join asistentes asis on asis.id_asistente = p.id_asistente "+
		" where id_equipo_local = ? or id_equipo_visitante = ?", id, id)
	if err != nil {
		log.Fatalln("Failed to query")
	}

	var partidos []gorms.PartidosFromDateGorm
	for rows.Next() {
		partido := gorms.PartidosFromDateGorm{}
		error := rows.Scan(
			&partido.IDPartidos,
			&partido.FechaEncuentro,
			&partido.LigaName,
			&partido.CampeonatoName,
			&partido.ELocalName,
			&partido.EVisitName,
			&partido.ResultadoLocal,
			&partido.ResultadoVisitante,
			&partido.Suspendido,
		)
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

var dateFormat = "2006-01-02"

func (ed *PartidosDaoImpl) SaveFixture(idLiga int, idCampeonato int, dateFrom time.Time, rondas [][]help.Partido) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	datesAvailability := getWeekendFromDate(dateFrom)
	dateCount := 0
	for i := 0; i < len(rondas); i++ {
		fmt.Print("Ronda " + strconv.Itoa((i + 1)) + ": ")
		for j := 0; j < len(rondas[i]); j++ {
			fmt.Print("   " + strconv.Itoa((1 + rondas[i][j].Local)) + "-" + strconv.Itoa((1 + rondas[i][j].Visitante)))

			_, error := db.Exec("insert into partidos(id_liga, id_campeonato, id_equipo_local, id_equipo_visitante, fecha_encuentro ) "+
				"values ( "+
				"?, "+
				"?, "+
				"(select id_equipo from equipos where id_campeonato = ? and nro_equipo = ?), "+
				"(select id_equipo from equipos where id_campeonato = ? and nro_equipo = ?), "+
				"? "+
				")",
				idLiga,
				idCampeonato,
				idCampeonato, 1+rondas[i][j].Local,
				idCampeonato, 1+rondas[i][j].Visitante,
				datesAvailability[dateCount])
			if error != nil {
				panic(error)
			}
			dateCount++
		}

		fmt.Println()
	}

	fmt.Println("VUELTA")

	for i := 0; i < len(rondas); i++ {
		fmt.Print("Ronda " + strconv.Itoa((i + 1)) + ": ")

		for j := 0; j < len(rondas[i]); j++ {
			fmt.Print("   " + strconv.Itoa((1 + rondas[i][j].Visitante)) + "-" + strconv.Itoa((1 + rondas[i][j].Local)))

			_, error := db.Exec("insert into partidos(id_liga, id_campeonato, id_equipo_local, id_equipo_visitante, fecha_encuentro ) "+
				"values ( "+
				"?, "+
				"?, "+
				"(select id_equipo from equipos where id_campeonato = ? and nro_equipo = ?), "+
				"(select id_equipo from equipos where id_campeonato = ? and nro_equipo = ?), "+
				"? "+
				")",
				idLiga,
				idCampeonato,
				idCampeonato, 1+rondas[i][j].Visitante,
				idCampeonato, 1+rondas[i][j].Local,
				datesAvailability[dateCount])
			if error != nil {
				panic(error)
			}
			dateCount++
		}

		fmt.Println()
	}
}

func getWeekendFromDate(start time.Time) []time.Time {
	var datesAvailability []time.Time
	start.Year()
	end, _ := time.Parse(dateFormat, "2021-12-01")
	end = end.Add(time.Hour * 24)

	for t := start; t.Before(end); t = t.Add(time.Hour * 24) {
		if t.Weekday() == time.Saturday {
			datesAvailability = append(datesAvailability, t)
		}
	}
	return datesAvailability
}
