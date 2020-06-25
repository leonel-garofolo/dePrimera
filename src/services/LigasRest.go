package services

import (
	"deprimera/src/application"
	"deprimera/src/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetLigas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	ligas := models.Ligas{}

	j, err := json.Marshal(ligas)
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(j)
		log.Println(string(j))
	}
}

func SaveLiga(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	ligas := &models.Ligas{}
	json.NewDecoder(r.Body).Decode(ligas)

	db, err := application.GetDB()
	defer db.Close()

	if err != nil {
		log.Println(err.Error())
	}

	ligaDB := db.Find(ligas.IDLiga)
	if ligaDB != nil {
		db.Create(ligas)
	} else {
		db.Update(ligas)
	}

	log.Println(ligas)
	w.Write([]byte("insertado"))
}

func InfoLigas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	ligas := &models.Ligas{}
	json.NewDecoder(r.Body).Decode(ligas)

	j, err := json.Marshal(ligas)
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(j)
		log.Println(string(j))
	}
}
