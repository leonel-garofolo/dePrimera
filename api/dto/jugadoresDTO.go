package models

type Jugadores struct {
	IDJugador int64 `json:"id_jugador"`
	IDPersona int64 `json:"id_persona"`
	IDEquipo  int64 `json:"id_equipo"`
}

// TableName sets the insert table name for this struct type
func (a *Jugadores) TableName() string {
	return "jugadores"
}
