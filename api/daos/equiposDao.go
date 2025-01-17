package daos

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
)

// EquiposDaoImpl sarasa
type EquiposDaoImpl struct{}

// GetAll object
func (ed *EquiposDaoImpl) GetAll() []models.Equipos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query(" select e.id_equipo, e.nombre, e.habilitado, e.foto, ce.id_campeonato, ce.nro_equipo " +
		"from equipos e " +
		"inner join campeonatos_equipos ce on ce.id_equipo = e.id_equipo ")
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}

	equipos := []models.Equipos{}
	for rows.Next() {
		equipo := models.Equipos{}
		error := rows.Scan(&equipo.IDEquipo, &equipo.Nombre, &equipo.Habilitado, &equipo.Foto, &equipo.IDCampeonato, &equipo.NroEquipo)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		row := db.QueryRow("select id_campeonato from campeonatos_equipos where id_equipo=$1 limit 1", &equipo.IDEquipo)
		errorCamp := row.Scan(&equipo.IDCampeonato)
		if errorCamp != nil {
			if errorCamp != sql.ErrNoRows {
				log.Println(errorCamp)
			}
		}
		equipos = append(equipos, equipo)
	}

	return equipos
}

func (ed *EquiposDaoImpl) GetEquiposFromUser(idUser string, idGrupo int) []models.Equipos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	//IdGrupo is Admin
	query := ""
	switch idGrupo {
	case 1: //ADMINISTRADOR
		query = " select e.id_equipo, e.nombre, e.habilitado, e.foto, ce.id_campeonato, ce.nro_equipo " +
			"from equipos e " +
			"inner join campeonatos_equipos ce on ce.id_equipo = e.id_equipo "

		break
	case 2: //DELEGADOS
		query = "select e.id_equipo, e.nombre, e.habilitado, e.foto, ce.id_campeonato, ce.nro_equipo  " +
			"from equipos e " +
			"inner join campeonatos_equipos ce on ce.id_equipo = e.id_equipo " +
			"inner join asistentes a on a.id_campeonato = ce.id_campeonato " +
			"inner join personas p on p.id_persona = a.id_persona " +
			"where p.id_user= '" + idUser + "' and p.idgrupo = " + strconv.Itoa(idGrupo)
		break
	case 3: //JUGADORES
		query = "select e.id_equipo, e.nombre, e.habilitado, e.foto, ce.id_campeonato, ce.nro_equipo  " +
			"from equipos e " +
			"inner join campeonatos_equipos ce on ce.id_equipo = e.id_equipo " +
			"inner join jugadores j on j.id_equipo = ce.id_equipo " +
			"inner join personas p on p.id_persona = j.id_persona " +
			"where p.id_user= '" + idUser + "' and p.idgrupo = " + strconv.Itoa(idGrupo)
		break
	case 4: //ARBITROS
		query = "select e.id_equipo, e.nombre, e.habilitado, e.foto, ce.id_campeonato, ce.nro_equipo  " +
			"from equipos e " +
			"inner join campeonatos_equipos ce on ce.id_equipo = e.id_equipo " +
			"inner join arbitros a on a.id_campeonato = ce.id_campeonato " +
			"inner join personas p on p.id_persona = a.id_persona " +
			"where p.id_user= '" + idUser + "' and p.idgrupo = " + strconv.Itoa(idGrupo)
		break
	}

	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		log.Panic(err)
		panic(err)
	}

	equipos := []models.Equipos{}
	for rows.Next() {
		equipo := models.Equipos{}
		error := rows.Scan(&equipo.IDEquipo, &equipo.Nombre, &equipo.Habilitado, &equipo.Foto, &equipo.IDCampeonato, &equipo.NroEquipo)
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

func (ed *EquiposDaoImpl) GetPlantel(idEquipo int) []models.JugadoresPlantel {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query(
		"select j.id_jugadores, p.apellido, p.nombre, j.nro_camiseta  "+
			"from jugadores j "+
			"inner join personas p on j.id_persona = p.id_persona "+
			"where p.apellido is not null and p.nombre is not null and j.id_equipo = $1 "+
			"order by p.apellido asc, p.nombre asc", idEquipo,
	)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	jugadores := []models.JugadoresPlantel{}
	for rows.Next() {
		jugador := models.JugadoresPlantel{}
		error := rows.Scan(&jugador.IDJugador, &jugador.Apellido, &jugador.Nombre, &jugador.NroCamiseta)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}

		jugadores = append(jugadores, jugador)
	}
	return jugadores
}

func (ed *EquiposDaoImpl) GetAllFromCampeonato(IDLiga int, IDCampeonato int) []gorms.EquiposGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select id_equipo from campeonatos_equipos where id_liga=$1 and id_campeonato = $2", IDLiga, IDCampeonato)
	if err != nil {
		log.Fatalln("Failed to query")
	}

	equipos := []gorms.EquiposGorm{}
	for rows.Next() {
		equipo := gorms.EquiposGorm{}
		error := rows.Scan(&equipo.IDEquipo)
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

func (ed *EquiposDaoImpl) UpdateNro(IDLiga int, IDCampeonato int) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select id_equipo from campeonatos_equipos where id_liga=$1 and id_campeonato = $2", IDLiga, IDCampeonato)
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}

	nroEquipo := 1
	for rows.Next() {
		idEquipo := int64(0)
		error := rows.Scan(&idEquipo)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}

		_, errorInsert := db.Exec("update campeonatos_equipos set nro_equipo = $1 where id_liga =$2 and id_campeonato=$3 and id_equipo=$4",
			nroEquipo, IDLiga, IDCampeonato, idEquipo)
		if errorInsert != nil {
			log.Println(errorInsert)
			panic(errorInsert)
		}
		nroEquipo++
	}
}

// Get equipo
func (ed *EquiposDaoImpl) Get(id int) gorms.EquiposGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select * from equipos where id_equipo = $1", id)
	equipo := gorms.EquiposGorm{}
	error := row.Scan(&equipo.IDEquipo, &equipo.Nombre, &equipo.Habilitado, &equipo.Foto)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return equipo
}

// Save equipos
func (ed *EquiposDaoImpl) Save(e *gorms.EquiposGorm) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDEquipo > 0 {
		_, error := db.Exec("update equipos"+
			" set nombre=$1, habilitado=$2, foto=$3 "+
			" where id_equipo = $4", e.Nombre, e.Habilitado, e.Foto, e.IDEquipo)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into equipos"+
			" (id_equipo, nombre, foto) "+
			" values($1,$2,$3)", e.IDEquipo, e.Nombre, e.Foto)
		if error != nil {
			log.Println(error)
			panic(error)
		}
		e.IDEquipo, _ = res.LastInsertId()
	}
	return e.IDEquipo
}

func (ed *EquiposDaoImpl) SaveEquiposCampeonatos(idCampeonato int64, idEquipo int64) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("insert IGNORE into campeonatos_equipos "+
		"select c.id_liga, c.id_campeonato, $1, (select count(*) + 1 from campeonatos_equipos aux where aux.id_campeonato = $2) as nro_equipo, "+
		"	0 as p_gan, "+
		"	0 as p_emp, "+
		"	0 as p_per, "+
		"	0 as puntos "+
		"from campeonatos c "+
		"where c.id_campeonato = $3", idEquipo, idCampeonato, idCampeonato)
	if error != nil {
		log.Println(error)
		return false
	}
	return true
}

// Delete equipos
func (ed *EquiposDaoImpl) Delete(id int) (bool, error) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from equipos where id_equipo = $1", id)
	if error != nil {
		return false, error
	}
	return true, nil
}
