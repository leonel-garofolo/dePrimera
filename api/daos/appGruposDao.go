package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
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

// GetAll appGrupos
func (ed *AppGruposDaoImpl) GetUserAppGrupos(idUser string) gorms.AppGruposGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select p.idgrupo, ag.descripcion "+
		" from personas p "+
		" inner join app_grupos ag on p.idgrupo = ag.idgrupo "+
		" where p.id_user = ?", idUser)
	if err != nil {
		log.Fatalln("Failed to query")
	}

	grupo := gorms.AppGruposGorm{}
	for rows.Next() {
		error := rows.Scan(&grupo.Idgrupo, &grupo.Descripcion)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
	}
	return grupo
}
