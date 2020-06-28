package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Shelnutt2/db2struct"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Program to reverse engineer your mysql database into gorm models
func main() {
	user := "root"
	pass := "root"
	host := "localhost"
	database := "de_primera_app"
	port := 3306
	packagename := "models"
	fmt.Println("Connecting to mysql server " + host + ":" + strconv.Itoa(port))
	db, err := gorm.Open("mysql", user+":"+pass+"@/"+database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln("Failed to connect database")
	}
	defer db.Close()
	//Get all the tables from Database
	rows, err := db.Raw("SHOW TABLES").Rows()
	defer rows.Close()
	for rows.Next() {
		var table string
		rows.Scan(&table)
		columnDataTypes, err := db2struct.GetColumnsFromMysqlTable(user, pass, host, port, database, table)
		if err != nil {
			fmt.Println("Error in selecting column data information from mysql information schema")
			return
		}
		// Generate struct string based on columnDataTypes
		struc, err := db2struct.Generate(*columnDataTypes, table, table, packagename, true, true, false)
		if err != nil {
			fmt.Println("Error in creating struct from json: " + err.Error())
			return
		}
		file, err := os.Create(packagename + "/" + table + ".go")
		if err != nil {
			log.Printf("exist: " + table)
			log.Fatal("exception", err)
			continue
		}
		defer file.Close()
		fmt.Fprintf(file, string(struc))
		log.Println("Wrote " + table + ".go to disk")
	}

}
