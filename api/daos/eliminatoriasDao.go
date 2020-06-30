package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"fmt"
	"log"
)

type EliminatoriasDaoImpl struct{}

func (ed *EliminatoriasDaoImpl) GetAll() []models.Eliminatorias {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	eliminatorias := []models.Eliminatorias{}
	db.Find(&eliminatorias)
	return eliminatorias
}

func (ed *EliminatoriasDaoImpl) Get(id int) models.Eliminatorias {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	eliminatoria := models.Eliminatorias{}
	db.Find(&eliminatoria, id)
	return eliminatoria
}

func (ed *EliminatoriasDaoImpl) Save(e *models.Eliminatorias) int {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	eliminatoriaDB := db.Find(&e)
	if eliminatoriaDB == nil {
		db.Create(&e).Last(&e)
	} else {
		db.Save(&e)
	}
	return e.IDEliminatoria
}

func (ed *EliminatoriasDaoImpl) Delete(id int) bool {
	eliminatoria := models.Eliminatorias{}

	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}
	db.Where("id_eliminatoria = ?", id).First(&eliminatoria)
	if eliminatoria.IDEliminatoria > 0 {
		db.Where("id_eliminatoria=?", id).Delete(&models.Eliminatorias{})
		fmt.Println("delete ID is:", id)
		return true
	} else {
		fmt.Println("no exist ID:", id)
		return false
	}
}

func (ed *EliminatoriasDaoImpl) Query(sql string) []models.Eliminatorias {
	eliminatorias := []models.Eliminatorias{}
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	db.First(eliminatorias, 1)
	return eliminatorias
}
