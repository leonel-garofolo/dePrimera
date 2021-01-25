package models

import (
	"time"
)

type Campeonatos struct {
	IDCampeonato     int64     `json:"id_campeonato"`
	IDLiga           int       `json:"id_liga"`
	IDModelo         string    `json:"id_modelo"`
	Descripcion      string    `json:"descripcion"`
	FechaInicio      time.Time `json:"fecha_inicio"`
	FechaFin         time.Time `json:"fecha_fin"`
	GenFixture       bool      `json:"gen_fixture"`
	GenFixtureFinish bool      `json:"gen_fixture_finish"`
}

type CampeonatosGoleadores struct {
	IDJugador int64  `json:"id_campeonato"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Equipo    string `json:"equipo"`
	Goles     int64  `json:"goles"`
}

// TableName sets the insert table name for this struct type
func (c *Campeonatos) TableName() string {
	return "campeonatos"
}
