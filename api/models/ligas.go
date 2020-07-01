package models

type Ligas struct {
	Cuit             string `gorm:"column:cuit" json:"cuit"`
	Domicilio        string `gorm:"column:domicilio" json:"domicilio"`
	IDLiga           int64  `gorm:"column:id_liga;primary_key" json:"idLiga"`
	MailContacto     string `gorm:"column:mail_contacto" json:"mailContacto"`
	Nombre           string `gorm:"column:nombre" json:"nombre"`
	NombreContacto   string `gorm:"column:nombre_contacto" json:"nombreContacto"`
	Telefono         string `gorm:"column:telefono" json:"telefono"`
	TelefonoContacto string `gorm:"column:telefono_conctacto" json:"telefonoContacto"`
}

// TableName sets the insert table name for this struct type
func (l *Ligas) TableName() string {
	return "ligas"
}
