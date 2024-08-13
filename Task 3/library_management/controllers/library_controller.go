package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
)

func Run(library services.LibraryManager) {
	// Initializing a scanner to read input from the console
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Display menu options
		fmt.Println("+------------------------------------------------------------+")
		fmt.Println("Library Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Add Member")
		fmt.Println("8. List Members")
		fmt.Println("9. Exit")
		fmt.Print("Enter choice: ")

		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		// Perform the corresponding action based on the user's choice
		case 1:
			fmt.Println("+------------------------------------------------------------+")
			addBook(scanner, library)
		case 2:
			fmt.Println("+------------------------------------------------------------+")
			removeBook(scanner, library)
		case 3:
			fmt.Println("+------------------------------------------------------------+")
			borrowBook(scanner, library)
		case 4:
			fmt.Println("+------------------------------------------------------------+")
			returnBook(scanner, library)
		case 5:
			fmt.Println("+------------------------------------------------------------+")
			listAvailableBooks(library)
		case 6:
			fmt.Println("+------------------------------------------------------------+")
			listBorrowedBooks(scanner, library)
		case 7:
			fmt.Println("+------------------------------------------------------------+")
			addMember(scanner, library)
		case 8:
			fmt.Println("+------------------------------------------------------------+")
			listMembers(library)
		case 9:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// Function to add a book to the library
// It accepts a scanner to read input from the console
// And prints the corresponding messages based on the result
func addBook(scanner *bufio.Scanner, library services.LibraryManager) {
	fmt.Print("Enter book ID: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter book title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Enter book author: ")
	scanner.Scan()
	author := scanner.Text()

	book := models.Book{ID: id, Title: title, Author: author, Status: "Available"}
	library.AddBook(book)
	fmt.Println("Book added successfully.")
}

// Function to remove a book from the library
// It accepts a scanner to read input from the console
// And prints the corresponding messages based on the result
func removeBook(scanner *bufio.Scanner, library services.LibraryManager) {
	fmt.Print("Enter book ID to remove: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	library.RemoveBook(id) // Remove the book with the given ID whether it exists or not
	fmt.Println("Book removed successfully.")
}

// Function to borrow a book from the library
// It accepts a scanner to read input from the console
// And prints the corresponding messages based on the result
func borrowBook(scanner *bufio.Scanner, library services.LibraryManager) {
	fmt.Print("Enter book ID to borrow: ")
	scanner.Scan()
	bookID, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())
	// Calls the respective borrow method of the library service
	// And prints the corresponding messages based on the result
	err := library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

// Function to return a book to the library
// It accepts a scanner to read input from the console
// And prints the corresponding messages based on the result
func returnBook(scanner *bufio.Scanner, library services.LibraryManager) {
	fmt.Print("Enter book ID to return: ")
	scanner.Scan()
	bookID, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())
	// Calls the respective return method of the library service
	// And prints the corresponding messages based on the result
	err := library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

// Function to list all available books in the library
// It accepts no arguments and prints the list of available books
func listAvailableBooks(library services.LibraryManager) {
	// Calls the ListAvailableBooks method of the library service to get the list of available books
	books := library.ListAvailableBooks()
	// If there are no available books, print the following message
	if len(books) == 0 {
		fmt.Println("No available books.")
		return
	}

	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

// Function to list all borrowed books by a member
// It accepts a scanner to read input from the console
// And prints the list of borrowed books by the member
func listBorrowedBooks(scanner *bufio.Scanner, library services.LibraryManager) {
	fmt.Print("Enter member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())
	// Calls the ListBorrowedBooks method of the library service
	// to get the list of borrowed books by the member
	books := library.ListBorrowedBooks(memberID)
	// If there are no borrowed books by the member, print the following message
	if len(books) == 0 {
		fmt.Println("No borrowed books for this member.")
		return
	}
	// Print the list of borrowed books by the member
	fmt.Println("Borrowed Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

// Function to add a new member to the library
// It accepts a scanner to read input from the console
// And prints the corresponding messages based on the result
func addMember(scanner *bufio.Scanner, library services.LibraryManager) {
	fmt.Print("Enter member ID: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter member name: ")
	scanner.Scan()
	name := scanner.Text()

	member := models.Member{ID: id, Name: name, BorrowedBooks: []models.Book{}}
	library.AddMember(member)
	fmt.Println("Member added successfully.")
}

// Function to list all members in the library
// It accepts no arguments and prints the list of members
// Calls the ListMembers method of the library service to get the list of members
func listMembers(library services.LibraryManager) {
	members := library.ListMembers()
	if len(members) == 0 {
		fmt.Println("No members.")
		return
	}

	fmt.Println("Members:")
	for _, member := range members {
		fmt.Printf("ID: %d, Name: %s\n", member.ID, member.Name)
	}
}
