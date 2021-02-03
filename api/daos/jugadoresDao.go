package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
)

type JugadoresDaoImpl struct{}

func (ed *JugadoresDaoImpl) GetAll() []models.Jugadores {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from jugadores")
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}

	jugadores := []models.Jugadores{}
	for rows.Next() {
		jugador := models.Jugadores{}
		error := rows.Scan(&jugador.IDJugador, &jugador.IDPersona, &jugador.IDEquipo, &jugador.NroCamiseta)
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

func (ed *JugadoresDaoImpl) Save(e *models.Jugadores) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	isDelete, _ := ed.Delete(e.IDPersona, e.IDEquipo)
	if isDelete == true {
		_, error := db.Exec("insert into jugadores (id_persona, id_equipo, nro_camiseta) values($1,$2, $3)", e.IDPersona, e.IDEquipo, e.NroCamiseta)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	}
	return e.IDJugador
}

func (ed *JugadoresDaoImpl) Delete(IDPersona int64, IDEquipo int64) (bool, error) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from jugadores where id_persona = $1 and id_equipo = $2", IDPersona, IDEquipo)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			return false, error
		}
	}
	return true, nil
}
