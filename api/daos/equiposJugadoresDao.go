package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"fmt"
	"log"
)

type EquiposJugadoresDaoImpl struct{}

func (ed *EquiposJugadoresDaoImpl) GetAll() []models.EquiposJugadores {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	equiposJugadores := []models.EquiposJugadores{}
	db.Find(&equiposJugadores)
	return equiposJugadores
}

func (ed *EquiposJugadoresDaoImpl) Get(id int) models.EquiposJugadores {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	equipoJugadores := models.EquiposJugadores{}
	db.Find(&equipoJugadores, id)
	return equipoJugadores
}

func (ed *EquiposJugadoresDaoImpl) Save(e models.EquiposJugadores) int {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	equipoJugadorDB := db.Find(&e)
	if equipoJugadorDB == nil {
		db.Create(&e).Last(&e)
	} else {
		db.Save(&e)
	}
	return e.IDEquipos
}

func (ed *EquiposJugadoresDaoImpl) Delete(id int) bool {
	equipoJugadores := models.EquiposJugadores{}

	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}
	db.Where("id_equipo = ?", id).First(&equipoJugadores)
	if equipoJugadores.IDEquipos > 0 && equipoJugadores.IDJugadores > 0 {
		db.Where("id_equipo=?", id).Delete(&models.EquiposJugadores{})
		fmt.Println("delete ID is:", id)
		return true
	} else {
		fmt.Println("no exist ID:", id)
		return false
	}
}

func (ed *EquiposJugadoresDaoImpl) Query(sql string) []models.EquiposJugadores {
	equiposJugadores := []models.EquiposJugadores{}
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	db.First(equiposJugadores, 1)
	return equiposJugadores
}
