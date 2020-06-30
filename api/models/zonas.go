package models

import "database/sql"

type Zonas struct {
	IDCampeonato int            `gorm:"column:id_campeonato"`
	IDZona       int            `gorm:"column:id_zona;primary_key"`
	Nombre       sql.NullString `gorm:"column:nombre"`
}

// TableName sets the insert table name for this struct type
func (z *Zonas) TableName() string {
	return "zonas"
}
