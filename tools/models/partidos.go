package models

type partidos struct {
	FechaEncuentro     time.Time   `gorm:"column:fecha_encuentro" json:"fecha_encuentro"`
	IDArbitro          null.Int    `gorm:"column:id_arbitro" json:"id_arbitro"`
	IDAsistente        null.Int    `gorm:"column:id_asistente" json:"id_asistente"`
	IDCampeonato       int         `gorm:"column:id_campeonato" json:"id_campeonato"`
	IDEquipoLocal      int         `gorm:"column:id_equipo_local" json:"id_equipo_local"`
	IDEquipoVisitante  int         `gorm:"column:id_equipo_visitante" json:"id_equipo_visitante"`
	IDLiga             int         `gorm:"column:id_liga" json:"id_liga"`
	IDPartidos         int         `gorm:"column:id_partidos;primary_key" json:"id_partidos"`
	MotivoSuspencion   null.String `gorm:"column:motivo_suspencion" json:"motivo_suspencion"`
	Observacion        null.String `gorm:"column:observacion" json:"observacion"`
	ResultadoLocal     null.Int    `gorm:"column:resultado_local" json:"resultado_local"`
	ResultadoVisitante null.Int    `gorm:"column:resultado_visitante" json:"resultado_visitante"`
	Suspendido         `gorm:"column:suspendido" json:"suspendido"`
}

// TableName sets the insert table name for this struct type
func (p *partidos) TableName() string {
	return "partidos"
}

// TableName sets the insert table name for this struct type
func (P *partidos) GetP() string {
	return "partidos"
}
