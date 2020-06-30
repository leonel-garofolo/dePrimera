package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"fmt"
	"log"
)

type SancionesEquiposDaoImpl struct{}

func (ed *SancionesEquiposDaoImpl) GetAll() []models.SancionesEquipos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	sancionEquiposes := []models.SancionesEquipos{}
	db.Find(&sancionEquiposes)
	return sancionEquiposes
}

func (ed *SancionesEquiposDaoImpl) Get(id int) models.SancionesEquipos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	sancionEquipos := models.SancionesEquipos{}
	db.Find(&sancionEquipos, id)
	return sancionEquipos
}

func (ed *SancionesEquiposDaoImpl) Save(e models.SancionesEquipos) int {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	sancionEquiposDB := db.Find(&e)
	if sancionEquiposDB == nil {
		db.Create(&e).Last(&e)
	} else {
		db.Save(&e)
	}
	return e.IDSancionesEquipos
}

func (ed *SancionesEquiposDaoImpl) Delete(id int) bool {
	sancionEquipos := models.SancionesEquipos{}

	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}
	db.Where("id_sancionEquipos = ?", id).First(&sancionEquipos)
	if sancionEquipos.IDSanciones > 0 && sancionEquipos.IDEquipo > 0 {
		db.Where("id_sancionEquipos=?", id).Delete(&models.SancionesEquipos{})
		fmt.Println("delete ID is:", id)
		return true
	} else {
		fmt.Println("no exist ID:", id)
		return false
	}
}

func (ed *SancionesEquiposDaoImpl) Query(sql string) []models.SancionesEquipos {
	sancionEquiposs := []models.SancionesEquipos{}
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	db.First(sancionEquiposs, 1)
	return sancionEquiposs
}
