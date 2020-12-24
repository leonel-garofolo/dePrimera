package gorms

import "database/sql"

type UsersGorm struct {
	IDUser      string  `gorm:"column:id_user;primary_key"`
	Clave       string 	`gorm:"column:clave"`
	Nombre      string 	`gorm:"column:nombre"`
	Apellido    string 	`gorm:"column:apellido"`	
	Habilitado  bool 	`gorm:"column:habilitado"`
	Telefono    sql.NullString `gorm:"column:telefono"`	
}

// TableName sets the insert table name for this struct type
func (u *UsersGorm) TableName() string {
	return "users"
}
