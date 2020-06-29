package models

type app_permisos struct {
	Descripcion null.String `gorm:"column:descripcion" json:"descripcion"`
	IDPermisos  int         `gorm:"column:id_permisos;primary_key" json:"id_permisos"`
}

// TableName sets the insert table name for this struct type
func (a *app_permisos) TableName() string {
	return "app_permisos"
}

// TableName sets the insert table name for this struct type
func (P *app_permisos) GetP() string {
	return "app_permisos"
}
