package daos

import "deprimera/api/models"

type DaosFactory interface {
	GetEquiposDao() *EquiposDao
}

type DePrimeraDaos struct{}

type EquiposDao interface {
	GetAll() []models.Equipos
	Get(id int) *models.Equipos
	Save() int
	Delete(id int) bool
	Query(filter string) []models.Equipos
}

func NewDePrimeraDaos() *DePrimeraDaos {
	return &DePrimeraDaos{}
}

func (dao *DePrimeraDaos) GetEquiposDao() *EquiposDaoImpl {
	return &EquiposDaoImpl{}
}
