package models

import (
	"time"
)

type Campeonatos struct {
	Descripcion  string    `json:"descripcion"`
	FechaFin     time.Time `json:"fecha_fin"`
	FechaInicio  time.Time `json:"fecha_inicio"`
	IDCampeonato int64     `json:"id_campeonato"`
	IDLiga       int       `json:"id_liga"`
	IDModelo     string    `json:"id_modelo"`
}

// TableName sets the insert table name for this struct type
func (c *Campeonatos) TableName() string {
	return "campeonatos"
}
