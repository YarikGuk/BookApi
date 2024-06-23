package transactions

import (
	"encoding/json"
	"github.com/YarikGuk/BookApi/internal/authors"
	"github.com/YarikGuk/BookApi/internal/books"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type UpdateRequest struct {
	BookTitle string `json:"book_title"`
	AuthorBio string `json:"author_bio"`
}

type AuthorBookHandler struct {
	BookRepo   *books.BookRepository
	AuthorRepo *authors.AuthorRepository
}

func (h *AuthorBookHandler) UpdateBookAndAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookID, err := strconv.Atoi(params["book_id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	authorID, err := strconv.Atoi(params["author_id"])
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	var updateReq UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tx, err := h.BookRepo.DB.Begin()
	if err != nil {
		log.Println("Error beginning transaction:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := tx.Commit(); err != nil {
			log.Println("Error committing transaction:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}()

	if err := h.BookRepo.UpdateTitleInTransaction(tx, bookID, updateReq.BookTitle); err != nil {
		log.Println("Error updating book title:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.AuthorRepo.UpdateBioInTransaction(tx, authorID, updateReq.AuthorBio); err != nil {
		log.Println("Error updating author bio:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
