package models

import (
	"time"
)

type Partidos struct {
	FechaEncuentro     time.Time `json:"fecha_encuentro"`
	IDArbitro          int64     `json:"id_arbitro"`
	IDAsistente        int64     `json:"id_asistente"`
	IDCampeonato       int       `json:"id_campeonato"`
	IDEquipoLocal      int       `json:"id_equipo_local"`
	IDEquipoVisitante  int       `json:"id_equipo_visitante"`
	IDLiga             int       `json:"id_liga"`
	IDPartidos         int64     `json:"id_partidos"`
	MotivoSuspencion   string    `json:"motivo_suspencion"`
	Observacion        string    `json:"observacion"`
	ResultadoLocal     int64     `json:"resultado_local"`
	ResultadoVisitante int64     `json:"resultado_visitante"`
	Suspendido         bool      `json:"suspendido"`
}


type PartidosFromDate struct {	
	IDPartidos         int64          `json:"id_partidos"`
	FechaEncuentro     time.Time      `json:"fecha_encuentro"`	
	LigaName           string         `json:"liga_name"`
	CampeonatoName     string 		  `json:"campeonato_name"`
	ELocalName         string         `json:"e_local_name"`
	EVisitName         string         `json:"e_visit_name"`
	ResultadoLocal     int64 		  `json:"resultado_local"`
	ResultadoVisitante int64  		  `json:"resultado_visitante"`
	Suspendido         string 		  `json:"suspendido"`
}

// TableName sets the insert table name for this struct type
func (p *Partidos) TableName() string {
	return "partidos"
}
