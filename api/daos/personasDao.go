package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
)

type PersonasDaoImpl struct{}

func (ed *PersonasDaoImpl) GetAll() []gorms.PersonasGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select id_persona, nombre, apellido, domicilio, edad, localidad, id_pais, id_provincia, id_tipo_doc, nro_doc from personas")
	if err != nil {
		//log.Fatalln("Failed to query")
		log.Println(err)
		panic(err)
	}

	apellido := sql.NullString{}
	personas := []gorms.PersonasGorm{}
	for rows.Next() {
		persona := gorms.PersonasGorm{}
		error := rows.Scan(&persona.IDPersona, &persona.Nombre, &apellido, &persona.Domicilio, &persona.Edad, &persona.Localidad, &persona.IDPais, &persona.IDProvincia, &persona.IDTipoDoc, &persona.NroDoc)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		persona.Apellido = apellido.String
		personas = append(personas, persona)
	}
	return personas
}

func (ed *PersonasDaoImpl) Get(id int) gorms.PersonasGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select id_persona, nombre, apellido, domicilio, edad, localidad, id_pais, id_provincia, id_tipo_doc, nro_doc from personas where id_persona = $1", id)
	persona := gorms.PersonasGorm{}
	error := row.Scan(&persona.IDPersona, &persona.Nombre, &persona.Apellido, &persona.Domicilio, &persona.Edad, &persona.Localidad, &persona.IDPais, &persona.IDProvincia, &persona.IDTipoDoc, &persona.NroDoc)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return persona
}

func (ed *PersonasDaoImpl) GetPersonasFromUser(idUser string) gorms.PersonasGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select id_persona, nombre, apellido, domicilio, edad, localidad, id_pais, id_provincia, id_tipo_doc, nro_doc from personas where id_user = $1", idUser)
	persona := gorms.PersonasGorm{}
	error := row.Scan(&persona.IDPersona, &persona.Nombre, &persona.Apellido, &persona.Domicilio, &persona.Edad, &persona.Localidad, &persona.IDPais, &persona.IDProvincia, &persona.IDTipoDoc, &persona.NroDoc)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return persona
}

func (ed *PersonasDaoImpl) Save(e *gorms.PersonasGorm) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDPersona > 0 {
		_, error := db.Exec("update personas"+
			" set nombre=$1, apellido=$2, domicilio=$3, edad=$4, localidad=$5, id_pais=$6, id_provincia=$7, id_tipo_doc=$8, nro_doc=$9 "+
			" where id_persona = $10", e.Nombre, e.Apellido, e.Domicilio, e.Edad, e.Localidad, e.IDPais, e.IDProvincia, e.IDTipoDoc, e.NroDoc, e.IDPersona)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into personas"+
			" (nombre, apellido, domicilio, edad, localidad, id_pais, id_provincia, id_tipo_doc, nro_doc) "+
			" values($1,$2,$3,$4,$5,$6,$7,$8,$9)",
			e.Nombre,
			e.Apellido,
			e.Domicilio,
			e.Edad,
			e.Localidad,
			e.IDPais,
			e.IDProvincia,
			e.IDTipoDoc,
			e.NroDoc)

		if error != nil {
			log.Println(error)
			panic(error)
		} else {
			IDPersona, _ := res.LastInsertId()
			if error != nil {
				log.Println(error)
				panic(error)
			}
			e.IDPersona = IDPersona
		}
	}
	return e.IDPersona
}

func (ed *PersonasDaoImpl) Delete(id int) (bool, error) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from personas where id_persona = $1", id)
	if error != nil {
		log.Println(error)
		return false, error
	}
	return true, nil
}
