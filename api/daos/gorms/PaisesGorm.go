package gorms

type PaisesGorm struct {
	IDPais int    `gorm:"column:id_pais;primary_key"`
	Nombre string `gorm:"column:nombre"`
}

// TableName sets the insert table name for this struct type
func (a *PaisesGorm) TableName() string {
	return "app_paises"
}
