package sql

import (
	"database/sql"
	"fmt"
	"log"
)

func dbSelect(db *sql.DB) {
	//Get all the tables from Database
	rows, err := db.Query("select 1")
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var tables string
		//rows.Scan(&tables)
		log.Println(tables)
	}
}

func dbSelectOne(db *sql.DB) {
	//Get all the tables from Database
	id := 6
	row := db.QueryRow("select nombre from ligas where id_liga = ?", id)
	var nombre string
	err := row.Scan(&nombre)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	}

	log.Println(nombre)
}

func dbInsert(db *sql.DB) {
	nombre := "leonel"
	domicilio := "colombia 396"
	telefono := "3413553810"

	idLiga := 2
	_, error := db.Exec("insert into ligas(id_liga, nombre, domicilio, telefono) values(?, ?,?,?)", idLiga, nombre, domicilio, telefono)
	if error != nil {
		panic(error)
	}
	fmt.Println("New record ID is:", idLiga)
}

func dbInsertRecord(db *sql.DB) {
	idLiga := 2
	nombre := "leonel"

	fmt.Println(idLiga)
	fmt.Println(nombre)

	/*
		stmt, error := db.Prepare("insert into equipos(id_liga, nombre) values(?,?)")
		res, error := stmt.Exec(idLiga, nombre)
		idEquipo, error := res.LastInsertId()

		if error != nil {
			panic(error)
		}
		fmt.Println("New record equipo ID is:", idEquipo)
	*/
}

func dbUpdate(db *sql.DB) {
	nombre := "leonel 2"
	cuit := "31631073"
	idLiga := 2
	_, error := db.Exec("update ligas set nombre = ?, cuit = ?  where id_liga = ?", nombre, cuit, idLiga)
	if error != nil {
		panic(error)
	}
	fmt.Println("New record ID is:", idLiga)
}

func dbDelete(db *sql.DB) {
	idEquipo := 3
	_, error := db.Exec("delete from equipos where id_equipo = ?", idEquipo)
	if error != nil {
		panic(error)
	}
	fmt.Println("New record ID is:", idEquipo)
}
