package sql

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func createConn() *sql.DB {

	//db, err := sql.Open("postgres", user+":"+pass+"@/"+database+"?charset=utf8&parseTime=True&loc=Local")
	db, err := sql.Open("postgres", "postgres://kfaneklolueulx:838051822695796f4ba0a1764a70c02cca4a71abeda6a08af177389f260966c3@ec2-54-221-220-82.compute-1.amazonaws.com:5432/d6afam55njg5u6")
	if err != nil {
		log.Fatalln(err)
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
