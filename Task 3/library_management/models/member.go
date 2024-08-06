package models

// Member represents a library member.
type Member struct {
	ID            int    // Unique identifier for the member.
	Name          string // Name of the member.
	BorrowedBooks []Book // List of books borrowed by the member.
}
