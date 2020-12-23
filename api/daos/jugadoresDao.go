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
		jugador := gorms.JugadoresGorm{}
		error := rows.Scan(&jugador.IDJugador, &jugador.IDPersona, &jugador.IDEquipo)
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

func (ed *JugadoresDaoImpl) Save(e *gorms.JugadoresGorm) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	isDelete := ed.Delete(e.IDPersona, e.IDEquipo)
	if isDelete == true {
		_, error := db.Exec("insert into jugadores (id_persona, id_equipo) values(?,?)", e.IDPersona, e.IDEquipo)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	}
	return e.IDJugador
}

func (ed *JugadoresDaoImpl) Delete(IDPersona int64, IDEquipo int64) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from jugadores where id_persona = ? and id_equipo = ?", IDPersona, IDEquipo)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}	
	}
	return true
}
