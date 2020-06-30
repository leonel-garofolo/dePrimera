package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"fmt"
	"log"
)

type AsistentesDaoImpl struct{}

func (ed *AsistentesDaoImpl) GetAll() []models.Asistentes {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	asistentes := []models.Asistentes{}
	db.Find(&asistentes)
	return asistentes
}

func (ed *AsistentesDaoImpl) Get(id int) models.Asistentes {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	asistente := models.Asistentes{}
	db.Find(&asistente, id)
	return asistente
}

func (ed *AsistentesDaoImpl) Save(e models.Asistentes) int {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	asistenteDB := db.Find(&e)
	if asistenteDB == nil {
		db.Create(&e).Last(&e)
	} else {
		db.Save(&e)
	}
	return e.IDAsistente
}

func (ed *AsistentesDaoImpl) Delete(id int) bool {
	asistente := models.Asistentes{}

	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}
	db.Where("id_asistente = ?", id).First(&asistente)
	if asistente.IDAsistente > 0 {
		db.Where("id_asistente=?", id).Delete(&models.Asistentes{})
		fmt.Println("delete ID is:", id)
		return true
	} else {
		fmt.Println("no exist ID:", id)
		return false
	}
}

func (ed *AsistentesDaoImpl) Query(sql string) []models.Asistentes {
	asistentes := []models.Asistentes{}
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	db.First(asistentes, 1)
	return asistentes
}
