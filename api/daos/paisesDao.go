package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
)

type PaisesDaoImpl struct{}

func (ed *PaisesDaoImpl) GetAll() []models.Paises {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from app_paises")
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}

	paises := []models.Paises{}
	for rows.Next() {
		pais := models.Paises{}
		error := rows.Scan(&pais.IDPais, &pais.Nombre)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		paises = append(paises, pais)
	}
	return paises
}

func (ed *PaisesDaoImpl) Get(id int) models.Paises {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select * from app_paises where id_pais = $1", id)
	pais := models.Paises{}
	error := row.Scan(&pais.IDPais, &pais.Nombre)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return pais
}
