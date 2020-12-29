package gorms

import "database/sql"

type SancionesGorm struct {
	IDSanciones   int64          `gorm:"column:id_sanciones;primary_key"`
	Descripcion   sql.NullString `gorm:"column:descripcion"`
	Observaciones sql.NullString `gorm:"column:observaciones"`
}

// TableName sets the insert table name for this struct type
func (s *SancionesGorm) TableName() string {
	return "sanciones"
}
