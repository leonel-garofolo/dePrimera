package models

type arbitros struct {
	IDArbitro int `gorm:"column:id_arbitro;primary_key"`
	IDPersona int `gorm:"column:id_persona"`
}

// TableName sets the insert table name for this struct type
func (a *arbitros) TableName() string {
	return "arbitros"
}