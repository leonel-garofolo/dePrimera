package sql

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func createConn() sql.DB {
	user := "root"
	pass := "root"
	database := "de_primera_app"

	db, err := sql.Open("mysql", user+":"+pass+"@/"+database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln("Failed to connect database")
	}

	return db
}

func TestDBSelect(t *testing.T) {
	db := createConn()
	defer db.Close()

	dbSelect(db)
}

func TestDBSelectOne(t *testing.T) {
	db := createConn()
	defer db.Close()

	dbSelectOne(db)
}

func TestDBInsert(t *testing.T) {
	db := createConn()
	defer db.Close()

	dbInsert(db)
}

func TestDBInsertRecord(t *testing.T) {
	db := createConn()
	defer db.Close()

	dbInsertRecord(db)
}

func TestDBUpdate(t *testing.T) {
	db := createConn()
	defer db.Close()

	dbUpdate(db)
}

func TestDBDelete(t *testing.T) {
	db := createConn()
	defer db.Close()

	dbDelete(db)
}
