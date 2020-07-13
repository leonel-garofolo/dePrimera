package daos

import (
	"database/sql"
	"deprimera/api/application"
	"deprimera/api/models"
	"log"
)

type AppGruposDaoImpl struct{}

func (ed *AppGruposDaoImpl) GetAll() []models.AppGrupos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from app_grupos")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	grupos := []models.AppGrupos{}
	for rows.Next() {
		grupo := models.AppGrupos{}
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
