package gorms

type JugadoresGorm struct {
	IDJugador   int64 `gorm:"column:id_jugador;primary_key"`
	IDPersona   int64 `gorm:"column:id_persona"`
	IDEquipo    int64 `gorm:"column:id_equipo"`
	NroCamiseta int64 `gorm:"column:nro_camiseta"`
}

type JugadoresPlantelGorm struct {
	IDJugador   int64  `gorm:"column:id_jugador;primary_key"`
	Nombre      string `gorm:"column:nombre"`
	Apellido    string `gorm:"column:apellido"`
	NroCamiseta int64  `gorm:"column:nro_camiseta"`
}

// TableName sets the insert table name for this struct type
func (j *JugadoresGorm) TableName() string {
	return "jugadores"
}
