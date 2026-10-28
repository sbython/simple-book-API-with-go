package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"main/datafunc"
	"net/http"
)

func main() {
	db := datafunc.Connect()
	defer db.Close()
	fmt.Println("Connected to the database successfully!")
	var app = datafunc.Book{DB: db}
	router := mux.NewRouter()

	router.HandleFunc("/books", app.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", app.GetBookById).Methods("GET")
	router.HandleFunc("/books", app.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", app.DeleteBook).Methods("DELETE")
	router.HandleFunc("/booksSize", app.GetBookSize).Methods("GET")

	log.Fatal(http.ListenAndServe(":1337", router))
}
