package models

type eliminatorias struct {
	IDCampeonato   int      `gorm:"column:id_campeonato" json:"id_campeonato"`
	IDEliminatoria int      `gorm:"column:id_eliminatoria;primary_key" json:"id_eliminatoria"`
	IDPartido      int      `gorm:"column:id_partido" json:"id_partido"`
	NroLlave       null.Int `gorm:"column:nro_llave" json:"nro_llave"`
}

// TableName sets the insert table name for this struct type
func (e *eliminatorias) TableName() string {
	return "eliminatorias"
}

// TableName sets the insert table name for this struct type
func (P *eliminatorias) GetP() string {
	return "eliminatorias"
}
