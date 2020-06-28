package models

type equipos struct {
	Foto       []byte `gorm:"column:foto"`
	Habilitado bool   `gorm:"column:habilitado"`
	IDEquipo   int    `gorm:"column:id_equipo;primary_key"`
	IDLiga     int    `gorm:"column:id_liga"`
	Nombre     string `gorm:"column:nombre"`
}

// TableName sets the insert table name for this struct type
func (e *equipos) TableName() string {
	return "equipos"
}
