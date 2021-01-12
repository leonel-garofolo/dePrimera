package daos

import (
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
)

type ComentariosDaoImpl struct{}

func (ed *ComentariosDaoImpl) Save(e *models.Comentarios) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDComentario > 0 {
		_, error := db.Exec("update comentarios"+
			" set mail=?, puntaje=?, comentario=?"+
			" where id_comentario = ?", e.Mail, e.Puntaje, e.Comentario, e.IDComentario)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into comentarios"+
			" (mail, puntaje, comentario) "+
			" values(?,?,?)", e.Mail, e.Puntaje, e.Comentario)
		if error != nil {
			log.Println(error)
			panic(error)
		}
		e.IDComentario, _ = res.LastInsertId()
	}
	return int64(e.IDComentario)
}
