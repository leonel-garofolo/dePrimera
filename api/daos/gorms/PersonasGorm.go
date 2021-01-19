package gorms

import "database/sql"

type PersonasGorm struct {
	Nombre      string         `gorm:"column:nombre"`
	Apellido    string         `gorm:"column:apellido"`
	Domicilio   sql.NullString `gorm:"column:domicilio"`
	Edad        sql.NullInt64  `gorm:"column:edad"`
	Localidad   string         `gorm:"column:localidad"`
	IDPais      int64          `gorm:"column:id_pais"`
	IDPersona   int64          `gorm:"column:id_persona;primary_key"`
	IDProvincia int64          `gorm:"column:id_provincia"`
	IDTipoDoc   int            `gorm:"column:id_tipo_doc"`
	NroDoc      int            `gorm:"column:nro_doc"`
}

// TableName sets the insert table name for this struct type
func (p *PersonasGorm) TableName() string {
	return "personas"
}
