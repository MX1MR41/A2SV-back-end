# Library Management System Documentation

## Overview

The Library Management System is a command-line application written in Go that allows users to manage a library's books and members. The system supports adding and removing books, borrowing and returning books, listing available and borrowed books, and managing library members.

## Project Structure

```
library_management/
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   └── book.go
│   └── member.go
├── services/
│   └── library_service.go
├── docs/
│   └── documentation.md
└── go.mod
```

## Module Descriptions

### `main.go`
This file contains the main entry point of the application. It imports the `controllers` package and calls the `Run` function to start the application.

### `controllers/library_controller.go`
This file contains the controller logic that interacts with the user. It displays the menu, handles user input, and calls the appropriate service functions to perform actions such as adding books, borrowing books, and listing available books.

### `models/book.go`
Defines the `Book` struct, which represents a book in the library. A `Book` has an ID, Title, Author, and Status to indicate availability.

### `models/member.go`
Defines the `Member` struct, which represents a library member. A `Member` has an ID, Name, and a list of borrowed books.

### `services/library_service.go`
This file contains the `Library` struct and methods to manage books and members. The `Library` struct includes maps to store books and members and provides methods to add, remove, borrow, and return books, as well as to list available and borrowed books.

## Usage

1. **Run the Application**:
   - Execute the `main.go` file to start the application.
      ```
      go run main.go
      ```

2. **Menu Options**:
   - **Add Book**: Adds a new book to the library.
   - **Remove Book**: Removes a book from the library using its ID.
   - **Borrow Book**: Borrows a book by specifying the book ID and member ID.
   - **Return Book**: Returns a borrowed book by specifying the book ID and member ID.
   - **List Available Books**: Lists all books that are currently available in the library.
   - **List Borrowed Books**: Lists all books borrowed by a specific member.
   - **Add Member**: Adds a new member to the library.
   - **List Members**: Lists all members of the library.
   - **Exit**: Exits the application.

3. **Adding a Book**:
   - Enter the book ID, title, and author when prompted.

4. **Removing a Book**:
   - Enter the book ID of the book to be removed.

5. **Borrowing a Book**:
   - Enter the book ID and the member ID when prompted.

6. **Returning a Book**:
   - Enter the book ID and the member ID when prompted.

7. **Listing Available Books**:
   - Displays all books with the status "Available".

8. **Listing Borrowed Books**:
   - Enter the member ID to list all books borrowed by that member.

9. **Adding a Member**:
   - Enter the member ID and name when prompted.

10. **Listing Members**:
    - Displays all members registered in the library.

## Error Handling
- The application handles common errors such as attempting to borrow a book that is already borrowed or returning a book that is not borrowed.

## Extending the Application
- New features and functionalities can be added by extending the `Library` struct and implementing new methods in the `services` package.

## Conclusion
This documentation provides an overview of the Library Management System's structure and usage. For detailed implementation, refer to the source code in the respective files.

---