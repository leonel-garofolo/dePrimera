package models

import (
	"deprimera/api/application"
	"log"
)

type Equipos struct {
	Foto       []byte `gorm:"column:foto"`
	Habilitado bool   `gorm:"column:habilitado"`
	IDEquipo   int    `gorm:"column:id_equipo;primary_key"`
	IDLiga     int    `gorm:"column:id_liga"`
	Nombre     string `gorm:"column:nombre"`
}

// TableName sets the insert table name for this struct type
func (e *Equipos) TableName() string {
	return "equipos"
}

func (e *Equipos) SaveEquipo() int {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	equipoDB := db.Find(&e)
	if equipoDB == nil {
		db.Create(&e).Last(&e)
	} else {
		db.Save(&e)
	}
	return e.IDEquipo
}

func (e *Equipos) GetEquipo() {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	db.Find(&e, e.IDEquipo)
}

func GetAllEquipo() []Equipos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	equipos := []Equipos{}
	db.Find(&equipos)
	return equipos
}
