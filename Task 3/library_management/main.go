/*
A simple console-based library management system that allows users
to add, remove, borrow, return, list available and borrowed books in the library.
It also allows users to add members and list in the library.
*/

package main

import (
	"library_management/controllers"
	"library_management/services"
)

// Main entry point of the application
func main() {
	// Initialize the LibraryManager Interface and assign it with a new Library struct
	var library services.LibraryManager = services.NewLibrary()

	// Pass the LibraryManager to the controller
	controllers.Run(library)
}
