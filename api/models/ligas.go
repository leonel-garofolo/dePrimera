package models

import (
	"database/sql"
	"deprimera/api/application"
	"log"
)

type Ligas struct {
	Cuit              sql.NullString `gorm:"column:cuit"`
	Domicilio         string         `gorm:"column:domicilio"`
	IDLiga            int            `gorm:"column:id_liga;primary_key"`
	MailContacto      sql.NullString `gorm:"column:mail_contacto"`
	Nombre            string         `gorm:"column:nombre"`
	NombreContacto    sql.NullString `gorm:"column:nombre_contacto"`
	Telefono          sql.NullString `gorm:"column:telefono"`
	TelefonoConctacto sql.NullString `gorm:"column:telefono_conctacto"`
}

// TableName sets the insert table name for this struct type
func (l *Ligas) TableName() string {
	return "ligas"
}

func (l *Ligas) SaveLigas() int {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	ligaDB := db.Find(&l)
	if ligaDB == nil {
		db.Create(&l).Last(&l)
	} else {
		db.Save(&l)
	}
	return l.IDLiga
}

func (l *Ligas) GetLiga() {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	db.Find(&l, l.IDLiga)
}

func GetAllLigas() []Ligas {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	ligas := []Ligas{}
	db.Find(&ligas)
	return ligas
}
