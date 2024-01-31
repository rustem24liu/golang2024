package main

import (
	"github.com/gorilla/mux"
	"github.com/rustem24liu/golang2024/TSIS1/internal/handlers"
	"log"
	"net/http"
)

func main() {

	log.Println("Starting API server")
	r := mux.NewRouter()

	r.HandleFunc("/home/bands", handlers.GetBands).Methods("GET")
	r.HandleFunc("/home/bands/{id}", handlers.GetBand).Methods("GET")
	r.HandleFunc("/home/bands", handlers.CreateBand).Methods("POST")
	r.HandleFunc("/home/bands", handlers.UpdateBand).Methods("PUT")
	r.HandleFunc("/home/bands/{id}", handlers.DeleteBand).Methods("DELETE")

	log.Println("Server in :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
