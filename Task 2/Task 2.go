package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// this function checks whether a string is a palindrome or not
// after removing non-alphanumeric characters and ignoring case
func palindrome(s string) bool {
	s = strings.ToLower(s)
	// remove non-alphanumeric characters
	var cleaned string
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			cleaned += string(r)
		}
	}

	// Check if cleaned string is a palindrome
	for i := 0; i < len(cleaned)/2; i++ {
		if cleaned[i] != cleaned[len(cleaned)-1-i] {
			return false
		}
	}
	return true
}

// this function returns a map with the frequency of each word in the string,
// after removing non-alphanumeric characters and ignoring case
func wordFrequencyCount(s string) map[string]int {
	s = strings.ToLower(s)
	words := strings.Fields(s)

	// count frequency of each word and store in map
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	return m
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var choice int

	// Show the possible options to the user
	fmt.Println("Press 1 to count the frequency of words in a string.")
	fmt.Println("Press 2 to check if a string is a palindrome.")
	fmt.Println("Press 3 to exit.")

	fmt.Scanln(&choice)
	if choice == 1 {

		fmt.Print("Enter a string: ")
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)

		// Display word frequency count
		frequencies := wordFrequencyCount(s)
		for word, count := range frequencies {
			fmt.Printf("%s: %d\n", word, count)
		}

	} else if choice == 2 {

		fmt.Print("Enter a string: ")
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)

		// Check and display if the string is a palindrome
		if palindrome(s) {
			fmt.Println("The string is a palindrome.")
		} else {
			fmt.Println("The string is not a palindrome.")
		}

	} else if choice == 3 {

		return

	} else {
		fmt.Println("Invalid choice, please try again.")
	}
}
