package models

type equipos_jugadores struct {
	IDEquipos   int `gorm:"column:id_equipos;primary_key" json:"id_equipos"`
	IDJugadores int `gorm:"column:id_jugadores;primary_key" json:"id_jugadores"`
}

// TableName sets the insert table name for this struct type
func (e *equipos_jugadores) TableName() string {
	return "equipos_jugadores"
}

// TableName sets the insert table name for this struct type
func (P *equipos_jugadores) GetP() string {
	return "equipos_jugadores"
}
