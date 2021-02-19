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
	Iniciado           bool      `json:"iniciado"`
	Finalizado         bool      `json:"finalizado"`
}

type PartidosFromDate struct {
	IDPartidos         int64     `json:"id_partidos"`
	FechaEncuentro     time.Time `json:"fecha_encuentro"`
	LigaName           string    `json:"liga_name"`
	CampeonatoName     string    `json:"campeonato_name"`
	ELocalName         string    `json:"e_local_name"`
	EVisitName         string    `json:"e_visit_name"`
	ResultadoLocal     int64     `json:"resultado_local"`
	ResultadoVisitante int64     `json:"resultado_visitante"`
	Suspendido         bool      `json:"suspendido"`
	Iniciado           bool      `json:"iniciado"`
	Finalizado         bool      `json:"finalizado"`
	Motivo             string    `json:"motivo"`
	GoleadoresLocal    string    `json:"goleadores_local"`
	GoleadoresVisit    string    `json:"goleadores_visit"`
	SancLocalAmar      string    `json:"sanc_local_amar"`
	SancLocalRojas     string    `json:"sanc_local_rojas"`
	SancVisitAmar      string    `json:"sanc_visit_amar"`
	SancVisitRojas     string    `json:"sanc_visit_rojas"`
}

type PartidoResult struct {
	IDPartidos                int64  `json:"id_partidos"`
	ResultadoLocal            int64  `json:"resultado_local"`
	GoleadoresLocal           string `json:"goleadores_local"`
	SancionAmarillasLocal     string `json:"sancion_amarillos_local"`
	SancionRojasLocal         string `json:"sancion_rojos_local"`
	ResultadoVisitante        int64  `json:"resultado_visitante"`
	GoleadoresVisitante       string `json:"goleadores_visitante"`
	SancionAmarillasVisitante string `json:"sancion_amarillos_visitantes"`
	SancionRojasVisitante     string `json:"sancion_rojos_visitantes"`
	Iniciado                  bool   `json:"iniciado"`
	Finalizado                bool   `json:"finalizado"`
	Suspendido                bool   `json:"suspendido"`
	Motivo                    string `json:"motivo"`
}

// TableName sets the insert table name for this struct type
func (p *Partidos) TableName() string {
	return "partidos"
}
