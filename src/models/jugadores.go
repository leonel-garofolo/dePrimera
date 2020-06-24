package models

import (
	"database/sql"
)

type jugadores struct {
	Foto        []byte        `gorm:"column:foto"`
	IDJugadores int           `gorm:"column:id_jugadores;primary_key"`
	IDPersona   int           `gorm:"column:id_persona"`
	Puesto      sql.NullInt64 `gorm:"column:puesto"`
}

// TableName sets the insert table name for this struct type
func (j *jugadores) TableName() string {
	return "jugadores"
}
