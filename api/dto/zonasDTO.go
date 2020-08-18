package models

type Zonas struct {
	IDCampeonato int    `json:"id_campeonato"`
	IDZona       int64  `json:"id_zona"`
	Nombre       string `json:"nombre"`
}

// TableName sets the insert table name for this struct type
func (z *Zonas) TableName() string {
	return "zonas"
}
