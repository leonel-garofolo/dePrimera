package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"fmt"
	"log"
)

type CampeonatosDaoImpl struct{}

func (ed *CampeonatosDaoImpl) GetAll() []models.Campeonatos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	campeonatos := []models.Campeonatos{}
	db.Find(&campeonatos)
	return campeonatos
}

func (ed *CampeonatosDaoImpl) Get(id int) models.Campeonatos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	campeonato := models.Campeonatos{}
	db.Find(&campeonato, id)
	return campeonato
}

func (ed *CampeonatosDaoImpl) Save(e *models.Campeonatos) int {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	campeonatoDB := db.Find(&e)
	if campeonatoDB == nil {
		db.Create(&e).Last(&e)
	} else {
		db.Save(&e)
	}
	return e.IDCampeonato
}

func (ed *CampeonatosDaoImpl) Delete(id int) bool {
	campeonato := models.Campeonatos{}

	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}
	db.Where("id_campeonato = ?", id).First(&campeonato)
	if campeonato.IDCampeonato > 0 {
		db.Where("id_campeonato=?", id).Delete(&models.Campeonatos{})
		fmt.Println("delete ID is:", id)
		return true
	} else {
		fmt.Println("no exist ID:", id)
		return false
	}
}

func (ed *CampeonatosDaoImpl) Query(sql string) []models.Campeonatos {
	campeonatos := []models.Campeonatos{}
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	db.First(campeonatos, 1)
	return campeonatos
}
