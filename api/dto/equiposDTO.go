package models

type Equipos struct {
	Foto         []byte `json:"foto"`
	Habilitado   bool   `json:"habilitado" copier:"nopanic"`
	IDEquipo     int64  `json:"id_equipo"`
	IDCampeonato int    `json:"id_campeonato"`
	Nombre       string `json:"nombre"`
}

// TableName sets the insert table name for this struct type
func (e *Equipos) TableName() string {
	return "equipos"
}
