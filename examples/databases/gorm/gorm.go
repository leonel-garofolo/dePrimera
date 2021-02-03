package gorm

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
)

func dbSelect(db *gorm.DB) {
	//Get all the tables from Database
	idLiga := 2
	ligas := &models.Ligas{}
	db.First(ligas, idLiga)
	log.Println(ligas)
}

func dbSelectOne(db *gorm.DB) {
	//Get all the tables from Database
	id := 2
	rows, err := db.Raw("select nombre from ligas where id_liga = ?", id).Rows()
	defer rows.Close()
	var nombre string
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	}

	for rows.Next() {
		rows.Scan(&nombre)
	}

	log.Println(nombre)
}

func dbInsert(db *gorm.DB) { //verificar como atrapar el error del insert
	equipo := models.Equipos{
		Nombre: "test",
	}

	db.Create(&equipo)
}

func dbInsertRecord(db *gorm.DB) {
	equipo := models.Equipos{
		Nombre:     "test 2 ",
		Habilitado: true,
	}

	db.Create(&equipo).Last(&equipo)
	fmt.Println("New record equipo ID is:", equipo.IDEquipo)
}

func dbUpdate(db *gorm.DB) {
	equipo := models.Equipos{
		IDEquipo: 5,
		Nombre:   "test 2",
	}

	db.Save(&equipo)
}

func dbDelete(db *gorm.DB) {
	idEquipo := 2
	equipo := models.Equipos{}
	db.Where("id_equipo = ?", idEquipo).First(&equipo)
	/*
		if equipo.IDLiga > 0 {
			db.Where("id_equipo=?", idEquipo).Delete(&models.Equipos{})
			fmt.Println("delete ID is:", idEquipo)
		} else {
			fmt.Println("no exist ID:", idEquipo)
		}
	*/

}
