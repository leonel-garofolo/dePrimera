package models

type personas struct {
	ApellidoNombre string      `gorm:"column:apellido_nombre" json:"apellido_nombre"`
	Domicilio      null.String `gorm:"column:domicilio" json:"domicilio"`
	Edad           null.Int    `gorm:"column:edad" json:"edad"`
	IDLiga         null.Int    `gorm:"column:id_liga" json:"id_liga"`
	IDLocalidad    null.Int    `gorm:"column:id_localidad" json:"id_localidad"`
	IDPais         null.Int    `gorm:"column:id_pais" json:"id_pais"`
	IDPersona      int         `gorm:"column:id_persona;primary_key" json:"id_persona"`
	IDProvincia    null.Int    `gorm:"column:id_provincia" json:"id_provincia"`
	IDTipoDoc      int         `gorm:"column:id_tipo_doc" json:"id_tipo_doc"`
	NroDoc         int         `gorm:"column:nro_doc" json:"nro_doc"`
}

// TableName sets the insert table name for this struct type
func (p *personas) TableName() string {
	return "personas"
}

// TableName sets the insert table name for this struct type
func (P *personas) GetP() string {
	return "personas"
}
