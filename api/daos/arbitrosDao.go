package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"fmt"
	"log"
)

type ArbitrosDaoImpl struct{}

func (ed *ArbitrosDaoImpl) GetAll() []models.Arbitros {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	arbitros := []models.Arbitros{}
	db.Find(&arbitros)
	return arbitros
}

func (ed *ArbitrosDaoImpl) Get(id int) models.Arbitros {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	arbitro := models.Arbitros{}
	db.Find(&arbitro, id)
	return arbitro
}

func (ed *ArbitrosDaoImpl) Save(e models.Arbitros) int {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	arbitroDB := db.Find(&e)
	if arbitroDB == nil {
		db.Create(&e).Last(&e)
	} else {
		db.Save(&e)
	}
	return e.IDArbitro
}

func (ed *ArbitrosDaoImpl) Delete(id int) bool {
	arbitro := models.Arbitros{}

	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}
	db.Where("id_arbitro = ?", id).First(&arbitro)
	if arbitro.IDArbitro > 0 {
		db.Where("id_arbitro=?", id).Delete(&models.Arbitros{})
		fmt.Println("delete ID is:", id)
		return true
	} else {
		fmt.Println("no exist ID:", id)
		return false
	}
}

func (ed *ArbitrosDaoImpl) Query(sql string) []models.Arbitros {
	arbitros := []models.Arbitros{}
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	db.First(arbitros, 1)
	return arbitros
}
