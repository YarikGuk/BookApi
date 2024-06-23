package authors

import (
	"database/sql"
	"errors"
	appErrors "github.com/YarikGuk/BookApi/error"
)

type AuthorRepository struct {
	DB *sql.DB
}

func (r *AuthorRepository) Create(author *Author) error {
	_, err := r.DB.Exec(
		"INSERT INTO authors (first_name, last_name, biography, birth_date) VALUES ($1, $2, $3, $4)",
		author.FirstName,
		author.LastName,
		author.Biography,
		author.BirthDate,
	)
	if err != nil {
		return appErrors.ErrInternalServer
	}
	return nil
}

func (r *AuthorRepository) GetAll() ([]Author, error) {
	rows, err := r.DB.Query("SELECT id, first_name, last_name, biography, birth_date FROM authors")
	if err != nil {
		return nil, appErrors.ErrInternalServer
	}
	defer rows.Close()

	var authors []Author
	for rows.Next() {
		var author Author
		if err := rows.Scan(&author.ID, &author.FirstName, &author.LastName, &author.Biography, &author.BirthDate); err != nil {
			return nil, appErrors.ErrInternalServer
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (r *AuthorRepository) GetByID(id int) (*Author, error) {
	var author Author
	err := r.DB.QueryRow("SELECT id, first_name, last_name, biography, birth_date FROM authors WHERE id = $1", id).Scan(&author.ID, &author.FirstName, &author.LastName, &author.Biography, &author.BirthDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, appErrors.ErrNotFound
		}
		return nil, appErrors.ErrInternalServer
	}
	return &author, nil
}

func (r *AuthorRepository) Update(author *Author) error {
	_, err := r.DB.Exec("UPDATE authors SET first_name = $1, last_name = $2, biography = $3, birth_date = $4 WHERE id = $5", author.FirstName, author.LastName, author.Biography, author.BirthDate, author.ID)
	return err
}

func (r *AuthorRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM authors WHERE id = $1", id)
	if err != nil {
		return appErrors.ErrInternalServer
	}
	return nil
}

func (r *AuthorRepository) UpdateBioInTransaction(tx *sql.Tx, authorID int, bio string) error {
	_, err := tx.Exec("UPDATE authors SET biography = $1 WHERE id = $2", bio, authorID)
	if err != nil {
		return appErrors.ErrInternalServer
	}
	return nil
}
