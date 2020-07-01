package models

type SancionesEquipos struct {
	IDEquipo    int64 `json:"id_equipo"`
	IDSanciones int64 `json:"id_sanciones"`
}

// TableName sets the insert table name for this struct type
func (s *SancionesEquipos) TableName() string {
	return "sanciones_equipos"
}
