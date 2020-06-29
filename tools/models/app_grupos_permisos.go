package models

type app_grupos_permisos struct {
	IDGrupo   int      `gorm:"column:id_grupo;primary_key" json:"id_grupo"`
	IDPermiso null.Int `gorm:"column:id_permiso" json:"id_permiso"`
}

// TableName sets the insert table name for this struct type
func (a *app_grupos_permisos) TableName() string {
	return "app_grupos_permisos"
}

// TableName sets the insert table name for this struct type
func (P *app_grupos_permisos) GetP() string {
	return "app_grupos_permisos"
}
