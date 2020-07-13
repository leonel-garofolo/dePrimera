package daos

import (
	"database/sql"
	"deprimera/api/application"
	"deprimera/api/models"
	"log"
)

type NotificacionesDaoImpl struct{}

func (ed *NotificacionesDaoImpl) GetAll() []models.Notificaciones {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from notificaciones")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	notificaciones := []models.Notificaciones{}
	for rows.Next() {
		notificacion := models.Notificaciones{}
		error := rows.Scan(&notificacion.IDNotificacion, &notificacion.Titulo, &notificacion.Texto)
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

func (ed *NotificacionesDaoImpl) Save(e *models.Notificaciones) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDNotificacion > 0 {
		_, error := db.Exec("update notificaciones"+
			" set titulo=?, texto=?, id_grupo=?"+
			" where id_notificacion = ?", e.Titulo, e.Texto, e.IDGrupo, e.IDNotificacion)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into notificaciones"+
			" (id_notificacion, titulo, texto, id_grupo) "+
			" values(?,?,?,?,?)", e.IDNotificacion, e.Titulo, e.Texto, e.IDGrupo)
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

	_, error := db.Exec("delete from notificaciones where id_notificacion = ?", IDNotificacion)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return true
}
