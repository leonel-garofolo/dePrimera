package models

type Provincias struct {
	IDProvincia int    `json:"id_provincia"`
	IDPais      int    `json:"id_pais"`
	Nombre      string `json:"nombre"`
}
