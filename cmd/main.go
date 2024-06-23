package main

import (
	"github.com/YarikGuk/BookApi/internal/authors"
	"github.com/YarikGuk/BookApi/internal/books"
	"github.com/YarikGuk/BookApi/internal/database"
	"github.com/YarikGuk/BookApi/internal/transactions"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bookRepo := &books.BookRepository{DB: db}
	bookHandler := &books.BookHandler{Repo: bookRepo}

	authorRepo := &authors.AuthorRepository{DB: db}
	authorHandler := &authors.AuthorHandler{Repo: authorRepo}

	authBookHandler := &transactions.AuthorBookHandler{AuthorRepo: authorRepo, BookRepo: bookRepo}

	router := mux.NewRouter()
	router.HandleFunc("/books", bookHandler.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books", bookHandler.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", bookHandler.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", bookHandler.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", bookHandler.DeleteBook).Methods(http.MethodDelete)

	router.HandleFunc("/books/{book_id}/authors/{author_id}", authBookHandler.UpdateBookAndAuthor).Methods(http.MethodPut)

	router.HandleFunc("/authors", authorHandler.CreateAuthor).Methods(http.MethodPost)
	router.HandleFunc("/authors", authorHandler.GetAuthors).Methods(http.MethodGet)
	router.HandleFunc("/authors/{id}", authorHandler.GetAuthor).Methods(http.MethodGet)
	router.HandleFunc("/authors/{id}", authorHandler.UpdateAuthor).Methods(http.MethodPut)
	router.HandleFunc("/authors/{id}", authorHandler.DeleteAuthor).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", router))
}
