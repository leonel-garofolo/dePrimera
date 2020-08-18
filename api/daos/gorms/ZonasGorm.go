package gorms

import "database/sql"

type ZonasGorm struct {
	IDCampeonato int            `gorm:"column:id_campeonato"`
	IDZona       int            `gorm:"column:id_zona;primary_key"`
	Nombre       sql.NullString `gorm:"column:nombre"`
}

// TableName sets the insert table name for this struct type
func (z *ZonasGorm) TableName() string {
	return "zonas"
}
