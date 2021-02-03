package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
)

type ProvinciasDaoImpl struct{}

func (ed *ProvinciasDaoImpl) GetAll() []models.Provincias {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from app_provincias")
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}

	var provincias []models.Provincias
	for rows.Next() {
		provincia := models.Provincias{}
		error := rows.Scan(&provincia.IDProvincia, &provincia.Nombre, &provincia.IDPais)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		provincias = append(provincias, provincia)
	}
	return provincias
}

func (ed *ProvinciasDaoImpl) Get(idPais int, idProvincia int) models.Provincias {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select * from app_provincias where id_pais = $1 and id_provincia = $2", idPais, idProvincia)
	pais := models.Provincias{}
	error := row.Scan(&pais.IDPais, &pais.Nombre)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return pais
}
