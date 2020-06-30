package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"fmt"
	"log"
)

type PartidosDaoImpl struct{}

func (ed *PartidosDaoImpl) GetAll() []models.Partidos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	partidos := []models.Partidos{}
	db.Find(&partidos)
	return partidos
}

func (ed *PartidosDaoImpl) Get(id int) models.Partidos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	partido := models.Partidos{}
	db.Find(&partido, id)
	return partido
}

func (ed *PartidosDaoImpl) Save(e *models.Partidos) int {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	partidoDB := db.Find(&e)
	if partidoDB == nil {
		db.Create(&e).Last(&e)
	} else {
		db.Save(&e)
	}
	return e.IDPartidos
}

func (ed *PartidosDaoImpl) Delete(id int) bool {
	partido := models.Partidos{}

	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}
	db.Where("id_partido = ?", id).First(&partido)
	if partido.IDPartidos > 0 {
		db.Where("id_partido=?", id).Delete(&models.Partidos{})
		fmt.Println("delete ID is:", id)
		return true
	} else {
		fmt.Println("no exist ID:", id)
		return false
	}
}

func (ed *PartidosDaoImpl) Query(sql string) []models.Partidos {
	partidos := []models.Partidos{}
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	db.First(partidos, 1)
	return partidos
}
