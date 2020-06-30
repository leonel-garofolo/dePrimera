package models

import "database/sql"

type Sanciones struct {
	Descripcion   sql.NullString `gorm:"column:descripcion"`
	IDLigas       int            `gorm:"column:id_ligas"`
	IDSanciones   int            `gorm:"column:id_sanciones;primary_key"`
	Observaciones sql.NullString `gorm:"column:observaciones"`
}

// TableName sets the insert table name for this struct type
func (s *Sanciones) TableName() string {
	return "sanciones"
}
