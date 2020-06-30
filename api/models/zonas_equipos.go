package models

type ZonasEquipos struct {
	IDEquipo int `gorm:"column:id_equipo;primary_key"`
	IDZona   int `gorm:"column:id_zona;primary_key"`
}

// TableName sets the insert table name for this struct type
func (z *ZonasEquipos) TableName() string {
	return "zonas_equipos"
}
