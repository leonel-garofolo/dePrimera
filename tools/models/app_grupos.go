package models

type app_grupos struct {
	Descripcion null.String `gorm:"column:descripcion" json:"descripcion"`
	Idgrupo     int         `gorm:"column:idgrupo;primary_key" json:"idgrupo"`
}

// TableName sets the insert table name for this struct type
func (a *app_grupos) TableName() string {
	return "app_grupos"
}

// TableName sets the insert table name for this struct type
func (P *app_grupos) GetP() string {
	return "app_grupos"
}
