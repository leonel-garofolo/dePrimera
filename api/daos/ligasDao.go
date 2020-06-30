package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"fmt"
	"log"
)

type LigasDaoImpl struct{}

func (ed *LigasDaoImpl) GetAll() []models.Ligas {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	ligas := []models.Ligas{}
	db.Find(&ligas)
	return ligas
}

func (ed *LigasDaoImpl) Get(id int) models.Ligas {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	liga := models.Ligas{}
	db.Find(&liga, id)
	return liga
}

func (ed *LigasDaoImpl) Save(e *models.Ligas) int {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	ligaDB := db.Find(&e)
	if ligaDB == nil {
		db.Create(&e).Last(&e)
	} else {
		db.Save(&e)
	}
	return e.IDLiga
}

func (ed *LigasDaoImpl) Delete(id int) bool {
	liga := models.Ligas{}

	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}
	db.Where("id_liga = ?", id).First(&liga)
	if liga.IDLiga > 0 {
		db.Where("id_liga=?", id).Delete(&models.Ligas{})
		fmt.Println("delete ID is:", id)
		return true
	} else {
		fmt.Println("no exist ID:", id)
		return false
	}
}

func (ed *LigasDaoImpl) Query(sql string) []models.Ligas {
	ligas := []models.Ligas{}
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	db.First(ligas, 1)
	return ligas
}
