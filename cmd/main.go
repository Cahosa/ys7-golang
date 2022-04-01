package main

import (
	"Ys7/internal/router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()

	log.Println("Server started on 0.0.0.0:8080")
	err := http.ListenAndServe("0.0.0.0:8080", r)

	if err != nil {
		log.Println("An error occured starting HTTP listener at 8080")
		log.Println("Error: " + err.Error())
	}
}
