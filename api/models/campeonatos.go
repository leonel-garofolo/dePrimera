package models

import (
	"database/sql"
	"time"
)

type Campeonatos struct {
	Descripcion  sql.NullString `gorm:"column:descripcion"`
	FechaFin     time.Time      `gorm:"column:fecha_fin"`
	FechaInicio  time.Time      `gorm:"column:fecha_inicio"`
	IDCampeonato int            `gorm:"column:id_campeonato;primary_key"`
	IDLiga       int            `gorm:"column:id_liga"`
	IDModelo     sql.NullString `gorm:"column:id_modelo"`
}

// TableName sets the insert table name for this struct type
func (c *Campeonatos) TableName() string {
	return "campeonatos"
}
