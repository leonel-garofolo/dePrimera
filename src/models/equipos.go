package models

type Equipos struct {
	Foto []byte `json:"foto"; gorm:"column:foto"`
	//Habilitado    `gorm:"column:habilitado"`
	IDEquipo int    `json:"idEquipo"; gorm:"column:id_equipo;primary_key"`
	IDLiga   int    `json:"idLiga"; gorm:"column:id_liga"`
	Nombre   string `json:"nombre"; gorm:"column:nombre"`
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
