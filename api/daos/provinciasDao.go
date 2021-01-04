package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
)

type ProvinciasDaoImpl struct{}

func (ed *ProvinciasDaoImpl) GetAll() []gorms.ProvinciasGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from app_provincias")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	var provincias []gorms.ProvinciasGorm
	for rows.Next() {
		provincia := gorms.ProvinciasGorm{}
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

func (ed *ProvinciasDaoImpl) Get(idPais int, idProvincia int) gorms.ProvinciasGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select * from app_provincias where id_pais = ? and id_provincia = ?", idPais, idProvincia)
	pais := gorms.ProvinciasGorm{}
	error := row.Scan(&pais.IDPais, &pais.Nombre)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return pais
}
