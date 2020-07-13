package models

type AppGrupos struct {
	Idgrupo     int    `json:"id_grupo"`
	Descripcion string `json:"descripcion"`
}

// TableName sets the insert table name for this struct type
func (a *AppGrupos) TableName() string {
	return "app_grupos"
}
