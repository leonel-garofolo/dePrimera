package daos

import "deprimera/api/models"

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
	Save(e models.Arbitros) int
	Delete(id int) bool
	Query(filter string) []models.Arbitros
}

func (dao *DePrimeraDaos) GetArbitrosDao() *ArbitrosDaoImpl {
	return &ArbitrosDaoImpl{}
}

type AsistentesDao interface {
	GetAll() []models.Asistentes
	Get(id int) *models.Asistentes
	Save(e models.Asistentes) int
	Delete(id int) bool
	Query(filter string) []models.Asistentes
}

func (dao *DePrimeraDaos) GetAsistentesDao() *AsistentesDaoImpl {
	return &AsistentesDaoImpl{}
}

type CampeonatosDao interface {
	GetAll() []models.Campeonatos
	Get(id int) *models.Campeonatos
	Save(e models.Campeonatos) int
	Delete(id int) bool
	Query(filter string) []models.Campeonatos
}

func (dao *DePrimeraDaos) GetCampeonatosDao() *CampeonatosDaoImpl {
	return &CampeonatosDaoImpl{}
}

type EliminatoriasDao interface {
	GetAll() []models.Eliminatorias
	Get(id int) *models.Eliminatorias
	Save(e models.Eliminatorias) int
	Delete(id int) bool
	Query(filter string) []models.Eliminatorias
}

func (dao *DePrimeraDaos) GetEliminatoriasDao() *EliminatoriasDaoImpl {
	return &EliminatoriasDaoImpl{}
}

type EquiposDao interface {
	GetAll() []models.Equipos
	Get(id int) *models.Equipos
	Save(e models.Equipos) int
	Delete(id int) bool
	Query(filter string) []models.Equipos
}

func (dao *DePrimeraDaos) GetEquiposDao() *EquiposDaoImpl {
	return &EquiposDaoImpl{}
}

type LigasDao interface {
	GetAll() []models.Ligas
	Get(id int) *models.Ligas
	Save(e models.Ligas) int
	Delete(id int) bool
	Query(filter string) []models.Ligas
}

func (dao *DePrimeraDaos) GetLigasDao() *LigasDaoImpl {
	return &LigasDaoImpl{}
}
