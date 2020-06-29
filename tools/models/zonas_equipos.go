package models

type zonas_equipos struct {
	IDEquipo int `gorm:"column:id_equipo;primary_key" json:"id_equipo"`
	IDZona   int `gorm:"column:id_zona;primary_key" json:"id_zona"`
}

// TableName sets the insert table name for this struct type
func (z *zonas_equipos) TableName() string {
	return "zonas_equipos"
}

// TableName sets the insert table name for this struct type
func (P *zonas_equipos) GetP() string {
	return "zonas_equipos"
}
