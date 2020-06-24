package models

type equipos_jugadores struct {
	IDEquipos   int `gorm:"column:id_equipos;primary_key"`
	IDJugadores int `gorm:"column:id_jugadores;primary_key"`
}

// TableName sets the insert table name for this struct type
func (e *equipos_jugadores) TableName() string {
	return "equipos_jugadores"
}
