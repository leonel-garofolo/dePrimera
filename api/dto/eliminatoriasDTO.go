package models

import "database/sql"

type Eliminatorias struct {
	IDCampeonato   int           `json:"id_campeonato"`
	IDEliminatoria int64         `json:"id_eliminatoria"`
	IDPartido      int           `json:"id_partido"`
	NroLlave       sql.NullInt64 `json:"nro_llave"`
}

// TableName sets the insert table name for this struct type
func (e *Eliminatorias) TableName() string {
	return "eliminatorias"
}
