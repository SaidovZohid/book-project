package bookTesting

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBManager struct {
	db *sqlx.DB
}

type Book struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	AuthorName string    `json:"author_name"`
	Price      float64   `json:"price"`
	Amount     int       `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewDbmanager(db *sqlx.DB) *DBManager {
	return &DBManager{
		db: db,
	}
}

func (d *DBManager) CreateBook(book *Book) (*Book, error) {
	query := `
		INSERT INTO book_info (
			title,
			author_name,
			price,
			amount
		) VALUES ($1, $2, $3, $4)
		RETURNING id, title, author_name, price, amount, created_at
	`
	row := d.db.QueryRow(
		query,
		book.Title,
		book.AuthorName,
		book.Price, book.
			Amount,
	)
	var b Book
	err := row.Scan(
		&b.Id,
		&b.Title,
		&b.AuthorName,
		&b.Price,
		&b.Amount,
		&b.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (d *DBManager) GetBook(book_id int64) (*Book, error) {
	var b Book
	queryGetBook := `
		SELECT 
			id,
			title,
			author_name,
			price,
			amount,
			created_at
		FROM book_info WHERE id = $1

	`
	row := d.db.QueryRow(queryGetBook, book_id)
	err := row.Scan(
		&b.Id,
		&b.Title,
		&b.AuthorName,
		&b.Price,
		&b.Amount,
		&b.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (d *DBManager) UpdateBook(book *Book) (*Book, error){
	query := `
		UPDATE book_info SET
			title = $1,
			author_name = $2,
			price = $3,
			amount = $4
		WHERE id = $5 
		RETURNING id, title, author_name, price, amount, created_at
	`
	row := d.db.QueryRow(
		query, 
		book.Title,
		book.AuthorName,
		book.Price,
		book.Amount,
		book.Id,
	)
	var b Book
	err := row.Scan(
		&b.Id,
		&b.Title,
		&b.AuthorName,
		&b.Price,
		&b.Amount,
		&b.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (d *DBManager) DeleteBook(book_id int64) error {
	query := `
		DELETE FROM book_info WHERE id = $1
	`
	result, err := d.db.Exec(query, book_id)
	if err != nil {
		return err
	}
	res, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if res == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (d *DBManager) GetAllBook() ([]*Book, error) {
	query := `
		SELECT 
			id,
			title,
			author_name,
			price,
			amount,
			created_at
		FROM book_info
	`
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}

	var books []*Book
	for rows.Next() {
		var book Book
		err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.AuthorName,
			&book.Price,
			&book.Amount,
			&book.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}
