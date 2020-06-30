package models

import "database/sql"

type Eliminatorias struct {
	IDCampeonato   int           `gorm:"column:id_campeonato"`
	IDEliminatoria int           `gorm:"column:id_eliminatoria;primary_key"`
	IDPartido      int           `gorm:"column:id_partido"`
	NroLlave       sql.NullInt64 `gorm:"column:nro_llave"`
}

// TableName sets the insert table name for this struct type
func (e *Eliminatorias) TableName() string {
	return "eliminatorias"
}
