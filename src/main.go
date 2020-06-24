package main

import (
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world 2"}`))
	log.Println("Paso por aca")
}

func main() {
	http.Handle("/", &server{})

	http.Handle("/equipos", equipos.Get)
	log.Fatal(http.ListenAndServe(":8081", nil))
	log.Println("Server started on: http://localhost:8081")
}
