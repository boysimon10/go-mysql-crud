package repository

import (
    "database/sql"
    "github.com/boysimon10/go-mysql-crud/internal/models"
)

type BookRepository struct {
    db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
    return &BookRepository{db: db}
}

func (r *BookRepository) Create(book *models.Book) error {
    query := `INSERT INTO books (title, author, price) VALUES (?, ?, ?)`
    result, err := r.db.Exec(query, book.Title, book.Author, book.Price)
    if err != nil {
        return err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    book.ID = id
    return nil
}

func (r *BookRepository) GetByID(id int64) (*models.Book, error) {
    book := &models.Book{}
    query := `SELECT id, title, author, price, created_at FROM books WHERE id = ?`
    err := r.db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.CreatedAt)
    if err != nil {
        return nil, err
    }
    return book, nil
}

func (r *BookRepository) GetAll() ([]models.Book, error) {
    var books []models.Book
    query := `SELECT id, title, author, price, created_at FROM books`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var book models.Book
        if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.CreatedAt); err != nil {
            return nil, err
        }
        books = append(books, book)
    }
    return books, nil
}

func (r *BookRepository) Update(book *models.Book) error {
    query := `UPDATE books SET title=?, author=?, price=? WHERE id=?`
    _, err := r.db.Exec(query, book.Title, book.Author, book.Price, book.ID)
    return err
}

func (r *BookRepository) Delete(id int64) error {
    query := `DELETE FROM books WHERE id=?`
    _, err := r.db.Exec(query, id)
    return err
}