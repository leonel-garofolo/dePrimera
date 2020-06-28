package models

import "database/sql"

type app_grupos struct {
	Descripcion sql.NullString `gorm:"column:descripcion"`
	Idgrupo     int            `gorm:"column:idgrupo;primary_key"`
}

// TableName sets the insert table name for this struct type
func (a *app_grupos) TableName() string {
	return "app_grupos"
}
