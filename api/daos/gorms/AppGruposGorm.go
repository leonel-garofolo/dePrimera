package gorms

type AppGruposGorm struct {
	Descripcion string `gorm:"column:descripcion"`
	Idgrupo     int    `gorm:"column:idgrupo;primary_key"`
}

// TableName sets the insert table name for this struct type
func (a *AppGruposGorm) TableName() string {
	return "app_grupos"
}
