package daos

import (
	"database/sql"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"	
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	"log"
)

// AppGruposDaoImpl struct
type AppGruposDaoImpl struct{}

// GetAll appGrupos
func (ed *AppGruposDaoImpl) GetAll() []gorms.AppGruposGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from app_grupos")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	grupos := []gorms.AppGruposGorm{}
	for rows.Next() {
		grupo := gorms.AppGruposGorm{}
		error := rows.Scan(&grupo.Idgrupo, &grupo.Descripcion)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		grupos = append(grupos, grupo)
	}
	return grupos
}
