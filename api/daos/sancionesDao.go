package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
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
		log.Fatalln("Failed to query")
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

func (ed *SancionesDaoImpl) GetSancionesFromCampeonato(idCampeonato int) []gorms.SancionesJugadoresFromCampeonatoGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select p.apellido_nombre, e.nombre as e_nombre, "+
		" 	(case when sj.id_sancion = 1 then count(sj.id_sancion) else 0 end ) as c_rojas, "+
		" 	(case when sj.id_sancion = 2 then count(sj.id_sancion) else 0 end ) as c_amarillas, "+
		" 	(case when sj.id_sancion = 3 then count(sj.id_sancion) else 0 end ) as c_azules "+
		" from sanciones_jugadores sj "+
		" inner join jugadores j on j.id_jugadores = sj.id_jugador "+
		" inner join equipos e on e.id_equipo = j.id_equipo "+
		" inner join personas p on p.id_persona = j.id_persona "+
		" where sj.id_campeonato = ? "+
		" group by p.apellido_nombre, e.nombre, sj.id_sancion "+
		" order by p.apellido_nombre asc", idCampeonato)
	if err != nil {
		log.Fatalln("Failed to query")
	}
	var sancionesJugadores []gorms.SancionesJugadoresFromCampeonatoGorm
	for rows.Next() {
		sancion := gorms.SancionesJugadoresFromCampeonatoGorm{}
		error := rows.Scan(&sancion.ApellidoNombre, &sancion.ENombre, &sancion.CRojas, &sancion.CAmarillas, &sancion.CAzules)
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

	row := db.QueryRow("select * from sanciones where id_sancion = ?", id)
	if err != nil {
		log.Fatalln("Failed to query")
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
			" set  id_ligas=?, descripcion=?, observaciones=? "+
			" where id_sanciones=?", e.Descripcion, e.Observaciones, e.IDSanciones)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into sanciones"+
			" (id_sanciones, id_ligas, descripcion, observaciones) "+
			" values(?,?,?,?)", e.IDSanciones, e.Descripcion, e.Observaciones)
		if error != nil {
			log.Println(error)
			panic(error)
		}
		e.IDSanciones, _ = res.LastInsertId()
	}
	return e.IDSanciones
}

func (ed *SancionesDaoImpl) Delete(id int) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from sanciones where id_sancion = ?", id)
	if error != nil {
		log.Println(error)
		panic(error)
	}
	return true
}
