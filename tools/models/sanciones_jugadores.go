package models

type sanciones_jugadores struct {
	IDJugador   int `gorm:"column:id_jugador;primary_key" json:"id_jugador"`
	IDSanciones int `gorm:"column:id_sanciones;primary_key" json:"id_sanciones"`
}

// TableName sets the insert table name for this struct type
func (s *sanciones_jugadores) TableName() string {
	return "sanciones_jugadores"
}

// TableName sets the insert table name for this struct type
func (P *sanciones_jugadores) GetP() string {
	return "sanciones_jugadores"
}
