package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"fmt"
	"log"
)

type ZonasEquiposDaoImpl struct{}

func (ed *ZonasEquiposDaoImpl) GetAll() []models.ZonasEquipos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	zonasEquipos := []models.ZonasEquipos{}
	db.Find(&zonasEquipos)
	return zonasEquipos
}

func (ed *ZonasEquiposDaoImpl) Get(id int) models.ZonasEquipos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	zonaEquipo := models.ZonasEquipos{}
	db.Find(&zonaEquipo, id)
	return zonaEquipo
}

func (ed *ZonasEquiposDaoImpl) Save(e models.ZonasEquipos) int {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	zonaEquipoDB := db.Find(&e)
	if zonaEquipoDB == nil {
		db.Create(&e).Last(&e)
	} else {
		db.Save(&e)
	}
	return e.IDZona
}

func (ed *ZonasEquiposDaoImpl) Delete(id int) bool {
	zonaEquipo := models.ZonasEquipos{}

	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}
	db.Where("id_sancion = ?", id).First(&zonaEquipo)
	if zonaEquipo.IDZona > 0 && zonaEquipo.IDEquipo > 0 {
		db.Where("id_sancion=?", id).Delete(&models.ZonasEquipos{})
		fmt.Println("delete ID is:", id)
		return true
	} else {
		fmt.Println("no exist ID:", id)
		return false
	}
}

func (ed *ZonasEquiposDaoImpl) Query(sql string) []models.ZonasEquipos {
	zonas := []models.ZonasEquipos{}
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	db.First(zonas, 1)
	return zonas
}
