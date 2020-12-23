package models

type Ligas struct {
	Cuit             string `json:"cuit"`
	Domicilio        string `json:"domicilio"`
	IDLiga           int64  `json:"idLiga"`
	MailContacto     string `json:"mailContacto"`
	Nombre           string `json:"nombre"`
	NombreContacto   string `json:"nombreContacto"`
	Telefono         string `json:"telefono"`
	TelefonoContacto string `json:"telefonoContacto"`
}

// TableName sets the insert table name for this struct type
func (l *Ligas) TableName() string {
	return "ligas"
}
