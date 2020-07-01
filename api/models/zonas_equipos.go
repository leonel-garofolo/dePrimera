package models

type ZonasEquipos struct {
	IDEquipo int64 `json:"id_equipos"`
	IDZona   int64 `json:"id_zona"`
}

// TableName sets the insert table name for this struct type
func (z *ZonasEquipos) TableName() string {
	return "zonas_equipos"
}
