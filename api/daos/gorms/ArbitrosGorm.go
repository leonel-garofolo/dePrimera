package gorms

type ArbitrosGorm struct {
	IDArbitro int64 `gorm:"column:id_arbitro;primary_key"`
	IDPersona int64 `gorm:"column:id_persona"`
}

// TableName sets the insert table name for this struct type
func (a *ArbitrosGorm) TableName() string {
	return "arbitros"
}
