package models

type campeonatos struct {
	Descripcion  null.String `gorm:"column:descripcion" json:"descripcion"`
	FechaFin     null.Time   `gorm:"column:fecha_fin" json:"fecha_fin"`
	FechaInicio  null.Time   `gorm:"column:fecha_inicio" json:"fecha_inicio"`
	IDCampeonato int         `gorm:"column:id_campeonato;primary_key" json:"id_campeonato"`
	IDLiga       int         `gorm:"column:id_liga" json:"id_liga"`
	IDModelo     null.String `gorm:"column:id_modelo" json:"id_modelo"`
}

// TableName sets the insert table name for this struct type
func (c *campeonatos) TableName() string {
	return "campeonatos"
}

// TableName sets the insert table name for this struct type
func (P *campeonatos) GetP() string {
	return "campeonatos"
}
