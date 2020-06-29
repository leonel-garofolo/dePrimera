package models

type arbitros struct {
	IDArbitro int `gorm:"column:id_arbitro;primary_key" json:"id_arbitro"`
	IDPersona int `gorm:"column:id_persona" json:"id_persona"`
}

// TableName sets the insert table name for this struct type
func (a *arbitros) TableName() string {
	return "arbitros"
}

// TableName sets the insert table name for this struct type
func (P *arbitros) GetP() string {
	return "arbitros"
}
