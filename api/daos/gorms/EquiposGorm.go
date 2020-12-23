package gorms

import "database/sql"

type EquiposGorm struct {
	Foto       []byte        `gorm:"column:foto"`
	Habilitado sql.NullInt64 `gorm:"column:habilitado"`
	IDEquipo   int64           `gorm:"column:id_equipo;primary_key"`
	IDLiga     int64           `gorm:"column:id_liga"`
	Nombre     string        `gorm:"column:nombre"`
}

// TableName sets the insert table name for this struct type
func (e *EquiposGorm) TableName() string {
	return "equipos"
}
