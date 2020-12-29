package gorms

type SancionesEquiposGorm struct {
	IDSanciones  int64 `gorm:"column:id_sanciones;primary_key"`
	IDEquipo     int64 `gorm:"column:id_equipo;primary_key"`
	IDCampeonato int64 `gorm:"column:id_campeonato;primary_key"`
}

// TableName sets the insert table name for this struct type
func (s *SancionesEquiposGorm) TableName() string {
	return "sanciones_equipos"
}
