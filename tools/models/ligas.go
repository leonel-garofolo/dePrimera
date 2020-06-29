package models

type ligas struct {
	Cuit              null.String `gorm:"column:cuit" json:"cuit"`
	Domicilio         string      `gorm:"column:domicilio" json:"domicilio"`
	IDLiga            int         `gorm:"column:id_liga;primary_key" json:"id_liga"`
	MailContacto      null.String `gorm:"column:mail_contacto" json:"mail_contacto"`
	Nombre            string      `gorm:"column:nombre" json:"nombre"`
	NombreContacto    null.String `gorm:"column:nombre_contacto" json:"nombre_contacto"`
	Telefono          null.String `gorm:"column:telefono" json:"telefono"`
	TelefonoConctacto null.String `gorm:"column:telefono_conctacto" json:"telefono_conctacto"`
}

// TableName sets the insert table name for this struct type
func (l *ligas) TableName() string {
	return "ligas"
}

// TableName sets the insert table name for this struct type
func (P *ligas) GetP() string {
	return "ligas"
}
