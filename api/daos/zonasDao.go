package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"fmt"
	"log"
)

type ZonasDaoImpl struct{}

func (ed *ZonasDaoImpl) GetAll() []models.Zonas {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	zonas := []models.Zonas{}
	db.Find(&zonas)
	return zonas
}

func (ed *ZonasDaoImpl) Get(id int) models.Zonas {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	zona := models.Zonas{}
	db.Find(&zona, id)
	return zona
}

func (ed *ZonasDaoImpl) Save(e models.Zonas) int {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	zonaDB := db.Find(&e)
	if zonaDB == nil {
		db.Create(&e).Last(&e)
	} else {
		db.Save(&e)
	}
	return e.IDZona
}

func (ed *ZonasDaoImpl) Delete(id int) bool {
	sancion := models.Zonas{}

	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}
	db.Where("id_sancion = ?", id).First(&sancion)
	if sancion.IDZonas > 0 {
		db.Where("id_sancion=?", id).Delete(&models.Zonas{})
		fmt.Println("delete ID is:", id)
		return true
	} else {
		fmt.Println("no exist ID:", id)
		return false
	}
}

func (ed *ZonasDaoImpl) Query(sql string) []models.Zonas {
	zonas := []models.Zonas{}
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	db.First(zonas, 1)
	return zonas
}
