package models

type asistentes struct {
	IDAsistente int `gorm:"column:id_asistente;primary_key" json:"id_asistente"`
	IDPersona   int `gorm:"column:id_persona" json:"id_persona"`
}

// TableName sets the insert table name for this struct type
func (a *asistentes) TableName() string {
	return "asistentes"
}

// TableName sets the insert table name for this struct type
func (P *asistentes) GetP() string {
	return "asistentes"
}
