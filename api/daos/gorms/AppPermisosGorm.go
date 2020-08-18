package gorms

import "database/sql"

type AppPermisosGorm struct {
	Descripcion sql.NullString `gorm:"column:descripcion"`
	IDPermisos  int            `gorm:"column:id_permisos;primary_key"`
}

// TableName sets the insert table name for this struct type
func (a *AppPermisosGorm) TableName() string {
	return "app_permisos"
}
