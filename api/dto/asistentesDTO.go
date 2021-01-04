package models

type Asistentes struct {
	IDAsistente  int64 `json:"id_asistente"`
	IDPersona    int64 `json:"id_persona"`
	IDCampeonato int64 `json:"id_campeonato"`
}

// TableName sets the insert table name for this struct type
func (a *Asistentes) TableName() string {
	return "asistentes"
}
