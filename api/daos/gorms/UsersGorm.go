package gorms

import "database/sql"

type UsersGorm struct {
	UserID     string         `gorm:"column:id_user;primary_key"`
	Password   string         `gorm:"column:clave"`
	Habilitado bool           `gorm:"column:habilitado"`
	Telefono   sql.NullString `gorm:"column:telefono"`
}

// TableName sets the insert table name for this struct type
func (u *UsersGorm) TableName() string {
	return "users"
}
