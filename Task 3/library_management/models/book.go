package models

// Book represents a book in the library.
type Book struct {
	ID     int    // ID is the unique identifier of the book.
	Title  string // Title is the title of the book.
	Author string // Author is the author of the book.
	Status string // Status represents the availability status of the book.
}
