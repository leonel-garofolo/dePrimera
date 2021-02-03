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

// CampeonatosDaoImpl struct
type CampeonatosDaoImpl struct{}

// GetAll campeonatos
func (ed *CampeonatosDaoImpl) GetAll() []gorms.CampeonatosGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select id_campeonato, id_liga, id_modelo, descripcion, fecha_inicio, fecha_fin, gen_fixture, gen_fixture_finish from campeonatos")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	var campeonatos []gorms.CampeonatosGorm
	for rows.Next() {
		campeonato := gorms.CampeonatosGorm{}
		error := rows.Scan(&campeonato.IDCampeonato, &campeonato.IDLiga, &campeonato.IDModelo, &campeonato.Descripcion, &campeonato.FechaInicio, &campeonato.FechaFin, &campeonato.GenFixture, &campeonato.GenFixtureFinish)
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

func (ed *CampeonatosDaoImpl) GetCampeonatoForUser(idUser string, idGrupo int) []models.Campeonatos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	//IdGrupo is Admin
	query := ""
	switch idGrupo {
	case 1: //ADMINISTRADOR
		query = " select c.id_campeonato, c.id_liga, c.id_modelo, c.descripcion, c.fecha_inicio, c.fecha_fin, c.gen_fixture ,c.gen_fixture_finish " +
			" from campeonatos c "
		break
	case 2: //DELEGADOS
		query = "select c.id_campeonato, c.id_liga, c.id_modelo, c.descripcion, c.fecha_inicio, c.fecha_fin, c.gen_fixture ,c.gen_fixture_finish " +
			"from campeonatos c " +
			"inner join campeonatos_equipos ce on ce.id_campeonato = c.id_campeonato " +
			"inner join asistentes a on a.id_campeonato = ce.id_campeonato " +
			"inner join personas p on p.id_persona = a.id_persona " +
			"where p.id_user= '" + idUser + "' and p.idgrupo = " + strconv.Itoa(idGrupo)
		break
	case 3: //JUGADORES
		query = "select c.id_campeonato, c.id_liga, c.id_modelo, c.descripcion, c.fecha_inicio, c.fecha_fin, c.gen_fixture ,c.gen_fixture_finish " +
			"from campeonatos c " +
			"inner join campeonatos_equipos ce on ce.id_campeonato = c.id_campeonato " +
			"inner join jugadores j on j.id_equipo = ce.id_equipo " +
			"inner join personas p on p.id_persona = j.id_persona " +
			"where p.id_user= '" + idUser + "' and p.idgrupo = " + strconv.Itoa(idGrupo)
		break
	case 4: //ARBITROS
		query = "select c.id_campeonato, c.id_liga, c.id_modelo, c.descripcion, c.fecha_inicio, c.fecha_fin, c.gen_fixture ,c.gen_fixture_finish " +
			"from campeonatos c " +
			"inner join campeonatos_equipos ce on ce.id_campeonato = c.id_campeonato " +
			"inner join arbitros a on a.id_campeonato = ce.id_campeonato " +
			"inner join personas p on p.id_persona = a.id_persona " +
			"where p.id_user= '" + idUser + "' and p.idgrupo = " + strconv.Itoa(idGrupo)
		break
	}

	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		log.Println("query: " + query)
		log.Println(err)
		panic(err)
	}

	var campeonatos []models.Campeonatos
	for rows.Next() {
		campeonato := models.Campeonatos{}
		error := rows.Scan(&campeonato.IDCampeonato, &campeonato.IDLiga, &campeonato.IDModelo, &campeonato.Descripcion, &campeonato.FechaInicio, &campeonato.FechaFin, &campeonato.GenFixture, &campeonato.GenFixtureFinish)
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

// Get campeonatos
func (ed *CampeonatosDaoImpl) Get(id int) gorms.CampeonatosGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select id_campeonato, id_liga, id_modelo, descripcion, fecha_inicio, fecha_fin, gen_fixture, get_fixture_finish from campeonatos where id_campeonato = $1", id)
	campeonato := gorms.CampeonatosGorm{}
	error := row.Scan(&campeonato.IDCampeonato, &campeonato.IDLiga, &campeonato.IDModelo, &campeonato.Descripcion, &campeonato.FechaInicio, &campeonato.FechaFin, &campeonato.GenFixture, &campeonato.GenFixtureFinish)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return campeonato
}

// Save campeonatos
func (ed *CampeonatosDaoImpl) Save(e *gorms.CampeonatosGorm) int64 {
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

		_, error := db.Exec("update campeonatos set descripcion=$1, fecha_fin=$2, fecha_inicio=$3, id_liga=$4, id_modelo=$5, gen_fixture =$6 where id_campeonato = $7",
			e.Descripcion, fechaFin, fechaInicio, e.IDLiga, e.IDModelo, e.GenFixture, e.IDCampeonato)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec(
			"insert into campeonatos (descripcion, fecha_fin, fecha_inicio, id_campeonato, id_liga, id_modelo) "+
				" values($1,$2,$3,$4,$5,$6)", e.Descripcion, fechaFin, fechaInicio, e.IDCampeonato, e.IDLiga, e.IDModelo)
		IDCampeonato, error := res.LastInsertId()

		if error != nil {
			log.Println(error)
			panic(error)
		}
		e.IDCampeonato = IDCampeonato
	}
	return e.IDCampeonato
}

// Delete campeonatos
func (ed *CampeonatosDaoImpl) Delete(id int64) (bool, error) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from campeonatos where id_campeonato = $1", id)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			return false, error
		}
	}
	return true, nil
}

// Get campeonatos
func (ed *CampeonatosDaoImpl) SaveCampeonatosGoleadores(idPartido int64, GoleadoresLocal string, GoleadoresVisitante string) (bool, error) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
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
		//Goles local
		saveGoles(db, idPartido, idJugLocal, nroCamLocal, GoleadoresLocal)

		//Goles visitante
		saveGoles(db, idPartido, idJugVisit, nroCamVisit, GoleadoresVisitante)
	}

	return true, nil
}

func saveGoles(db *sql.DB, idPartido int64, idJugador sql.NullInt64, nroCamDb sql.NullInt64, Goleadores string) {
	if Goleadores == "" {
		return
	}

	list := strings.Split(Goleadores, " ")
	for i, s := range list {
		fmt.Println(i, s)
		nroCam, _ := strconv.Atoi(s)
		if nroCam == int(nroCamDb.Int64) {
			_, error := db.Exec(
				"insert into campeonatos_goleadores(id_partido, id_jugadores, goles) values($1,$2,1) on DUPLICATE KEY UPDATE goles = goles + 1",
				idPartido,
				idJugador)
			if error != nil {
				log.Println(error)
				continue
			}
		}
	}
}

func (ed *CampeonatosDaoImpl) GetGoleadores(idCampeonato int) []models.CampeonatosGoleadores {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select cg.id_jugadores, e.nombre, per.nombre, per.apellido, sum(goles) as goles "+
		"from campeonatos_goleadores cg "+
		"inner join partidos p on p.id_partidos = cg.id_partido "+
		"inner join jugadores j on j.id_jugadores = cg.id_jugadores "+
		"inner join personas per on per.id_persona = j.id_persona "+
		"inner join equipos e on e.id_equipo = j.id_equipo "+
		"where p.id_campeonato = $1 "+
		"group by cg.id_jugadores, e.nombre, per.nombre, per.apellido", idCampeonato)
	if err != nil {
		log.Fatalln("Failed to query")
	}

	var goleadores []models.CampeonatosGoleadores
	for rows.Next() {
		goleador := models.CampeonatosGoleadores{}
		error := rows.Scan(&goleador.IDJugador, &goleador.Equipo, &goleador.Nombre, &goleador.Apellido, &goleador.Goles)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		goleadores = append(goleadores, goleador)
	}

	return goleadores
}
