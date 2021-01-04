package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
)

// ArbitrosDaoImpl struct
type QueryDaoImpl struct{}

type QueryConfiguracion struct {
	Ligas       int64 `json:"ligas"`
	Campeonatos int64 `json:"campeonatos"`
	Equipos     int64 `json:"equipos"`
	Arbitros    int64 `json:"arbitros"`
	Asistentes  int64 `json:"asistentes"`
	Jugadores   int64 `json:"jugadores"`
}

// Login user
func (ed *QueryDaoImpl) GetConfiguracionesSize() QueryConfiguracion {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	queryConfiguracion := QueryConfiguracion{}
	row := db.QueryRow("select " +
		"(select count(*) from ligas) as ligas, " +
		"(select count(*) from campeonatos) as campeonatos, " +
		"(select count(*) from equipos) as equipos, " +
		"(select count(*) from arbitros) as arbitros, " +
		"(select count(*) from asistentes) as asistentes, " +
		"(select count(*) from jugadores) as jugadores " +
		";")
	error := row.Scan(&queryConfiguracion.Ligas, &queryConfiguracion.Campeonatos, &queryConfiguracion.Equipos, &queryConfiguracion.Arbitros, &queryConfiguracion.Asistentes, &queryConfiguracion.Jugadores)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return queryConfiguracion
}
