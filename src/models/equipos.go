package models

type Equipos struct {
	Foto []byte `gorm:"column:foto"; json:"foto"`
	//Habilitado `gorm:"column:habilitado"`
	IDEquipo int    `json:"idEquipo"; gorm:"column:id_equipo;primary_key"`
	IDLiga   int    `gorm:"column:id_liga"`
	Nombre   string `gorm:"column:nombre"`
}

// TableName sets the insert table name for this struct type
func (e *Equipos) TableName() string {
	return "equipos"
}

func NewEquipo(equipo int, liga int, nombre string, foto []byte) *Equipos {
	e := &Equipos{
		IDEquipo: equipo,
		IDLiga:   liga,
		Nombre:   nombre,
		Foto:     foto,
	}
	return e
}
