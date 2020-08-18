package models

type EquiposJugadores struct {
	IDEquipos   int64 `json:"id_equipos"`
	IDJugadores int64 `json:"id_jugadores"`
}

// TableName sets the insert table name for this struct type
func (e *EquiposJugadores) TableName() string {
	return "equipos_jugadores"
}
