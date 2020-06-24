package router

import (
	"log"
	"net/http"

	"deprimera/src/services"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world 3"}`))
	log.Println("Paso por aca")
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)

	router.HandleFunc("/equipos", services.GetEquipos).Methods("GET")
	router.HandleFunc("/equipos", services.SaveEquipos).Methods("POST")
	router.HandleFunc("/equipos/info", services.Info).Methods("GET")
	return router
}
