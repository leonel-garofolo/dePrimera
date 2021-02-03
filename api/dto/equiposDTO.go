package models

type Equipos struct {
	Foto         []byte `json:"foto"`
	Habilitado   bool   `json:"habilitado" copier:"nopanic"`
	IDEquipo     int64  `json:"id_equipo"`
	Nombre       string `json:"nombre"`
	NroEquipo    int64  `json:"nro_equipo" copier:"nopanic"`
	IDCampeonato int64  `json:"id_campeonato" copier:"nopanic"`
}

type EquiposTablePos struct {
	IDEquipo        int64  `json:"id_equipo"`
	IDCampeonato    int64  `json:"id_campeonato"`
	Nombre          string `json:"nombre"`
	NroEquipo       int    `json:"nro_equipo"`
	Puntos          int    `json:"puntos"`
	PartidoGanado   int    `json:"p_gan"`
	PartidoEmpatado int    `json:"p_emp"`
	PartidoPerdido  int    `json:"p_per"`
}

// TableName sets the insert table name for this struct type
func (e *Equipos) TableName() string {
	return "equipos"
}
