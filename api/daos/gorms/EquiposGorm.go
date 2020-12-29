package gorms

import "database/sql"

type EquiposGorm struct {
	Foto         []byte         `gorm:"column:foto"`
	Habilitado   sql.NullString `gorm:"column:habilitado"`
	IDEquipo     int64          `gorm:"column:id_equipo;primary_key"`
	IDCampeonato int64          `gorm:"column:id_campeonato"`
	Nombre       string         `gorm:"column:nombre"`
	NroEquipo    int64          `gorm:"column:nro_equipo"`
}

type EquiposTablePosGorm struct {
	IDEquipo        int64  `gorm:"column:id_equipo"`
	IDCampeonato    int64  `gorm:"column:id_campeonato"`
	Nombre          string `gorm:"column:nombre"`
	NroEquipo       string `gorm:"column:nro_equipo"`
	Puntos          int    `gorm:"column:puntos"`
	PartidoGanado   int    `gorm:"column:p_gan"`
	PartidoEmpatado int    `gorm:"column:p_emp"`
	PartidoPerdido  int    `gorm:"column:p_per"`
}

// TableName sets the insert table name for this struct type
func (e *EquiposGorm) TableName() string {
	return "equipos"
}
