package gorms

import (
	"database/sql"
	"time"
)

type PartidosGorm struct {	
	IDArbitro          sql.NullInt64  `gorm:"column:id_arbitro"`
	IDAsistente        sql.NullInt64  `gorm:"column:id_asistente"`
	IDCampeonato       int64            `gorm:"column:id_campeonato"`
	IDEquipoLocal      int64            `gorm:"column:id_equipo_local"`
	IDEquipoVisitante  int64            `gorm:"column:id_equipo_visitante"`
	IDLiga             int64            `gorm:"column:id_liga"`
	IDPartidos         int64            `gorm:"column:id_partidos;primary_key"`
	FechaEncuentro     time.Time      `gorm:"column:fecha_encuentro"`
	MotivoSuspencion   sql.NullString `gorm:"column:motivo_suspencion"`
	Observacion        sql.NullString `gorm:"column:observacion"`
	ResultadoLocal     sql.NullInt64  `gorm:"column:resultado_local"`
	ResultadoVisitante sql.NullInt64  `gorm:"column:resultado_visitante"`
	Suspendido         string `gorm:"column:suspendido"`
}

// TableName sets the insert table name for this struct type
func (p *PartidosGorm) TableName() string {
	return "partidos"
}

type PartidosFromDateGorm struct {	
	IDPartidos         int64          `gorm:"column:id_partidos"`
	FechaEncuentro     time.Time      `gorm:"column:fecha_encuentro"`	
	LigaName           string         `gorm:"column:liga_name"`
	CampeonatoName     string 		  `gorm:"column:campeonato_name"`
	ELocalName         string         `gorm:"column:e_local_name"`
	EVisitName         string         `gorm:"column:e_visit_name"`
	ResultadoLocal     sql.NullInt64  `gorm:"column:resultado_local"`
	ResultadoVisitante sql.NullInt64  `gorm:"column:resultado_visitante"`
	Suspendido         string `gorm:"column:suspendido"`
}