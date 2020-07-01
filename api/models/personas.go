package models

type Personas struct {
	ApellidoNombre string `json:"apellido_nombre"`
	Domicilio      string `json:"domicilio"`
	Edad           int64  `json:"edad"`
	IDLiga         int    `json:"id_liga"`
	IDLocalidad    int64  `json:"id_localidad"`
	IDPais         int64  `json:"id_pais"`
	IDPersona      int64  `json:"id_persona"`
	IDProvincia    int64  `json:"id_provincia"`
	IDTipoDoc      int    `json:"id_tipo_doc"`
	NroDoc         int    `json:"nro_doc"`
}

// TableName sets the insert table name for this struct type
func (p *Personas) TableName() string {
	return "personas"
}
