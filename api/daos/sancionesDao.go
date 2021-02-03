package daos

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
)

type SancionesDaoImpl struct{}

func (ed *SancionesDaoImpl) GetAll() []gorms.SancionesGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from sanciones")
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}
	var sanciones []gorms.SancionesGorm
	for rows.Next() {
		sancion := gorms.SancionesGorm{}
		error := rows.Scan(&sancion.IDSanciones, &sancion.Descripcion, &sancion.Observaciones)
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

func (ed *SancionesDaoImpl) GetSancionesFromCampeonato(idCampeonato int) []models.SancionesJugadoresFromCampeonato {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select p.nombre as p_nombre, p.apellido as p_apellido, e.nombre as e_nombre, "+
		" 	(case when sj.id_sancion = 1 then count(sj.id_sancion) else 0 end ) as c_rojas, "+
		" 	(case when sj.id_sancion = 2 then count(sj.id_sancion) else 0 end ) as c_amarillas, "+
		" 	(case when sj.id_sancion = 3 then count(sj.id_sancion) else 0 end ) as c_azules "+
		" from sanciones_jugadores sj "+
		" inner join partidos partido on partido.id_partidos = sj.id_partidos "+
		" inner join jugadores j on j.id_jugadores = sj.id_jugador "+
		" inner join equipos e on e.id_equipo = j.id_equipo "+
		" inner join personas p on p.id_persona = j.id_persona "+
		" where partido.id_campeonato = $1 "+
		" group by p.nombre, p.apellido, e.nombre, sj.id_sancion "+
		" order by p.apellido asc, p.nombre asc", idCampeonato)
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}
	var sancionesJugadores []models.SancionesJugadoresFromCampeonato
	for rows.Next() {
		sancion := models.SancionesJugadoresFromCampeonato{}
		error := rows.Scan(&sancion.Nombre, &sancion.Apellido, &sancion.ENombre, &sancion.CRojas, &sancion.CAmarillas, &sancion.CAzules)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		sancionesJugadores = append(sancionesJugadores, sancion)
	}
	return sancionesJugadores
}

func (ed *SancionesDaoImpl) Get(id int) gorms.SancionesGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select * from sanciones where id_sancion = $1", id)
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}
	sancion := gorms.SancionesGorm{}
	error := row.Scan(&sancion.IDSanciones, &sancion.Descripcion, &sancion.Observaciones)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return sancion
}

func (ed *SancionesDaoImpl) Save(e *gorms.SancionesGorm) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDSanciones > 0 {
		_, error := db.Exec("update sanciones"+
			" set  id_ligas=$1, descripcion=$2, observaciones=$3 "+
			" where id_sanciones=$4", e.Descripcion, e.Observaciones, e.IDSanciones)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into sanciones"+
			" (id_sanciones, id_ligas, descripcion, observaciones) "+
			" values($1,$2,$3,$4)", e.IDSanciones, e.Descripcion, e.Observaciones)
		if error != nil {
			log.Println(error)
			panic(error)
		}
		e.IDSanciones, _ = res.LastInsertId()
	}
	return e.IDSanciones
}

func (ed *SancionesDaoImpl) Delete(id int) (bool, error) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from sanciones where id_sancion = $1", id)
	if error != nil {
		log.Println(error)
		return false, error
	}
	return true, nil
}

func (ed *SancionesDaoImpl) SavePartido(idPartido int64, amarillasLocal string, rojasLocal string, amarillasVisitante string, rojasVisitante string) (bool, error) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from sanciones_jugadores where id_partidos = $1", idPartido)
	if error != nil {
		log.Println(error)
		return false, error
	}

	rows, err := db.Query("select  "+
		"	jlocal.id_jugadores as jug_local, jlocal.nro_camiseta as nro_camiseta_local, "+
		"	jvisit.id_jugadores as jug_visit, jvisit.nro_camiseta as nro_camiseta_visit  "+
		"from partidos p "+
		"left join jugadores jlocal on jlocal.id_equipo = p.id_equipo_local "+
		"left join jugadores jvisit on jvisit.id_equipo = p.id_equipo_visitante "+
		"where id_partidos = $1", idPartido)
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}

	var idJugLocal sql.NullInt64
	var nroCamLocal sql.NullInt64
	var idJugVisit sql.NullInt64
	var nroCamVisit sql.NullInt64
	for rows.Next() {
		error := rows.Scan(&idJugLocal, &nroCamLocal, &idJugVisit, &nroCamVisit)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		//Amarillas local
		saveSancionForTeam(db, idPartido, idJugLocal, nroCamLocal, amarillasLocal, 2)

		//Rojas Local
		saveSancionForTeam(db, idPartido, idJugLocal, nroCamLocal, rojasLocal, 1)

		//Amarillas Visitante
		saveSancionForTeam(db, idPartido, idJugVisit, nroCamVisit, amarillasVisitante, 2)

		//Rojas Visitante
		saveSancionForTeam(db, idPartido, idJugVisit, nroCamVisit, rojasVisitante, 1)
	}

	return true, nil
}

func (ed *SancionesDaoImpl) SavePartidoFinalizado(idPartido int64, finalizado bool) (bool, error) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if finalizado {
		_, error := db.Exec("update partidos set finalizado = 1 where id_partidos = $1", idPartido)
		if error != nil {
			log.Println(error)
			return false, error
		}
	}

	return true, nil
}

func saveSancionForTeam(db *sql.DB, idPartido int64, idJugador sql.NullInt64, nroCamDb sql.NullInt64, card string, idSancion int64) {
	if card == "" && idJugador.Valid && nroCamDb.Valid {
		return
	}

	list := strings.Split(card, " ")
	for i, s := range list {
		fmt.Println(i, s)
		nroCam, _ := strconv.Atoi(s)
		if nroCam == int(nroCamDb.Int64) {
			_, error := db.Exec("insert into sanciones_jugadores(id_sancion, id_jugador, id_partidos) values($1,$2,$3) ", idSancion, idJugador, idPartido)
			if error != nil {
				log.Println(error)
				continue
			}
		}
	}
}
