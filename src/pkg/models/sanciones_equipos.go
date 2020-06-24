package models

type sanciones_equipos struct {
	IDEquipo    int `gorm:"column:id_equipo;primary_key"`
	IDSanciones int `gorm:"column:id_sanciones;primary_key"`
}

// TableName sets the insert table name for this struct type
func (s *sanciones_equipos) TableName() string {
	return "sanciones_equipos"
}
