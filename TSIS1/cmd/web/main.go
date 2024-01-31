package main

import (
	"github.com/gorilla/mux"
	"github.com/rustem24liu/golang2024/TSIS1/internal"
	"log"
)

func main() {

	log.Println("Starting API server")
	r := mux.NewRouter()

	r.HandleFunc("/home/bands", internal.getBands).Methods("GET")

}
