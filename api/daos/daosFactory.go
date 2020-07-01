package daos

import (
	"deprimera/api/models"
)

type DaosFactory interface {
	GetEquiposDao() *EquiposDao
}

type DePrimeraDaos struct{}

func NewDePrimeraDaos() *DePrimeraDaos {
	return &DePrimeraDaos{}
}

type ArbitrosDao interface {
	GetAll() []models.Arbitros
	Get(id int) *models.Arbitros
	Save(e *models.Arbitros) int
	Delete(id int) bool
	Query(filter string) []models.Arbitros
}

func (dao *DePrimeraDaos) GetArbitrosDao() *ArbitrosDaoImpl {
	return &ArbitrosDaoImpl{}
}

type AsistentesDao interface {
	GetAll() []models.Asistentes
	Get(id int) *models.Asistentes
	Save(e *models.Asistentes) int
	Delete(id int) bool
	Query(filter string) []models.Asistentes
}

func (dao *DePrimeraDaos) GetAsistentesDao() *AsistentesDaoImpl {
	return &AsistentesDaoImpl{}
}

type CampeonatosDao interface {
	GetAll() []models.Campeonatos
	Get(id int) *models.Campeonatos
	Save(e *models.Campeonatos) int
	Delete(id int) bool
	Query(filter string) []models.Campeonatos
}

func (dao *DePrimeraDaos) GetCampeonatosDao() *CampeonatosDaoImpl {
	return &CampeonatosDaoImpl{}
}

type EliminatoriasDao interface {
	GetAll() []models.Eliminatorias
	Get(id int) *models.Eliminatorias
	Save(e *models.Eliminatorias) int
	Delete(id int) bool
	Query(filter string) []models.Eliminatorias
}

func (dao *DePrimeraDaos) GetEliminatoriasDao() *EliminatoriasDaoImpl {
	return &EliminatoriasDaoImpl{}
}

type EquiposDao interface {
	GetAll() []models.Equipos
	Get(id int) *models.Equipos
	Save(e *models.Equipos) int
	Delete(id int) bool
	Query(filter string) []models.Equipos
}

func (dao *DePrimeraDaos) GetEquiposDao() *EquiposDaoImpl {
	return &EquiposDaoImpl{}
}

type EquiposJugadoresDao interface {
	Save(e *models.EquiposJugadores) int64
	Delete(id int64) bool
}

func (dao *DePrimeraDaos) GetEquiposJugadoresDao() *EquiposJugadoresDaoImpl {
	return &EquiposJugadoresDaoImpl{}
}

type LigasDao interface {
	GetAll() []models.Ligas
	Get(id int) *models.Ligas
	Save(e *models.Ligas) int
	Delete(id int) bool
	Query(filter string) []models.Ligas
}

func (dao *DePrimeraDaos) GetLigasDao() *LigasDaoImpl {
	return &LigasDaoImpl{}
}

type PartidosDao interface {
	GetAll() []models.Partidos
	Get(id int) *models.Partidos
	Save(e *models.Partidos) int
	Delete(id int) bool
	Query(filter string) []models.Partidos
}

func (dao *DePrimeraDaos) GetPartidosDao() *PartidosDaoImpl {
	return &PartidosDaoImpl{}
}

type PersonasDao interface {
	GetAll() []models.Personas
	Get(id int) *models.Personas
	Save(e *models.Personas) int
	Delete(id int) bool
	Query(filter string) []models.Personas
}

func (dao *DePrimeraDaos) GetPersonasDao() *PersonasDaoImpl {
	return &PersonasDaoImpl{}
}

type SancionesDao interface {
	GetAll() []models.Sanciones
	Get(id int) *models.Sanciones
	Save(e *models.Sanciones) int
	Delete(id int) bool
	Query(filter string) []models.Sanciones
}

func (dao *DePrimeraDaos) GetSancionesDao() *SancionesDaoImpl {
	return &SancionesDaoImpl{}
}

type SancionesEquiposDao interface {
	Save(e *models.SancionesEquipos) int64
	Delete(id int64) bool
}

func (dao *DePrimeraDaos) GetSancionesEquiposDao() *SancionesEquiposDaoImpl {
	return &SancionesEquiposDaoImpl{}
}

type ZonasDao interface {
	GetAll() []models.Zonas
	Get(id int) *models.Zonas
	Save(e *models.Zonas) int
	Delete(id int) bool
	Query(filter string) []models.Zonas
}

func (dao *DePrimeraDaos) GetZonasDao() *ZonasDaoImpl {
	return &ZonasDaoImpl{}
}

type ZonasEquiposDao interface {
	Save(e *models.ZonasEquipos) int64
	Delete(id int64) bool
}

func (dao *DePrimeraDaos) GetZonasEquiposDao() *ZonasEquiposDaoImpl {
	return &ZonasEquiposDaoImpl{}
}
