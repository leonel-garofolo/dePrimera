package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
)

type PaisesDaoImpl struct{}

func (ed *PaisesDaoImpl) GetAll() []gorms.PaisesGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from app_paises")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	var paises []gorms.PaisesGorm
	for rows.Next() {
		pais := gorms.PaisesGorm{}
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

func (ed *PaisesDaoImpl) Get(id int) gorms.PaisesGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select * from app_paises where id_pais = ?", id)
	pais := gorms.PaisesGorm{}
	error := row.Scan(&pais.IDPais, &pais.Nombre)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return pais
}
