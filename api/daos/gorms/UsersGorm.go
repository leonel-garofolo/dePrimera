package gorms

import "database/sql"

type UsersGorm struct {
	Clave       sql.NullString `gorm:"column:clave"`
	Descripcion sql.NullString `gorm:"column:descripcion"`
	Email       sql.NullString `gorm:"column:email"`
	Habilitado  `gorm:"column:habilitado"`
	Telefono    sql.NullString `gorm:"column:telefono"`
	UserID      string         `gorm:"column:user_id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (u *UsersGorm) TableName() string {
	return "users"
}
