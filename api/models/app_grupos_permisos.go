package models

import "database/sql"

type app_grupos_permisos struct {
	IDGrupo   int           `gorm:"column:id_grupo;primary_key"`
	IDPermiso sql.NullInt64 `gorm:"column:id_permiso"`
}

// TableName sets the insert table name for this struct type
func (a *app_grupos_permisos) TableName() string {
	return "app_grupos_permisos"
}
