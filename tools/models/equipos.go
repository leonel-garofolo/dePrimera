package models

type equipos struct {
	Foto       []byte `gorm:"column:foto" json:"foto"`
	Habilitado `gorm:"column:habilitado" json:"habilitado"`
	IDEquipo   int    `gorm:"column:id_equipo;primary_key" json:"id_equipo"`
	IDLiga     int    `gorm:"column:id_liga" json:"id_liga"`
	Nombre     string `gorm:"column:nombre" json:"nombre"`
}

// TableName sets the insert table name for this struct type
func (e *equipos) TableName() string {
	return "equipos"
}

// TableName sets the insert table name for this struct type
func (P *equipos) GetP() string {
	return "equipos"
}
