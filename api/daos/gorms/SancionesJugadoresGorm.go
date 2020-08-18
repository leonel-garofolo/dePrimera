package gorms

type SancionesJugadoresGorm struct {
	IDJugador   int `gorm:"column:id_jugador;primary_key"`
	IDSanciones int `gorm:"column:id_sanciones;primary_key"`
}

// TableName sets the insert table name for this struct type
func (s *SancionesJugadoresGorm) TableName() string {
	return "sanciones_jugadores"
}
