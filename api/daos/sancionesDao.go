package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"fmt"
	"log"
)

type SancionesDaoImpl struct{}

func (ed *SancionesDaoImpl) GetAll() []models.Sanciones {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	sanciones := []models.Sanciones{}
	db.Find(&sanciones)
	return sanciones
}

func (ed *SancionesDaoImpl) Get(id int) models.Sanciones {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	sancion := models.Sanciones{}
	db.Find(&sancion, id)
	return sancion
}

func (ed *SancionesDaoImpl) Save(e *models.Sanciones) int {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	sancionDB := db.Find(&e)
	if sancionDB == nil {
		db.Create(&e).Last(&e)
	} else {
		db.Save(&e)
	}
	return e.IDSanciones
}

func (ed *SancionesDaoImpl) Delete(id int) bool {
	sancion := models.Sanciones{}

	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}
	db.Where("id_sancion = ?", id).First(&sancion)
	if sancion.IDSanciones > 0 {
		db.Where("id_sancion=?", id).Delete(&models.Sanciones{})
		fmt.Println("delete ID is:", id)
		return true
	} else {
		fmt.Println("no exist ID:", id)
		return false
	}
}

func (ed *SancionesDaoImpl) Query(sql string) []models.Sanciones {
	sancions := []models.Sanciones{}
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	db.First(sancions, 1)
	return sancions
}
