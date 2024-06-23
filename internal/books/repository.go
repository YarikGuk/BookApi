package books

import (
	"database/sql"
	appError "github.com/YarikGuk/BookApi/error"
)

type BookRepository struct {
	DB *sql.DB
}

func (r *BookRepository) Create(book *Book) error {
	_, err := r.DB.Exec(
		"INSERT INTO books (title, author_id, year, isbn) VALUES ($1, $2, $3, $4)",
		book.Title,
		book.AuthorID,
		book.Year,
		book.ISBN,
	)
	if err != nil {
		return appError.ErrInternalServer
	}
	return nil
}

func (r *BookRepository) GetAll() ([]Book, error) {
	rows, err := r.DB.Query("SELECT id, title, author_id, year, isbn FROM books")
	if err != nil {
		return nil, appError.ErrInternalServer
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.Year, &book.ISBN); err != nil {
			return nil, appError.ErrInternalServer
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *BookRepository) GetByID(id int) (*Book, error) {
	var book Book
	err := r.DB.QueryRow("SELECT id, title, author_id, year, isbn FROM books WHERE id = $1", id).Scan(&book.ID, &book.Title, &book.AuthorID, &book.Year, &book.ISBN)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, appError.ErrNotFound
		}
		return nil, appError.ErrInternalServer
	}
	return &book, nil
}

func (r *BookRepository) Update(book *Book) error {
	_, err := r.DB.Exec("UPDATE books SET title = $1, author_id = $2, year = $3, isbn = $4 WHERE id = $5", book.Title, book.AuthorID, book.Year, book.ISBN, book.ID)
	if err != nil {
		return appError.ErrInternalServer
	}
	return nil
}

func (r *BookRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return appError.ErrInternalServer
	}
	return nil
}

func (r *BookRepository) UpdateTitleInTransaction(tx *sql.Tx, bookID int, title string) error {
	_, err := tx.Exec("UPDATE books SET title = $1 WHERE id = $2", title, bookID)
	if err != nil {
		return appError.ErrInternalServer
	}
	return nil
}
