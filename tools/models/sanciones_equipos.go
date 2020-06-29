package models

type sanciones_equipos struct {
	IDEquipo    int `gorm:"column:id_equipo;primary_key" json:"id_equipo"`
	IDSanciones int `gorm:"column:id_sanciones;primary_key" json:"id_sanciones"`
}

// TableName sets the insert table name for this struct type
func (s *sanciones_equipos) TableName() string {
	return "sanciones_equipos"
}

// TableName sets the insert table name for this struct type
func (P *sanciones_equipos) GetP() string {
	return "sanciones_equipos"
}
