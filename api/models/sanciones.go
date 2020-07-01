package models

type Sanciones struct {
	Descripcion   string `json:"descripcion"`
	IDLigas       int    `json:"id_ligas"`
	IDSanciones   int64  `json:"id_sanciones"`
	Observaciones string `json:"observaciones"`
}

// TableName sets the insert table name for this struct type
func (s *Sanciones) TableName() string {
	return "sanciones"
}
