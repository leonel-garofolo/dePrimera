package daos

import (
	"database/sql"
	"deprimera/api/application"
	"deprimera/api/models"
	"log"
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
		log.Fatalln("Failed to query")
	}

	jugadores := []models.Jugadores{}
	for rows.Next() {
		jugador := models.Jugadores{}
		error := rows.Scan(&jugador.IDJugador, &jugador.IDPersona)
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
