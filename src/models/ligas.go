package models

import (
	"database/sql"
)

type ligas struct {
	Cuit              sql.NullString `gorm:"column:cuit"`
	Domicilio         string         `gorm:"column:domicilio"`
	IDLiga            int            `gorm:"column:id_liga;primary_key"`
	MailContacto      sql.NullString `gorm:"column:mail_contacto"`
	Nombre            string         `gorm:"column:nombre"`
	NombreContacto    sql.NullString `gorm:"column:nombre_contacto"`
	Telefono          sql.NullString `gorm:"column:telefono"`
	TelefonoConctacto sql.NullString `gorm:"column:telefono_conctacto"`
}

// TableName sets the insert table name for this struct type
func (l *ligas) TableName() string {
	return "ligas"
}
