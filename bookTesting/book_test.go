package bookTesting

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/bxcodec/faker/v4"
)

func createBook(t *testing.T) *Book {
	book, err := DBmanager.CreateBook(&Book{
		Title: faker.Sentence(),
		AuthorName: faker.Name(),
		Price: 5000,
		Amount: 10,
	})
	require.NoError(t, err)
	require.NotEmpty(t, book)
	return book
} 

func deleteBook(t *testing.T, book_id int64) {
	err := DBmanager.DeleteBook(book_id)
	require.NoError(t, err)
} 


func TestCreateBook(t *testing.T) {
	book := createBook(t)
	deleteBook(t, book.Id)
	require.NotEmpty(t, book)
}

func TestGetBook(t *testing.T) {
	book := createBook(t)
	book_get, err := DBmanager.GetBook(book.Id)
	deleteBook(t, book_get.Id)
	require.NoError(t, err)
	require.NotEmpty(t, book_get)
}

func TestUpdateBook(t *testing.T) {
	book := createBook(t)
	book_update, err := DBmanager.UpdateBook(book)
	deleteBook(t, book.Id)
	require.NoError(t, err)
	require.NotEmpty(t, book_update)
}

func TestGetAllBook(t *testing.T) {
	book := createBook(t)
	books, err := DBmanager.GetAllBook(&BookParam{
		Limit: 10,
		Page: 1,
	})
	require.GreaterOrEqual(t, len(books.Books), 1)
	require.NoError(t, err)
	deleteBook(t, book.Id)
	require.NotEmpty(t, books)
	require.NotEmpty(t, book)
}