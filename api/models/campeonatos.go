package models

import (
	"time"
)

type Campeonatos struct {
	IDCampeonato int64     `json:"id_campeonato"`
	IDLiga       int       `json:"id_liga"`
	IDModelo     string    `json:"id_modelo"`
	Descripcion  string    `json:"descripcion"`
	FechaInicio  time.Time `json:"fecha_inicio"`
	FechaFin     time.Time `json:"fecha_fin"`
}

// TableName sets the insert table name for this struct type
func (c *Campeonatos) TableName() string {
	return "campeonatos"
}
