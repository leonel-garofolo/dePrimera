package models

type Sanciones struct {
	IDSanciones   int64  `json:"id_sanciones"`
	Descripcion   string `json:"descripcion"`
	Observaciones string `json:"observaciones"`
}

type SancionesJugadoresFromCampeonatoGorm struct {
	ApellidoNombre string `json:"apellido_nombre"`
	ENombre        string `json:"e_nombre"`
	CRojas         int    `json:"c_rojas"`
	CAmarillas     int    `json:"c_amarillas"`
	CAzules        int    `json:"c_azules"`
}

// TableName sets the insert table name for this struct type
func (s *Sanciones) TableName() string {
	return "sanciones"
}
