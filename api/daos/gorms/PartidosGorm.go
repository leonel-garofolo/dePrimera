package gorms

import (
	"database/sql"
	"time"
)

type PartidosGorm struct {
	FechaEncuentro     time.Time      `gorm:"column:fecha_encuentro"`
	IDArbitro          sql.NullInt64  `gorm:"column:id_arbitro"`
	IDAsistente        sql.NullInt64  `gorm:"column:id_asistente"`
	IDCampeonato       int64            `gorm:"column:id_campeonato"`
	IDEquipoLocal      int64            `gorm:"column:id_equipo_local"`
	IDEquipoVisitante  int64            `gorm:"column:id_equipo_visitante"`
	IDLiga             int64            `gorm:"column:id_liga"`
	IDPartidos         int64            `gorm:"column:id_partidos;primary_key"`
	MotivoSuspencion   sql.NullString `gorm:"column:motivo_suspencion"`
	Observacion        sql.NullString `gorm:"column:observacion"`
	ResultadoLocal     sql.NullInt64  `gorm:"column:resultado_local"`
	ResultadoVisitante sql.NullInt64  `gorm:"column:resultado_visitante"`
	Suspendido         bool `gorm:"column:suspendido"`
}

// TableName sets the insert table name for this struct type
func (p *PartidosGorm) TableName() string {
	return "partidos"
}
