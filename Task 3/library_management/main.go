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
