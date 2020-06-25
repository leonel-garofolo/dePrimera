package models

import (
	"database/sql"
)

type Ligas struct {
	Cuit              sql.NullString `json:"cuit"; gorm:"column:cuit"`
	Domicilio         string         `json:"domicilio"; gorm:"column:domicilio"`
	IDLiga            int            `json:"idLiga"; gorm:"column:id_liga;primary_key"`
	MailContacto      sql.NullString `json:"mailContacto"; gorm:"column:mail_contacto"`
	Nombre            string         `json:"nombre"; gorm:"column:nombre"`
	NombreContacto    sql.NullString `json:"nombreContacto"; gorm:"column:nombre_contacto"`
	Telefono          sql.NullString `json:"telefono"; gorm:"column:telefono"`
	TelefonoConctacto sql.NullString `json:"telefonoConctacto"; gorm:"column:telefono_conctacto"`
}

// TableName sets the insert table name for this struct type
func (l *Ligas) TableName() string {
	return "ligas"
}
