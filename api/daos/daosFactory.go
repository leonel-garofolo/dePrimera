package daos

type EntityDao interface {
	GetAll()
	Get(id int)
	Save() int
	Delete(id int) bool
	Query(filter string)
}

type DaosFactory interface {
	GetEquiposDao() EntityDao
}

type DAOs struct{}

func GetEquiposDao() *EquiposDao {
	return &EquiposDao{}
}
