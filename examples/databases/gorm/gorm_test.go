package gorm

import (
	"log"
	"testing"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
)

func TestMensajeAltaPrioridad(t *testing.T) {
	db, err := application.GetDB()
	if err != nil {
		log.Fatalln("fail to database connection")
	}
	defer db.Close()

	//dbSelect(db)
	//dbSelectOne(db)
	//dbInsert(db)
	//dbInsertRecord(db)
	//dbUpdate(db)
	//dbDelete(db)

}
