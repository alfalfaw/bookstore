package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct (Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// return a empty book when not found
	json.NewEncoder(w).Encode(&Book{})
}

// Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	// 生成id
	book.ID = strconv.Itoa(rand.Intn(100000000))
	books = append(books, book)
	// return the book created
	json.NewEncoder(w).Encode(book)

}

// Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {

			books = append(books[:index], books[index+1:]...)

			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)

			book.ID = params["id"]
			books = append(books, book)
			// return the book created
			json.NewEncoder(w).Encode(book)

			return
		}

	}
	json.NewEncoder(w).Encode(books)
}

// Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(books)

}

func main() {
	r := mux.NewRouter()
	books = append(books, Book{ID: "1", Isbn: "13824", Title: "Book1", Author: &Author{Firstname: "abc", Lastname: "edf"}})
	books = append(books, Book{ID: "2", Isbn: "13824", Title: "Book2", Author: &Author{Firstname: "abc", Lastname: "edf"}})
	books = append(books, Book{ID: "3", Isbn: "13824", Title: "Book3", Author: &Author{Firstname: "abc", Lastname: "edf"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	// 打印严重错误
	log.Fatal(http.ListenAndServe(":8000", r))

}
