package daos

import (
	"database/sql"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	"log"
)

type JugadoresDaoImpl struct{}

func (ed *JugadoresDaoImpl) GetAll() []gorms.JugadoresGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from jugadores")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	jugadores := []gorms.JugadoresGorm{}
	for rows.Next() {
		asistente := gorms.JugadoresGorm{}
		error := rows.Scan(&asistente.IDJugador, &asistente.IDPersona)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		jugadores = append(jugadores, asistente)
	}
	return jugadores
}

func (ed *JugadoresDaoImpl) Save(e *gorms.JugadoresGorm) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	isDelete := ed.Delete(e.IDJugador, e.IDPersona)
	if isDelete == true {
		_, error := db.Exec("insert into jugadores (id_jugador, id_persona) values(?,?)", e.IDJugador, e.IDPersona)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	}
	return e.IDJugador
}

func (ed *JugadoresDaoImpl) Delete(IDJugador int64, IDPersona int64) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from jugadores where id_jugador = ? and id_persona = ?", IDJugador, IDPersona)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return true
}
