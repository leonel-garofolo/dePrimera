package gorms

type ZonasEquiposGorm struct {
	IDEquipo int64 `gorm:"column:id_equipo;primary_key"`
	IDZona   int64 `gorm:"column:id_zona;primary_key"`
}

// TableName sets the insert table name for this struct type
func (z *ZonasEquiposGorm) TableName() string {
	return "zonas_equipos"
}
