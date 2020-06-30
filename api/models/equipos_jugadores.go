package models

type EquiposJugadores struct {
	IDEquipos   int `gorm:"column:id_equipos;primary_key"`
	IDJugadores int `gorm:"column:id_jugadores;primary_key"`
}

// TableName sets the insert table name for this struct type
func (e *EquiposJugadores) TableName() string {
	return "equipos_jugadores"
}
