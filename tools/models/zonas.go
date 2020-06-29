package models

type zonas struct {
	IDCampeonato int         `gorm:"column:id_campeonato" json:"id_campeonato"`
	IDZona       int         `gorm:"column:id_zona;primary_key" json:"id_zona"`
	Nombre       null.String `gorm:"column:nombre" json:"nombre"`
}

// TableName sets the insert table name for this struct type
func (z *zonas) TableName() string {
	return "zonas"
}

// TableName sets the insert table name for this struct type
func (P *zonas) GetP() string {
	return "zonas"
}
