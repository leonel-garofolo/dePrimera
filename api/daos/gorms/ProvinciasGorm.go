package gorms

type ProvinciasGorm struct {
	IDProvincia int    `gorm:"column:id_provincia;primary_key"`
	IDPais      int    `gorm:"column:id_pais"`
	Nombre      string `gorm:"column:nombre"`
}

// TableName sets the insert table name for this struct type
func (a *ProvinciasGorm) TableName() string {
	return "app_provincias"
}
