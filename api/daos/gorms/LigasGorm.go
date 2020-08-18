package gorms

import "database/sql"

type LigasGorm struct {
	Cuit             sql.NullString `gorm:"column:cuit"`
	Domicilio        string         `gorm:"column:domicilio"`
	IDLiga           int            `gorm:"column:id_liga;primary_key"`
	MailContacto     sql.NullString `gorm:"column:mail_contacto"`
	Nombre           string         `gorm:"column:nombre"`
	NombreContacto   sql.NullString `gorm:"column:nombre_contacto"`
	Telefono         sql.NullString `gorm:"column:telefono"`
	TelefonoContacto sql.NullString `gorm:"column:telefono_contacto"`
}

// TableName sets the insert table name for this struct type
func (l *LigasGorm) TableName() string {
	return "ligas"
}
