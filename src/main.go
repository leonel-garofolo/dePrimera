package main

import (
	"log"
	"net/http"

	"deprimera/src/router"
)

func main() {
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8081", r))
	log.Println("Server started on: http://localhost:8081")
}
