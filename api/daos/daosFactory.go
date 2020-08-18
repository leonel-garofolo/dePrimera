package daos

import "deprimera/api/daos/gorms"

type DaosFactory interface {
	GetEquiposDao() *EquiposDao
}

type DePrimeraDaos struct{}

func NewDePrimeraDaos() *DePrimeraDaos {
	return &DePrimeraDaos{}
}

type ArbitrosDao interface {
	GetAll() []gorms.ArbitrosGorm
	Get(id int) *gorms.ArbitrosGorm
	Save(e *gorms.ArbitrosGorm) int
	Delete(id int) bool
	Query(filter string) []gorms.ArbitrosGorm
}

func (dao *DePrimeraDaos) GetArbitrosDao() *ArbitrosDaoImpl {
	return &ArbitrosDaoImpl{}
}

type AsistentesDao interface {
	GetAll() []gorms.AsistentesGorm
	Get(id int) *gorms.AsistentesGorm
	Save(e *gorms.AsistentesGorm) int
	Delete(id int) bool
	Query(filter string) []gorms.AsistentesGorm
}

func (dao *DePrimeraDaos) GetAsistentesDao() *AsistentesDaoImpl {
	return &AsistentesDaoImpl{}
}

type CampeonatosDao interface {
	GetAll() []gorms.CampeonatosGorm
	Get(id int) *gorms.CampeonatosGorm
	Save(e *gorms.CampeonatosGorm) int
	Delete(id int) bool
	Query(filter string) []gorms.CampeonatosGorm
}

func (dao *DePrimeraDaos) GetCampeonatosDao() *CampeonatosDaoImpl {
	return &CampeonatosDaoImpl{}
}

type EliminatoriasDao interface {
	GetAll() []gorms.EliminatoriasGorm
	Get(id int) *gorms.EliminatoriasGorm
	Save(e *gorms.EliminatoriasGorm) int
	Delete(id int) bool
	Query(filter string) []gorms.EliminatoriasGorm
}

func (dao *DePrimeraDaos) GetEliminatoriasDao() *EliminatoriasDaoImpl {
	return &EliminatoriasDaoImpl{}
}

type EquiposDao interface {
	GetAll() []gorms.EquiposGorm
	Get(id int) *gorms.EquiposGorm
	Save(e *gorms.EquiposGorm) int
	Delete(id int) bool
	Query(filter string) []gorms.EquiposGorm
}

func (dao *DePrimeraDaos) GetEquiposDao() *EquiposDaoImpl {
	return &EquiposDaoImpl{}
}

type EquiposJugadoresDao interface {
	Save(e *gorms.EquiposJugadoresGorm) int64
	Delete(id int64) bool
}

func (dao *DePrimeraDaos) GetEquiposJugadoresDao() *EquiposJugadoresDaoImpl {
	return &EquiposJugadoresDaoImpl{}
}

type LigasDao interface {
	GetAll() []gorms.LigasGorm
	Get(id int) *gorms.LigasGorm
	Save(e *gorms.LigasGorm) int
	Delete(id int) bool
	Query(filter string) []gorms.LigasGorm
}

func (dao *DePrimeraDaos) GetLigasDao() *LigasDaoImpl {
	return &LigasDaoImpl{}
}

type PartidosDao interface {
	GetAll() []gorms.PartidosGorm
	Get(id int) *gorms.PartidosGorm
	Save(e *gorms.PartidosGorm) int
	Delete(id int) bool
	Query(filter string) []gorms.PartidosGorm
}

func (dao *DePrimeraDaos) GetPartidosDao() *PartidosDaoImpl {
	return &PartidosDaoImpl{}
}

type PersonasDao interface {
	GetAll() []gorms.PersonasGorm
	Get(id int) *gorms.PersonasGorm
	Save(e *gorms.PersonasGorm) int
	Delete(id int) bool
	Query(filter string) []gorms.PersonasGorm
}

func (dao *DePrimeraDaos) GetPersonasDao() *PersonasDaoImpl {
	return &PersonasDaoImpl{}
}

type SancionesDao interface {
	GetAll() []gorms.SancionesGorm
	Get(id int) *gorms.SancionesGorm
	Save(e *gorms.SancionesGorm) int
	Delete(id int) bool
	Query(filter string) []gorms.SancionesGorm
}

func (dao *DePrimeraDaos) GetSancionesDao() *SancionesDaoImpl {
	return &SancionesDaoImpl{}
}

type SancionesEquiposDao interface {
	Save(e *gorms.SancionesEquiposGorm) int64
	Delete(id int64) bool
}

func (dao *DePrimeraDaos) GetSancionesEquiposDao() *SancionesEquiposDaoImpl {
	return &SancionesEquiposDaoImpl{}
}

type ZonasDao interface {
	GetAll() []gorms.ZonasGorm
	Get(id int) *gorms.ZonasGorm
	Save(e *gorms.ZonasGorm) int
	Delete(id int) bool
	Query(filter string) []gorms.ZonasGorm
}

func (dao *DePrimeraDaos) GetZonasDao() *ZonasDaoImpl {
	return &ZonasDaoImpl{}
}

type ZonasEquiposDao interface {
	Save(e *gorms.ZonasEquiposGorm) int64
	Delete(id int64) bool
}

func (dao *DePrimeraDaos) GetZonasEquiposDao() *ZonasEquiposDaoImpl {
	return &ZonasEquiposDaoImpl{}
}

type JugadoresDao interface {
	GetAll() []gorms.JugadoresGorm
	Get(id int) *gorms.JugadoresGorm
	Save(e *gorms.JugadoresGorm) int
	Delete(id int) bool
	Query(filter string) []gorms.JugadoresGorm
}

func (dao *DePrimeraDaos) GetJugadoresDao() *JugadoresDaoImpl {
	return &JugadoresDaoImpl{}
}

type NotificacionesDao interface {
	GetAll() []gorms.NotificacionesGorm
	Get(id int) *gorms.NotificacionesGorm
	Save(e *gorms.NotificacionesGorm) int
	Delete(id int) bool
	Query(filter string) []gorms.NotificacionesGorm
}

func (dao *DePrimeraDaos) GetNotificacionesDao() *NotificacionesDaoImpl {
	return &NotificacionesDaoImpl{}
}

type AppGruposDao interface {
	GetAll() []gorms.AppGruposGorm
}

func (dao *DePrimeraDaos) GetAppGruposDao() *AppGruposDaoImpl {
	return &AppGruposDaoImpl{}
}
