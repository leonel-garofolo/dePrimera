package models

import "database/sql"

type app_permisos struct {
	Descripcion sql.NullString `gorm:"column:descripcion"`
	IDPermisos  int            `gorm:"column:id_permisos;primary_key"`
}

// TableName sets the insert table name for this struct type
func (a *app_permisos) TableName() string {
	return "app_permisos"
}
