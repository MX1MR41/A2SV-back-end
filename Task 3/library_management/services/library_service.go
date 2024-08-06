package services

import (
	"errors"
	"library_management/models"
)

// Library represents a library.
type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

// NewLibrary creates a new Library.
func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

// LibraryManager is an interface that contains list of methods for managing a library.
type LibraryManager interface {
	AddBook(book models.Book)                     // AddBook adds a book to the library.
	RemoveBook(bookID int)                        // RemoveBook removes a book from the library.
	BorrowBook(bookID int, memberID int) error    // BorrowBook borrows a book from the library.
	ReturnBook(bookID int, memberID int) error    // ReturnBook returns a borrowed book to the library.
	ListAvailableBooks() []models.Book            // ListAvailableBooks lists all available books in the library.
	ListBorrowedBooks(memberID int) []models.Book // ListBorrowedBooks lists all books borrowed by a member.
	AddMember(member models.Member)               // AddMember adds a new member to the library.
	ListMembers() []models.Member                 // ListMembers lists all members in the library.
}

// Implemntation of LibraryManager interface

func (l *Library) AddMember(member models.Member) {
	l.Members[member.ID] = member
}

func (l *Library) ListMembers() []models.Member {
	members := []models.Member{}
	for _, member := range l.Members {
		members = append(members, member)
	}
	return members
}

func (l *Library) AddBook(book models.Book) {
	l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.Books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}

	member, exists := l.Members[memberID]
	if !exists {
		member = models.Member{ID: memberID, BorrowedBooks: []models.Book{}}
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member

	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Available" {
		return errors.New("book is not borrowed")
	}

	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book.Status = "Available"
	l.Books[bookID] = book

	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}
	l.Members[memberID] = member

	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	availableBooks := []models.Book{}
	for _, book := range l.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.Members[memberID]
	if !exists {
		return []models.Book{}
	}
	return member.BorrowedBooks
}
