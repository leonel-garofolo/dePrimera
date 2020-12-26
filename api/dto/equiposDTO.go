package models

import "database/sql"

type Equipos struct {
	Foto         []byte         `json:"foto"`
	Habilitado   sql.NullString `json:"habilitado" copier:"nopanic"`
	IDEquipo     int64          `json:"id_equipo"`
	IDCampeonato int            `json:"id_campeonato"`
	Nombre       string         `json:"nombre"`
	NroEquipo    int64          `json:"nro_equipo"`
}

// TableName sets the insert table name for this struct type
func (e *Equipos) TableName() string {
	return "equipos"
}
