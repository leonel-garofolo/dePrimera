package gorms

type JugadoresGorm struct {
	IDJugador int64 `gorm:"column:id_jugador;primary_key"`
	IDPersona int64 `gorm:"column:id_persona"`
	IDEquipo int64  `gorm:"column:id_equipo"`
}

// TableName sets the insert table name for this struct type
func (j *JugadoresGorm) TableName() string {
	return "jugadores"
}
