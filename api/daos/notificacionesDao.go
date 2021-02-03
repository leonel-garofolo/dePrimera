package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
)

type NotificacionesDaoImpl struct{}

func (ed *NotificacionesDaoImpl) GetAll() []gorms.NotificacionesGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from notificaciones")
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}

	notificaciones := []gorms.NotificacionesGorm{}
	for rows.Next() {
		notificacion := gorms.NotificacionesGorm{}
		error := rows.Scan(&notificacion.IDNotificacion, &notificacion.IDGrupo, &notificacion.Titulo, &notificacion.Texto)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		notificaciones = append(notificaciones, notificacion)
	}
	return notificaciones
}

func (ed *NotificacionesDaoImpl) Save(e *gorms.NotificacionesGorm) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDNotificacion > 0 {
		_, error := db.Exec("update notificaciones"+
			" set titulo=$1, texto=$2, id_grupo=$3"+
			" where id_notificacion = $4", e.Titulo, e.Texto, e.IDGrupo, e.IDNotificacion)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into notificaciones"+
			" (id_notificacion, titulo, texto, id_grupo) "+
			" values($1,$2,$3,$4)", e.IDNotificacion, e.Titulo, e.Texto, e.IDGrupo)
		if error != nil {
			log.Println(error)
			panic(error)
		}
		e.IDNotificacion, _ = res.LastInsertId()
	}
	return e.IDNotificacion
}

func (ed *NotificacionesDaoImpl) Delete(IDNotificacion int64) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from notificaciones where id_notificacion = $1", IDNotificacion)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return true
}
