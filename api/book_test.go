package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0123456789"}`, string(json), "Book JSON marshalling wrong.")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0123456789"}`)
	book := FromJSON(json)
	assert.Equal(t, Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"}, book, "Book JSON unmarshalling wrong.")
}

func TestAllBooks(t *testing.T) {
	books := AllBooks()
	assert.Len(t, books, 2, "Wrong number of books.")
}

func TestCreateNewBook(t *testing.T) {
	book := Book{Title: "Test", Author: "Me", ISBN: "1234567890"}
	isbn, created := CreateBook(book)
	assert.True(t, created, "Book was not created.")
	assert.Equal(t, "1234567890", isbn, "Wrong ISBN.")
}

func TestDoNotCreateExistingBook(t *testing.T) {
	book := Book{ISBN: "0000000000"}
	_, created := CreateBook(book)
	assert.False(t, created, "Book was created.")
}

func TestUpdateExistingBook(t *testing.T) {
	book := Book{Title: "Test Update", Author: "Me Again", ISBN: "1234567890", Description: "Dummy Text"}
	updated := UpdateBook("1234567890", book)
	assert.True(t, updated, "Book not updated.")

	book, _ = GetBook("1234567890")
	assert.Equal(t, "Test Update", book.Title, "Title not updated.")
	assert.Equal(t, "Me Again", book.Author, "Author not updated.")
	assert.Equal(t, "Dummy Text", book.Description, "Description not updated.")
}

func TestDeleteBook(t *testing.T) {
	DeleteBook("1234567890")
	assert.Len(t, AllBooks(), 2, "Wrong number of books after delete.")
}
