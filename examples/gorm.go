package main

import (
	"database/sql"
	"deprimera/api/application"
	"deprimera/api/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := application.GetDB()
	if err != nil {
		log.Fatalln("fail to database connection")
	}
	defer db.Close()

	//dbSelect(db)
	//dbSelectOne(db)
	//dbInsert(db)
	//dbInsertRecord(db)
	dbUpdate(db)
	//dbDelete(db)

}

func dbSelect(db *gorm.DB) {
	//Get all the tables from Database
	ligas := &models.Ligas{
		IDLiga: 1,
	}
	err := db.First(ligas)
	if err != nil {
		panic(err)
	}
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
		IDLiga: 2,
		Nombre: "test",
	}

	db.Create(&equipo)
}

func dbInsertRecord(db *sql.DB) {
	idLiga := 2
	nombre := "leonel"

	stmt, error := db.Prepare("insert into equipos(id_liga, nombre) values(?,?)")
	res, error := stmt.Exec(idLiga, nombre)
	idEquipo, error := res.LastInsertId()

	if error != nil {
		panic(error)
	}
	fmt.Println("New record equipo ID is:", idEquipo)
}

func dbUpdate(db *gorm.DB) {
	equipo := models.Equipos{
		IDEquipo: 5,
		IDLiga:   2,
		Nombre:   "test 2",
	}

	db.Save(&equipo)
}

func dbDelete(db *sql.DB) {
	idEquipo := 3
	_, error := db.Exec("delete from equipos where id_equipo = ?", idEquipo)
	if error != nil {
		panic(error)
	}
	fmt.Println("New record ID is:", idEquipo)
}
