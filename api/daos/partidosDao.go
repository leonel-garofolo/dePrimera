package daos

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
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
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}

	partidos := []gorms.PartidosGorm{}
	for rows.Next() {
		partido := gorms.PartidosGorm{}
		error := rows.Scan(&partido.IDPartidos, &partido.IDLiga, &partido.IDCampeonato, &partido.IDEquipoLocal, &partido.IDEquipoVisitante, &partido.IDArbitro, &partido.IDAsistente, &partido.FechaEncuentro, &partido.ResultadoLocal, &partido.ResultadoVisitante, &partido.Suspendido, &partido.MotivoSuspencion, &partido.Observacion, &partido.Iniciado, &partido.Finalizado)
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

func (ed *PartidosDaoImpl) GetAllFromEquipo(idEquipo int) []models.PartidosFromDate {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select p.id_partidos, p.fecha_encuentro, "+
		" l.nombre as liga_name, c.descripcion as campeonato_name, "+
		" e_local.nombre as e_local_name, e_visit.nombre as e_visit_name, "+
		" p.resultado_local, p.resultado_visitante, "+
		" p.suspendido,  p.iniciado, p.finalizado, p.motivo_suspencion, "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from campeonatos_goleadores aux_cg "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_cg.id_jugadores and aux_jug.id_equipo = e_local.id_equipo "+
		" 	where aux_cg.id_partido = p.id_partidos) as goleadores_local, "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from campeonatos_goleadores aux_cg  "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_cg.id_jugadores and aux_jug.id_equipo = e_visit.id_equipo "+
		" 	where aux_cg.id_partido = p.id_partidos) as goleadores_visit , "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from sanciones_jugadores aux_sj  "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_local.id_equipo "+
		" 	where aux_sj.id_sancion=2 and aux_sj.id_partidos = p.id_partidos) as sanciones_local_amarillas, "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from sanciones_jugadores aux_sj  "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_local.id_equipo "+
		" 	where aux_sj.id_sancion=1 and aux_sj.id_partidos = p.id_partidos) as sanciones_local_rojas, "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from sanciones_jugadores aux_sj  "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_visit.id_equipo "+
		" 	where aux_sj.id_sancion=2 and aux_sj.id_partidos = p.id_partidos) as sanciones_visit_amarillas, "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from sanciones_jugadores aux_sj  "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_visit.id_equipo "+
		" 	where aux_sj.id_sancion=1 and aux_sj.id_partidos = p.id_partidos) as sanciones_visit_rojas  "+
		" from partidos p "+
		" inner join ligas l on l.id_liga = p.id_liga "+
		" inner join campeonatos c on c.id_campeonato = p.id_campeonato "+
		" inner join equipos e_local on e_local.id_equipo = p.id_equipo_local "+
		" inner join equipos e_visit on e_visit.id_equipo = p.id_equipo_visitante "+
		" left join arbitros a on a.id_arbitro = p.id_arbitro "+
		" left join asistentes asis on asis.id_asistente = p.id_asistente "+
		" where e_local.id_equipo = $1 or e_visit.id_equipo = $2", idEquipo, idEquipo)
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}

	goleadoresLocal := sql.NullString{}
	goleadoresVisit := sql.NullString{}
	sancLocalAmar := sql.NullString{}
	sancLocalRojas := sql.NullString{}
	sancVisitAmar := sql.NullString{}
	sancVisitRojas := sql.NullString{}
	motivo := sql.NullString{}
	partidos := []models.PartidosFromDate{}
	for rows.Next() {
		partido := models.PartidosFromDate{}
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
			&partido.Iniciado,
			&partido.Finalizado,
			&motivo,
			&goleadoresLocal,
			&goleadoresVisit,
			&sancLocalAmar,
			&sancLocalRojas,
			&sancVisitAmar,
			&sancVisitRojas,
		)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		partido.GoleadoresLocal = goleadoresLocal.String
		partido.GoleadoresVisit = goleadoresVisit.String
		partido.SancLocalAmar = sancLocalAmar.String
		partido.SancLocalRojas = sancLocalRojas.String
		partido.SancVisitAmar = sancLocalAmar.String
		partido.SancVisitRojas = sancLocalRojas.String
		partido.Motivo = motivo.String
		partidos = append(partidos, partido)
	}
	return partidos
}

func (ed *PartidosDaoImpl) GetAllFromDate(datePartidos string) []models.PartidosFromDate {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select p.id_partidos, p.fecha_encuentro, "+
		" l.nombre as liga_name, c.descripcion as campeonato_name, "+
		" e_local.nombre as e_local_name, e_visit.nombre as e_visit_name, "+
		" p.resultado_local, p.resultado_visitante, "+
		" p.suspendido,  p.iniciado, p.finalizado, p.motivo_suspencion, "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from campeonatos_goleadores aux_cg "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_cg.id_jugadores and aux_jug.id_equipo = e_local.id_equipo "+
		" 	where aux_cg.id_partido = p.id_partidos) as goleadores_local, "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from campeonatos_goleadores aux_cg  "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_cg.id_jugadores and aux_jug.id_equipo = e_visit.id_equipo "+
		" 	where aux_cg.id_partido = p.id_partidos) as goleadores_visit , "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from sanciones_jugadores aux_sj  "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_local.id_equipo "+
		" 	where aux_sj.id_sancion=2 and aux_sj.id_partidos = p.id_partidos) as sanciones_local_amarillas, "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from sanciones_jugadores aux_sj  "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_local.id_equipo "+
		" 	where aux_sj.id_sancion=1 and aux_sj.id_partidos = p.id_partidos) as sanciones_local_rojas, "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from sanciones_jugadores aux_sj  "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_visit.id_equipo "+
		" 	where aux_sj.id_sancion=2 and aux_sj.id_partidos = p.id_partidos) as sanciones_visit_amarillas, "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from sanciones_jugadores aux_sj  "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_visit.id_equipo "+
		" 	where aux_sj.id_sancion=1 and aux_sj.id_partidos = p.id_partidos) as sanciones_visit_rojas  "+
		" from partidos p "+
		" inner join ligas l on l.id_liga = p.id_liga "+
		" inner join campeonatos c on c.id_campeonato = p.id_campeonato "+
		" inner join equipos e_local on e_local.id_equipo = p.id_equipo_local "+
		" inner join equipos e_visit on e_visit.id_equipo = p.id_equipo_visitante "+
		" left join arbitros a on a.id_arbitro = p.id_arbitro "+
		" left join asistentes asis on asis.id_asistente = p.id_asistente "+
		" where fecha_encuentro between $1 and $2"+
		" order by fecha_encuentro asc", datePartidos+" 00:00:01", datePartidos+" 23:59:59")
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}

	goleadoresLocal := sql.NullString{}
	goleadoresVisit := sql.NullString{}
	sancLocalAmar := sql.NullString{}
	sancLocalRojas := sql.NullString{}
	sancVisitAmar := sql.NullString{}
	sancVisitRojas := sql.NullString{}
	motivo := sql.NullString{}
	partidos := []models.PartidosFromDate{}
	for rows.Next() {
		partido := models.PartidosFromDate{}
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
			&partido.Iniciado,
			&partido.Finalizado,
			&motivo,
			&goleadoresLocal,
			&goleadoresVisit,
			&sancLocalAmar,
			&sancLocalRojas,
			&sancVisitAmar,
			&sancVisitRojas,
		)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		partido.GoleadoresLocal = goleadoresLocal.String
		partido.GoleadoresVisit = goleadoresVisit.String
		partido.SancLocalAmar = sancLocalAmar.String
		partido.SancLocalRojas = sancLocalRojas.String
		partido.SancVisitAmar = sancLocalAmar.String
		partido.SancVisitRojas = sancLocalRojas.String
		partido.Motivo = motivo.String
		partidos = append(partidos, partido)
	}
	return partidos
}

func (ed *PartidosDaoImpl) GetAllFromCampeonato(idTorneo int) []models.PartidosFromDate {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select p.id_partidos, p.fecha_encuentro, "+
		" l.nombre as liga_name, c.descripcion as campeonato_name, "+
		" e_local.nombre as e_local_name, e_visit.nombre as e_visit_name, "+
		" p.resultado_local, p.resultado_visitante, "+
		" p.suspendido,  p.iniciado, p.finalizado, p.motivo_suspencion, "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from campeonatos_goleadores aux_cg "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_cg.id_jugadores and aux_jug.id_equipo = e_local.id_equipo "+
		" 	where aux_cg.id_partido = p.id_partidos) as goleadores_local, "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from campeonatos_goleadores aux_cg  "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_cg.id_jugadores and aux_jug.id_equipo = e_visit.id_equipo "+
		" 	where aux_cg.id_partido = p.id_partidos) as goleadores_visit , "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from sanciones_jugadores aux_sj  "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_local.id_equipo "+
		" 	where aux_sj.id_sancion = 2 and aux_sj.id_partidos = p.id_partidos) as sanciones_local_amarillas, "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from sanciones_jugadores aux_sj  "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_local.id_equipo "+
		" 	where aux_sj.id_sancion = 1 and aux_sj.id_partidos = p.id_partidos) as sanciones_local_rojas, "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from sanciones_jugadores aux_sj  "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_visit.id_equipo "+
		" 	where aux_sj.id_sancion = 2 and aux_sj.id_partidos = p.id_partidos) as sanciones_visit_amarillas, "+
		" (select string_agg(aux_jug.nro_camiseta::text, ' ')  "+
		" 	from sanciones_jugadores aux_sj  "+
		" 	inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_visit.id_equipo "+
		" 	where aux_sj.id_sancion = 1 and aux_sj.id_partidos = p.id_partidos) as sanciones_visit_rojas  "+
		" from partidos p "+
		" inner join ligas l on l.id_liga = p.id_liga "+
		" inner join campeonatos c on c.id_campeonato = p.id_campeonato "+
		" inner join equipos e_local on e_local.id_equipo = p.id_equipo_local "+
		" inner join equipos e_visit on e_visit.id_equipo = p.id_equipo_visitante "+
		" left join arbitros a on a.id_arbitro = p.id_arbitro "+
		" left join asistentes asis on asis.id_asistente = p.id_asistente "+
		" where c.id_campeonato = $1"+
		" order by fecha_encuentro asc", idTorneo)
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}

	goleadoresLocal := sql.NullString{}
	goleadoresVisit := sql.NullString{}
	sancLocalAmar := sql.NullString{}
	sancLocalRojas := sql.NullString{}
	sancVisitAmar := sql.NullString{}
	sancVisitRojas := sql.NullString{}
	motivo := sql.NullString{}
	partidos := []models.PartidosFromDate{}
	for rows.Next() {
		partido := models.PartidosFromDate{}
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
			&partido.Iniciado,
			&partido.Finalizado,
			&motivo,
			&goleadoresLocal,
			&goleadoresVisit,
			&sancLocalAmar,
			&sancLocalRojas,
			&sancVisitAmar,
			&sancVisitRojas,
		)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		partido.GoleadoresLocal = goleadoresLocal.String
		partido.GoleadoresVisit = goleadoresVisit.String
		partido.SancLocalAmar = sancLocalAmar.String
		partido.SancLocalRojas = sancLocalRojas.String
		partido.SancVisitAmar = sancVisitAmar.String
		partido.SancVisitRojas = sancVisitRojas.String
		partido.Motivo = motivo.String
		partidos = append(partidos, partido)
	}
	return partidos
}

// Nombre Equipo, Pts, PG, PE, PP
func (ed *PartidosDaoImpl) GetTablePosition(idTorneo int) []models.EquiposTablePos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select ce.id_campeonato, ce.id_equipo, e.nombre, "+
		"	ce.nro_equipo, "+
		"	ce.puntos, ce.p_gan, ce.p_emp, ce.p_per "+
		" from campeonatos_equipos ce "+
		" inner join campeonatos c on c.id_campeonato = ce.id_campeonato "+
		" inner join equipos e on e.id_equipo = ce.id_equipo "+
		" where c.id_campeonato = $1 "+
		" order by ce.puntos desc", idTorneo)
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}

	equiposPos := []models.EquiposTablePos{}
	for rows.Next() {
		partido := models.EquiposTablePos{}
		error := rows.Scan(
			&partido.IDCampeonato,
			&partido.IDEquipo,
			&partido.Nombre,
			&partido.NroEquipo,
			&partido.Puntos,
			&partido.PartidoGanado,
			&partido.PartidoEmpatado,
			&partido.PartidoPerdido,
		)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		equiposPos = append(equiposPos, partido)
	}
	return equiposPos
}

func (ed *PartidosDaoImpl) Get(id int) gorms.PartidosGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select * from partidos where id_partidos = $1", id)
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
			" set id_arbitro=$1, id_asistente=$2, id_campeonato=$3, id_equipo_local=$4, id_equipo_visitante=$5, id_liga=$6, motivo_suspencion=$7, observacion=$8, resultado_local=$9, resultado_visitante=$10, suspendido=$11, fecha_encuentro= $12 "+
			" where id_partidos = $13", e.IDArbitro, e.IDAsistente, e.IDCampeonato, e.IDEquipoLocal, e.IDEquipoVisitante, e.IDLiga, e.MotivoSuspencion, e.Observacion, e.ResultadoLocal, e.ResultadoVisitante, e.FechaEncuentro, e.IDPartidos)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into partidos"+
			" (id_partidos, id_arbitro, id_asistente, id_campeonato, id_equipo_local, id_equipo_visitante, id_liga, motivo_suspencion, observacion, resultado_local, resultado_visitante, suspendido, fecha_encuentro) "+
			" values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)", e.IDPartidos, e.IDArbitro, e.IDAsistente, e.IDCampeonato, e.IDEquipoLocal, e.IDEquipoVisitante, e.IDLiga, e.MotivoSuspencion, e.Observacion, e.ResultadoLocal, e.ResultadoVisitante, e.Suspendido, e.FechaEncuentro)
		if error != nil {
			log.Println(error)
			panic(error)
		}
		e.IDPartidos, _ = res.LastInsertId()
	}
	return e.IDPartidos
}

func (ed *PartidosDaoImpl) SaveResult(e *gorms.PartidoResultGorm) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDPartidos > 0 {
		_, error := db.Exec("update partidos"+
			" set resultado_local=$1, resultado_visitante=$2, iniciado =$3, finalizado =$4, suspendido =$5, motivo_suspencion =$6 "+
			" where id_partidos = $7", e.ResultadoLocal, e.ResultadoVisitante, e.Iniciado, e.Finalizado, e.Suspendido, e.Motivo, e.IDPartidos)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	}
	return e.IDPartidos
}

func (ed *PartidosDaoImpl) Delete(id int) (bool, error) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from partidos where id_partidos = $1", id)
	if error != nil {
		return false, error
	}
	return true, nil
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
		" left join arbitros a on a.id_arbitro = p.id_arbitro "+
		" left join asistentes asis on asis.id_asistente = p.id_asistente "+
		" where id_equipo_local = $1 or id_equipo_visitante = $2", id, id)
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}

	partidos := []gorms.PartidosFromDateGorm{}
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

func (ed *PartidosDaoImpl) GetFuturePartidos() []string {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select fecha_encuentro from partidos p where p.fecha_encuentro > current_date group by fecha_encuentro")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	datesPartidos := []string{}
	for rows.Next() {
		datePartido := ""
		error := rows.Scan(
			&datePartido,
		)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		datesPartidos = append(datesPartidos, datePartido)
	}
	return datesPartidos
}

var dateFormat = "2006-01-02"

func (ed *PartidosDaoImpl) SaveFixture(idLiga int, idCampeonato int, dateFrom time.Time, rondas [][]help.Partido) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, errorDelete := db.Exec("delete from partidos where id_liga = $1 and id_campeonato = $2", idLiga, idCampeonato)
	if errorDelete != nil {
		fmt.Println(errorDelete)
		panic(errorDelete)
	}

	datesAvailability := getWeekendFromDate(dateFrom)
	dateCount := 0
	for i := 0; i < len(rondas); i++ {
		fmt.Print("Ronda " + strconv.Itoa((i + 1)) + ": ")
		for j := 0; j < len(rondas[i]); j++ {
			fmt.Print("   " + strconv.Itoa((1 + rondas[i][j].Local)) + "-" + strconv.Itoa((1 + rondas[i][j].Visitante)))

			_, error := db.Exec("insert into partidos(id_liga, id_campeonato, id_equipo_local, id_equipo_visitante, fecha_encuentro ) "+
				"values ( "+
				"$1, "+
				"$2, "+
				"(select id_equipo from campeonatos_equipos where id_liga =$3 and id_campeonato = $4 and nro_equipo = $5), "+
				"(select id_equipo from campeonatos_equipos where id_liga =$6 and id_campeonato = $7 and nro_equipo = $8), "+
				"$9 "+
				")",
				idLiga,
				idCampeonato,
				idLiga, idCampeonato, 1+rondas[i][j].Local,
				idLiga, idCampeonato, 1+rondas[i][j].Visitante,
				datesAvailability[dateCount])
			if error != nil {
				fmt.Println(error)
				panic(error)
			}
		}
		dateCount++
		fmt.Println()
	}

	fmt.Println("VUELTA")

	for i := 0; i < len(rondas); i++ {
		fmt.Print("Ronda " + strconv.Itoa((i + 1)) + ": ")

		for j := 0; j < len(rondas[i]); j++ {
			fmt.Print("   " + strconv.Itoa((1 + rondas[i][j].Visitante)) + "-" + strconv.Itoa((1 + rondas[i][j].Local)))

			_, error := db.Exec("insert into partidos(id_liga, id_campeonato, id_equipo_local, id_equipo_visitante, fecha_encuentro ) "+
				"values ( "+
				"$1, "+
				"$2, "+
				"(select id_equipo from campeonatos_equipos where id_liga =$3 and id_campeonato = $4 and nro_equipo = $5), "+
				"(select id_equipo from campeonatos_equipos where id_liga =$6 and id_campeonato = $7 and nro_equipo = $8), "+
				"$9 "+
				")",
				idLiga,
				idCampeonato,
				idLiga, idCampeonato, 1+rondas[i][j].Visitante,
				idLiga, idCampeonato, 1+rondas[i][j].Local,
				datesAvailability[dateCount])
			if error != nil {
				fmt.Println(error)
				panic(error)
			}
		}
		dateCount++
		fmt.Println()
	}
}

func (ed *PartidosDaoImpl) FinishFixtureGen(idLiga int, idCampeonato int) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, errorDelete := db.Exec("update campeonato set gen_fixture =1 where id_liga = $1 and id_campeonato = $2", idLiga, idCampeonato)
	if errorDelete != nil {
		fmt.Println(errorDelete)
		return false
	}
	return true
}

func (ed *PartidosDaoImpl) FinalizarPartido(idLiga int64, idCampeonato int64, idEquipoLocal int64, idEquipoVisit int64, statusResult string) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	switch statusResult {
	case "GL":
		_, errorUpdateGan := db.Exec("update campeonatos_equipos set puntos=(puntos + 3), p_gan =(p_gan + 1) where id_liga = $1 and id_campeonato = $2 and id_equipo =$3", idLiga, idCampeonato, idEquipoLocal)
		if errorUpdateGan != nil {
			fmt.Println(errorUpdateGan)
			panic(errorUpdateGan)
		}

		_, errorUpdatePer := db.Exec("update campeonatos_equipos set p_per =(p_per + 1) where id_liga = $1 and id_campeonato = $2 and id_equipo =$3", idLiga, idCampeonato, idEquipoVisit)
		if errorUpdatePer != nil {
			fmt.Println(errorUpdatePer)
			panic(errorUpdatePer)
		}

	case "GV":
		_, errorUpdateGan := db.Exec("update campeonatos_equipos set puntos=(puntos + 3), p_gan =(p_gan + 1) where id_liga = $1 and id_campeonato = $2 and id_equipo =$3", idLiga, idCampeonato, idEquipoVisit)
		if errorUpdateGan != nil {
			fmt.Println(errorUpdateGan)
			panic(errorUpdateGan)
		}

		_, errorUpdatePer := db.Exec("update campeonatos_equipos set p_per =(p_per + 1) where id_liga = $1 and id_campeonato = $2 and id_equipo =$3", idLiga, idCampeonato, idEquipoLocal)
		if errorUpdatePer != nil {
			fmt.Println(errorUpdatePer)
			panic(errorUpdatePer)
		}

	case "E":
		_, errorUpdateGan := db.Exec("update campeonatos_equipos set puntos=(puntos + 1), p_emp =(p_emp + 1) where id_liga = $1 and id_campeonato = $2 and id_equipo in ($3, $4)", idLiga, idCampeonato, idEquipoLocal, idEquipoVisit)
		if errorUpdateGan != nil {
			fmt.Println(errorUpdateGan)
			panic(errorUpdateGan)
		}
	}

}

func getWeekendFromDate(start time.Time) []time.Time {
	var datesAvailability []time.Time
	start.Year()
	end, _ := time.Parse(dateFormat, "2022-12-01")
	end = end.Add(time.Hour * 24)

	for t := start; t.Before(end); t = t.Add(time.Hour * 24) {
		if t.Weekday() == time.Saturday {
			datesAvailability = append(datesAvailability, t)
		}
	}
	return datesAvailability
}
