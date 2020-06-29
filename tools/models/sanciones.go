package models

type sanciones struct {
	Descripcion   null.String `gorm:"column:descripcion" json:"descripcion"`
	IDLigas       int         `gorm:"column:id_ligas" json:"id_ligas"`
	IDSanciones   int         `gorm:"column:id_sanciones;primary_key" json:"id_sanciones"`
	Observaciones null.String `gorm:"column:observaciones" json:"observaciones"`
}

// TableName sets the insert table name for this struct type
func (s *sanciones) TableName() string {
	return "sanciones"
}

// TableName sets the insert table name for this struct type
func (P *sanciones) GetP() string {
	return "sanciones"
}
