package models

type partidos struct {
	FechaEncuentro     time.Time      `gorm:"column:fecha_encuentro"`
	IDArbitro          sql.NullInt64  `gorm:"column:id_arbitro"`
	IDAsistente        sql.NullInt64  `gorm:"column:id_asistente"`
	IDCampeonato       int            `gorm:"column:id_campeonato"`
	IDEquipoLocal      int            `gorm:"column:id_equipo_local"`
	IDEquipoVisitante  int            `gorm:"column:id_equipo_visitante"`
	IDLiga             int            `gorm:"column:id_liga"`
	IDPartidos         int            `gorm:"column:id_partidos;primary_key"`
	MotivoSuspencion   sql.NullString `gorm:"column:motivo_suspencion"`
	Observacion        sql.NullString `gorm:"column:observacion"`
	ResultadoLocal     sql.NullInt64  `gorm:"column:resultado_local"`
	ResultadoVisitante sql.NullInt64  `gorm:"column:resultado_visitante"`
	Suspendido         `gorm:"column:suspendido"`
}

// TableName sets the insert table name for this struct type
func (p *partidos) TableName() string {
	return "partidos"
}
