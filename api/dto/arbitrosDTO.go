package models

type Arbitros struct {
	IDArbitro    int64 `json:"id_arbitro"`
	IDPersona    int64 `json:"id_persona"`
	IDCampeonato int64 `json:"id_campeonato"`
}

// TableName sets the insert table name for this struct type
func (a *Arbitros) TableName() string {
	return "arbitros"
}
