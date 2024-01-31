package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Book Struct (Model)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID: "1", Isbn: "8449", Title: "Crime and Punishment", Author: &Author{Firstname: "Fedor", Lastname: "Dostaevsky"}})

	books = append(books, Book{ID: "2", Isbn: "9284", Title: "Player", Author: &Author{Firstname: "Fedor", Lastname: "Dostaevsky"}})

	books = append(books, Book{ID: "3", Isbn: "8924", Title: "Fathers and Sons", Author: &Author{Firstname: "Ivan", Lastname: "Turgenev"}})

	books = append(books, Book{ID: "4", Isbn: "8729", Title: "Cards on table", Author: &Author{Firstname: "Agatha", Lastname: "Christie"}})

	// Route Handler / Endpoints

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}

func updateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(request.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(1000000))
			books = append(books, book)
			json.NewEncoder(writer).Encode(book)
			return

		}
	}
	json.NewEncoder(writer).Encode(books)
}

func deleteBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(books)
}

func createBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(writer).Encode(book)
}

func getBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode(&Book{})
}

func getBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}
