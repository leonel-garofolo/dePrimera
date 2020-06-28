package models

type asistentes struct {
	IDAsistente int `gorm:"column:id_asistente;primary_key"`
	IDPersona   int `gorm:"column:id_persona"`
}

// TableName sets the insert table name for this struct type
func (a *asistentes) TableName() string {
	return "asistentes"
}
