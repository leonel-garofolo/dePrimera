package gorms

type EquiposJugadoresGorm struct {
	IDEquipos   int64 `gorm:"column:id_equipos;primary_key"`
	IDJugadores int64 `gorm:"column:id_jugadores;primary_key"`
}

// TableName sets the insert table name for this struct type
func (e *EquiposJugadoresGorm) TableName() string {
	return "equipos_jugadores"
}
