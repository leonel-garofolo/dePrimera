package models

type Sanciones struct {
	IDSanciones   int64  `json:"id_sanciones"`
	Descripcion   string `json:"descripcion"`
	Observaciones string `json:"observaciones"`
}

type SancionesJugadoresFromCampeonato struct {
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	ENombre    string `json:"e_nombre"`
	CRojas     int    `json:"c_rojas"`
	CAmarillas int    `json:"c_amarillas"`
	CAzules    int    `json:"c_azules"`
}

// TableName sets the insert table name for this struct type
func (s *Sanciones) TableName() string {
	return "sanciones"
}
