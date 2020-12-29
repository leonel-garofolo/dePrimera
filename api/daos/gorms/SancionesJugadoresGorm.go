package gorms

type SancionesJugadoresGorm struct {
	IDSancionesJugadores int `gorm:"column:id_sanciones_jugadores;primary_key"`
	IDSanciones          int `gorm:"column:id_sancion"`
	IDJugador            int `gorm:"column:id_jugador"`
	IDCampeonato         int `gorm:"column:id_campeonato"`
}

type SancionesJugadoresFromCampeonatoGorm struct {
	ApellidoNombre string `gorm:"column:apellido_nombre"`
	ENombre        string `gorm:"column:e_nombre"`
	CRojas         int    `gorm:"column:c_rojas"`
	CAmarillas     int    `gorm:"column:c_amarillas"`
	CAzules        int    `gorm:"column:c_azules"`
}

// TableName sets the insert table name for this struct type
func (s *SancionesJugadoresGorm) TableName() string {
	return "sanciones_jugadores"
}
