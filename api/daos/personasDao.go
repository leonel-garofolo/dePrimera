package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"fmt"
	"log"
)

type PersonasDaoImpl struct{}

func (ed *PersonasDaoImpl) GetAll() []models.Personas {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	personas := []models.Personas{}
	db.Find(&personas)
	return personas
}

func (ed *PersonasDaoImpl) Get(id int) models.Personas {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	persona := models.Personas{}
	db.Find(&persona, id)
	return persona
}

func (ed *PersonasDaoImpl) Save(e *models.Personas) int {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	personaDB := db.Find(&e)
	if personaDB == nil {
		db.Create(&e).Last(&e)
	} else {
		db.Save(&e)
	}
	return e.IDPersona
}

func (ed *PersonasDaoImpl) Delete(id int) bool {
	persona := models.Personas{}

	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}
	db.Where("id_persona = ?", id).First(&persona)
	if persona.IDPersona > 0 {
		db.Where("id_persona=?", id).Delete(&models.Personas{})
		fmt.Println("delete ID is:", id)
		return true
	} else {
		fmt.Println("no exist ID:", id)
		return false
	}
}

func (ed *PersonasDaoImpl) Query(sql string) []models.Personas {
	personas := []models.Personas{}
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	db.First(personas, 1)
	return personas
}
