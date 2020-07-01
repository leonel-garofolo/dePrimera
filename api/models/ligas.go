package models

type Ligas struct {
	Cuit             string `db:"column:cuit" json:"cuit"`
	Domicilio        string `db:"column:domicilio" json:"domicilio"`
	IDLiga           int64  `db:"column:id_liga;primary_key" json:"idLiga"`
	MailContacto     string `db:"column:mail_contacto" json:"mailContacto"`
	Nombre           string `db:"column:nombre" json:"nombre"`
	NombreContacto   string `db:"column:nombre_contacto" json:"nombreContacto"`
	Telefono         string `db:"column:telefono" json:"telefono"`
	TelefonoContacto string `db:"column:telefono_conctacto" json:"telefonoContacto"`
}

// TableName sets the insert table name for this struct type
func (l *Ligas) TableName() string {
	return "ligas"
}
